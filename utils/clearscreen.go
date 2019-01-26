package utils

import (
	"os"
	"runtime"
	"os/exec"
)

func ClearScreen() {

	var clear = make(map[string]func())
    clear["linux"] = func() { 
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }

    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        panic(" Something went wrong!!\n Could not clear the screen")
    }
}