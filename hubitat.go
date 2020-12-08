package hubitat

import (
	"encoding/json"
	"fmt"
	"net/url"
)

var (
	ErrMandatoryID      error = fmt.Errorf("ID is mandatory")
	ErrMandatoryCommand error = fmt.Errorf("Command is mandatory")
	ErrMandatoryValue   error = fmt.Errorf("Value is mandatory")
)

type Config struct {
	URL         *url.URL
	AccessToken string
}

func New(host, accessToken string) (*Config, error) {
	dataUrl, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	c := Config{
		URL:         dataUrl,
		AccessToken: accessToken,
	}

	return &c, nil
}

func (c *Config) GetModes() (*Modes, error) {
	urlPath := "/modes"

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(Modes)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Config) SetMode(id string) error {
	if len(id) == 0 {
		return ErrMandatoryID
	}

	urlPath := fmt.Sprintf("/modes/%s", id)

	u := c.buildURL(urlPath)

	_, err := GetURL(u)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) GetHsmStatus() (*Hsm, error) {
	urlPath := "/hsm"

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(Hsm)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Config) SetHsmStatus(value string) error {
	if len(value) == 0 {
		return ErrMandatoryValue
	}

	urlPath := fmt.Sprintf("/hsm/%s", value)

	u := c.buildURL(urlPath)

	_, err := GetURL(u)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) GetDevices() (*Devices, error) {
	urlPath := "/devices"

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(Devices)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Config) GetDevicesWithFullDetails() (*DevicesAll, error) {
	urlPath := "/devices/all"

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(DevicesAll)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get Device Info with actual subscribed device id
func (c *Config) GetDevice(id string) (*DeviceInfo, error) {
	if len(id) == 0 {
		return nil, ErrMandatoryID
	}

	urlPath := fmt.Sprintf("/devices/%s", id)

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(DeviceInfo)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get Device Event History with actual subscribed device id
func (c *Config) GetDeviceEventHistory(id string) (*DeviceEventHistory, error) {
	if len(id) == 0 {
		return nil, ErrMandatoryID
	}

	urlPath := fmt.Sprintf("/devices/%s/events", id)

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(DeviceEventHistory)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get Device Commands with actual subscribed device id
func (c *Config) GetDeviceCommands(id string) (*DeviceCommands, error) {
	if len(id) == 0 {
		return nil, ErrMandatoryID
	}

	urlPath := fmt.Sprintf("/devices/%s/commands", id)

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(DeviceCommands)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Get Device Capabilities with actual subscribed device id
func (c *Config) GetDeviceCapabilities(id string) (*DeviceCapabilities, error) {
	if len(id) == 0 {
		return nil, ErrMandatoryID
	}

	urlPath := fmt.Sprintf("/devices/%s/capabilities", id)

	u := c.buildURL(urlPath)

	resp, err := GetURL(u)
	if err != nil {
		return nil, err
	}

	result := new(DeviceCapabilities)

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Send Device Command with actual subscribed device id and with a supported command. Supports optional secondary value
func (c *Config) SendDeviceCommand(id, command, value string) error {
	if len(id) == 0 {
		return ErrMandatoryID
	}
	if len(command) == 0 {
		return ErrMandatoryCommand
	}

	urlPath := fmt.Sprintf("/devices/%s/%s", id, command)

	if len(value) > 0 {
		urlPath = fmt.Sprintf("%s/%s", urlPath, value)
	}

	u := c.buildURL(urlPath)

	_, err := GetURL(u)
	if err != nil {
		return err
	}

	return nil
}
