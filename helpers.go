package main

import (
	"encoding/json"
	"io/ioutil"
)

func readFile(filepath string) ([]byte, error) {
	filebytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return filebytes, err
}

func getBotSettings() (*botSettings, error) {
	settingsBytes, err := readFile(settingsFilePath)
	if err != nil {
		return nil, err
	}
	settings := botSettings{}
	err = json.Unmarshal(settingsBytes, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}
