package main

import (
	"encoding/json"
	"log"
	"os"
)

type Server struct {
	User string `json:"user"`
	Host string `json:"host"`
}

type ServersFile struct {
	Servers []*Server
}

func (s *ServersFile) Load() error {
	file, err := os.Open("servers.json")

	if err != nil {
		log.Printf("[ServersFile.Load] Error while opening config file: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s)

	if err != nil {
		log.Printf("[ServersFile.Load] Error while decoding JSON: %v", err)
		return err
	}

	return nil
}
