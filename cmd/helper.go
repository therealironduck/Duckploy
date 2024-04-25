package cmd

import (
	"Duckploy/data"
	"encoding/json"
	"fmt"
	"os"
)

func exitWithErrorf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", a...)
	exitFunc(1)
}

var readConfig = func(path string) data.Config {
	jsonFile, err := os.Open(path)
	if err != nil {
		exitWithErrorf("Config cannot be read: %v", err)
	}

	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	var config data.Config
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&config)

	if err != nil {
		exitWithErrorf("Failed to decode config file: %v", err)
	}

	return config
}
