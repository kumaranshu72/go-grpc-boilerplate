package model

// Configuration for application to export configuration logic
type Configuration struct {
	Environment   string
	GRPCPort      string
	HTTPPort      string
	LogLevel      int
	LogTimeFormat string
}
