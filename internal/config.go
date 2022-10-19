package internal

type Configuration struct {
	Settings ConfigurationSettings `json:"settings"`
}

type ConfigurationSettings struct {
	Adapter        string         `json:"adapter"`
	AdapterOptions AdapterOptions `json:"adapterOptions"`
}

type AdapterOptions struct {
	Server   string             `json:"server"`
	Port     int64              `json:"port"`
	Database string             `json:"database"`
	UseTLS   bool               `json:"useTLS"`
	Auth     AdapterAuthOptions `json:"auth"`
}

type AdapterAuthOptions struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
