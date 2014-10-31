package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Port   int    `json:"port"`
	DbFile string `json:"dbFile"`
}

func (c *config) read(env *string) error {
	file, err := os.Open("config/" + *env + ".json")
	if err != nil {
		return err
	}
	return json.NewDecoder(file).Decode(&c)
}
