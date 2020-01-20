package config

import "github.com/spf13/viper"

// IConfig is
type IConfig interface {
	GetString(key, defaultValue string) string
	GetInt(key, defaultValue string) int
	GetBool(key, defaultValue string) bool
}

// SimpleConfig is
type SimpleConfig struct {
}

// NewSimpleConfig is
func NewSimpleConfig(confiName, path string) *SimpleConfig {
	viper.SetConfigName(confiName)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	rc := SimpleConfig{}
	return &rc
}

// GetString is
func (r *SimpleConfig) GetString(key, defaultValue string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultValue
	}
	return v
}

// GetInt is
func (r *SimpleConfig) GetInt(key, defaultValue string) int {
	v := viper.GetInt(key)
	return v
}

// GetBool is
func (r *SimpleConfig) GetBool(key, defaultValue string) bool {
	v := viper.GetBool(key)
	return v
}
