package hubitat

type Modes []Mode

type Mode struct {
	Active bool   `json:"active"`
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
}

type Hsm struct{}

type Devices []Device

type Device struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type DevicesAll []DeviceFullDetails

type DeviceFullDetails struct {
	Device
	Type         string                 `json:"type"`
	Date         string                 `json:"date"`
	Model        interface{}            `json:"model"`
	Manufacturer interface{}            `json:"manufacturer"`
	Capabilities []string               `json:"capabilities"`
	Attributes   map[string]interface{} `json:"attributes"`
	Commands     []DeviceCommand        `json:"commands"`
}

type DeviceInfo struct {
	Device
	Attributes []DeviceAttribute `json:"attributes"`
	//Capabilities DeviceCapabilities `json:"capabilities"`
	Commands []string `json:"commands"`
}

type DeviceAttribute struct {
	Name         string      `json:"name"`
	CurrentValue interface{} `json:"currentValue"`
	DataType     string      `json:"dataType"`
	Values       []string    `json:"values,omitempty"`
}

type DeviceEventHistory []DeviceEvent

type DeviceEvent struct {
	DeviceID      string `json:"device_id"`
	Label         string `json:"label"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	Date          string `json:"date"`
	Unit          string `json:"unit"`
	IsStateChange string `json:"isStateChange"`
	Source        string `json:"source"`
}

type DeviceCommands []DeviceCommand

type DeviceCommand struct {
	Command string   `json:"command"`
	Type    []string `json:"type,omitempty"`
}

type DeviceCapabilities []DeviceCapability

type DeviceCapability struct {
}
