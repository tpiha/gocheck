package main

import (
	"encoding/json"
	"log"
	"os"
)

// Server represents the server containing hostname and username
type Server struct {
	User string `json:"user"`
	Host string `json:"host"`
}

// ServersFile represents the servers JSON file object
type ServersFile struct {
	Servers []*Server
}

// Load loads the servers JSON file into the ServersFile object and parses it
func (s *ServersFile) Load(serversFile string) error {
	file, err := os.Open(serversFile)

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
