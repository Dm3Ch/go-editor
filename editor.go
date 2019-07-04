package editor

import (
	"os"
	"os/exec"
)

const defaultEditor = "vi"

func getDefaultEditor() string {
	var value string
	var ok bool

	value, ok = os.LookupEnv("VISUAL")
	if ok {
		return value
	}

	value, ok = os.LookupEnv("EDITOR")
	if ok {
		return value
	}

	return defaultEditor
}

func RunEditor(path string) error {
	cmd := exec.Command(getDefaultEditor(), path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
