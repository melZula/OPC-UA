package main

// Config struct
type Config struct {
	URL    string   `json:"url"`
	Nodes  []string `json:"nodes"`
	Min    float32  `json:"min_value"`
	Max    float32  `json:"max_value"`
	Freq   string   `json:"freq"`
	APIKey string   `json:"telegram_api_key"`
	ChanID int64    `json:"channel_id"`
}

// NewConfig init
func NewConfig() *Config {
	return &Config{
		URL: "opc.tcp://localhost:53530/OPCUA/SimulationServer",
	}
}
