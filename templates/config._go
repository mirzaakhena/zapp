package config

import "github.com/spf13/viper"

// IConfig is
type IConfig interface {
	GetString(key, defaultValue string) string
	GetInt(key, defaultValue string) int
	GetBool(key, defaultValue string) bool
}

// RealtimeConfig is
type RealtimeConfig struct {
}

// NewRealtimeConfig is
func NewRealtimeConfig(confiName, path string) *RealtimeConfig {
	viper.SetConfigName(confiName)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	rc := RealtimeConfig{}
	// rc.reload()
	return &rc
}

func (r *RealtimeConfig) GetString(key, defaultValue string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func (r *RealtimeConfig) GetInt(key, defaultValue string) int {
	v := viper.GetInt(key)

	return v
}

func (r *RealtimeConfig) GetBool(key, defaultValue string) bool {
	v := viper.GetBool(key)

	return v
}
