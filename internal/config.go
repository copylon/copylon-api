package internal

type Configuration struct {
	Settings ConfigurationSettings `json:"settings"`
	OAuth    OAuthOptions          `json:"oauth"`
}

type ConfigurationSettings struct {
	Port           uint           `json:"port"`
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

type OAuthOptions struct {
	ServerURI    string `json:"server_uri"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	CallbackURI  string `json:"callback_uri"`
}
