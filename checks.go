package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// ChecksManager represents a manager object for servers checks
type ChecksManager struct {
	Checks map[string]interface{}
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

// ProcessChecks goes through checks and calls the appropriate methods to do the checks
func (c *ChecksManager) ProcessChecks(server *Server) error {
	var err error
	var e error

	for checkName, check := range c.Checks {
		chkParsed := check.(map[string]interface{})
		switch chkParsed["type"] {
		case "file_contains":
			e = c.fileContains(server, chkParsed["path"].(string), chkParsed["check"].(string))
		case "file_exists":
			e = c.fileExists(server, chkParsed["path"].(string))
		case "process_running":
			e = c.processRunning(server, chkParsed["name"].(string))
		default:
			log.Printf("[CheckManager.ProcessChecks] Check not found: %s", chkParsed["type"])
		}
		if e != nil && err == nil {
			err = fmt.Errorf("Check '%s' of type '%s' failed on server '%s'", checkName, chkParsed["type"], server.Host)
		}
	}

	return err
}

func (c *ChecksManager) buildCommand() string {

}

// fileContains checks if some file contains some string
func (c *ChecksManager) fileContains(server *Server, path, content string) error {
	command := fmt.Sprintf("grep %s %s", content, path)
	err := RunSSHCommand(server.User, server.Host, command)
	return err
}

// fileExists checks if file exists on some path
func (c *ChecksManager) fileExists(server *Server, path string) error {
	command := fmt.Sprintf("ls %s", path)
	err := RunSSHCommand(server.User, server.Host, command)
	return err
}

// processRunning checks if some process is running on the server
func (c *ChecksManager) processRunning(server *Server, process string) error {
	command := fmt.Sprintf("ps -A | grep %s", process)
	err := RunSSHCommand(server.User, server.Host, command)
	return err
}
