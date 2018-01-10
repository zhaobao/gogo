package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func execShell(name string, args ...string) (error, string) {
	out := &bytes.Buffer{}
	cmd := exec.Command(name, args...)
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		return err, out.String()
	}
	return nil, out.String()
}

func main() {
	time.Sleep(time.Second * 2)
	err, str := execShell("osascript", "-e", "tell application \"System Events\"", "-e", "set frontApp to name of first application process whose frontmost is true", "-e", "end tell")
	if err == nil {
		fmt.Println(str)
	}
	execShell("osascript", "-e", "tell application \"System Events\" to keystroke \"t\"")
}
