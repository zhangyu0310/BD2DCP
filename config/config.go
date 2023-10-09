package config

import "sync/atomic"

type Config struct {
	VideoInfoMap map[string]*VideoInfo
	AutoShutdown bool
}

type VideoInfo struct {
	VideoPath    string `json:"video_path"`
	SubtitlePath string `json:"subtitle_path"`
	AudioChannel uint   `json:"audio_channel"`
	AutoBitrate  bool   `json:"auto_bitrate"`
	Bitrate      string `json:"bitrate"`
	AutoCrop     bool   `json:"auto_crop"`
	CropWidth    uint   `json:"crop_width"`
	CropHeight   uint   `json:"crop_height"`
}

var globalCfg atomic.Value

// InitializeConfig initialize the global config handler.
func InitializeConfig(enforceCmdArgs func(*Config)) {
	cfg := Config{}
	// Use command config cover config file.
	enforceCmdArgs(&cfg)
	StoreGlobalConfig(&cfg)
}

func InitializeConfigFromFile(configFile string) {

}

// GetGlobalConfig returns the global configuration for this server.
// It should store configuration from command line and configuration file.
// Other parts of the system can read the global configuration use this function.
func GetGlobalConfig() *Config {
	return globalCfg.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalCfg.Store(config)
}
