package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func restartCoreAudio() error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("this program is intended to run on macOS only")
	}
	cmd := exec.Command("sudo", "killall", "-9", "coreaudiod")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	err := restartCoreAudio()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
