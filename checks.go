package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// CheckDefinition represents the JSON definition of the check
type CheckDefinition struct {
	Name    string   `json:"name"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

// ChecksManager represents a manager object for servers checks
type ChecksManager struct {
	Checks      map[string]interface{}
	Definitions []*CheckDefinition
}

// Load loads the checks JSON file into the ChecksManager object and parses it
func (c *ChecksManager) Load(checksFile string) error {
	file, err := os.Open(checksFile)

	if err != nil {
		log.Printf("[Config.Load] Error while opening config file: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.Checks)

	if err != nil {
		log.Printf("[Config.Load] Error while decoding JSON: %v", err)
		return err
	}

	return nil
}

// Load loads the checks JSON file into the ChecksManager object and parses it
func (c *ChecksManager) LoadDefinitions(definitionsFile string) error {
	file, err := os.Open(definitionsFile)

	if err != nil {
		log.Printf("[Config.Load] Error while opening config file: %v", err)
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.Definitions)

	if err != nil {
		log.Printf("[Config.Load] Error while decoding JSON: %v", err)
		return err
	}

	return nil
}

// ProcessChecks goes through checks and calls the appropriate methods to do the checks
func (c *ChecksManager) ProcessChecks(server *Server) error {
	var err error
	var e error

	for checkName, check := range c.Checks {
		chkParsed := check.(map[string]interface{})
		command := c.buildCommand(chkParsed)

		if verbose {
			log.Printf("Built command for server '%s': %s", server.Host, command)
		}

		e = RunSSHCommand(server.User, server.Host, command)

		if e != nil && err == nil {
			err = fmt.Errorf("FAILED: Check '%s' of type '%s' failed on server '%s'", checkName, chkParsed["type"], server.Host)
		}
	}

	return err
}

func (c *ChecksManager) buildCommand(check map[string]interface{}) string {
	var command string

	for _, definition := range c.Definitions {
		if definition.Name == check["type"] {
			var args []interface{}
			for _, arg := range definition.Args {
				args = append(args, check[arg])
			}
			command = fmt.Sprintf(definition.Command, args...)
		}
	}

	return command
}
