package config

import (
	"Duckploy/helper"
	"encoding/json"
)

type Step struct {
	Action  string `json:"action"`
	Command string `json:"command"`
}

type Host struct {
	Hostname    string `json:"hostname"`
	SSHUser     string `json:"ssh_user"`
	SSHPassword string `json:"ssh_password"`
	Path        string `json:"path"`
}

type Config struct {
	Steps []Step `json:"steps"`
	Hosts []Host `json:"hosts"`
}

func ReadConfig(path string) (config Config, err error) {
	jsonFile, err := helper.AppFs.Open(path)
	if err != nil {
		return
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)

	return
}
