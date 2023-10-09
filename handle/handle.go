package handle

import (
	"BD2DCP/config"
	"fmt"
	"os/exec"
	"runtime"
)

func Handle() error {
	cfg := config.GetGlobalConfig()

	if cfg.AutoShutdown {
		// Need super permission.
		ShutdownSystem()
	}
	return nil
}

func ShutdownSystem() {
	goOS := runtime.GOOS
	switch goOS {
	case "windows":
		cmd := exec.Command("shutdown", "-s", "-t", "10")
		_ = cmd.Run()
	case "linux":
		cmd := exec.Command("poweroff")
		_ = cmd.Run()
	default:
		fmt.Printf("Unsupported auto shutdown OS %s,"+
			"only support windows and linux.\n", goOS)
	}
}
