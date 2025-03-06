package main

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName  string
	Port     int
	Database string
}

var (
	configInstance *Config
	onceConfig     sync.Once
)

func LoadConfig() *Config {
	onceConfig.Do(func() {
		fmt.Println("Loading the configuration")
		configInstance = &Config{
			AppName:  "Anshul",
			Port:     8080,
			Database: "Postgres",
		}
	})
	return configInstance
}

func main() {

	config1 := LoadConfig()
	config2 := LoadConfig()

	// If below statement return true it means both config1 and config2 are pointing to
	//  the same memory location

	fmt.Println("Checking config1 and config2 the same instance, If 'true' means both are pointing to same memory location--->", config1 == config2)

	fmt.Println("App Name ---->>", config1.AppName)
	fmt.Println("Port ---->>", config1.Port)
	fmt.Println("Database Name ---->>", config1.Database)

}
