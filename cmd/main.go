package main

import (
	"BD2DCP/config"
	"BD2DCP/handle"
	"BD2DCP/preprocess"
	"flag"
	"fmt"
)

var (
	cfgPath      = flag.String("config", "", "config path - json format, can handle multiple video")
	videoPath    = flag.String("video", "", "video path")
	subtitlePath = flag.String("subtitle", "", "subtitle path")
	audioChannel = flag.Uint("audio-channel", 0, "audio channel")
	bitrate      = flag.String("bitrate", "", "bitrate")
	autoCrop     = flag.Bool("auto-crop", false, "auto crop")
	autoShutDown = flag.Bool("auto-shutdown", false, "auto shutdown")
)

func main() {
	help := flag.Bool("help", false, "help")
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	if *cfgPath != "" {
		config.InitializeConfigFromFile(*cfgPath)
	} else {
		config.InitializeConfig(initConfig)
	}

	if err := preprocess.Process(); err != nil {
		fmt.Println(err)
		return
	}

	if err := handle.Handle(); err != nil {
		fmt.Println(err)
		return
	}
}

func initConfig(cfg *config.Config) {
	// If path not set, find video & subtitle file in current directory.
	cfg.AutoShutdown = *autoShutDown
	config.StoreGlobalConfig(cfg)
}
