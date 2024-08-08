package config

var configInstance *Config

func GetConfig() *Config {
    if configInstance == nil {
        configInstance, _ = LoadConfig()
    }
    return configInstance
}