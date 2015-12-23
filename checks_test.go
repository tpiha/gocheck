package main

import "testing"

// TestChecksManager checks the CheckManager object loading and parsing capabilities for checks JSON file
func TestChecksManager(t *testing.T) {
	cm := ChecksManager{}
	err := cm.Load("checks_test.json")

	if err != nil {
		t.Errorf("[TestChecksManager] ChecksManger.Load failed: %s", err)
	}

	chkParsed := cm.Checks["check1"].(map[string]interface{})

	if chkParsed["type"] != "file_contains" {
		t.Errorf("[TestChecksManager] Wrong check type: %s", chkParsed["type"])
	}

	if chkParsed["path"] != "/some/path/to/file" {
		t.Errorf("[TestChecksManager] Wrong check path: %s", chkParsed["path"])
	}

	if chkParsed["content"] != "some content" {
		t.Errorf("[TestChecksManager] Wrong check content: %s", chkParsed["content"])
	}
}
