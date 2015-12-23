package main

import "testing"

// TestChecksManager checks the CheckManager object loading and parsing capabilities for checks JSON file
func TestChecksManager(t *testing.T) {
	cm := ChecksManager{}
	err := cm.Load("checks.sample.json")

	if err != nil {
		t.Errorf("[TestChecksManager] ChecksManger.Load failed: %s", err)
	}

	chkParsed := cm.Checks["check_some_content"].(map[string]interface{})

	if chkParsed["type"] != "file_contains" {
		t.Errorf("[TestChecksManager] Wrong check type: %s", chkParsed["type"])
	}

	if chkParsed["path"] != "/path/to/some/file" {
		t.Errorf("[TestChecksManager] Wrong check path: %s", chkParsed["path"])
	}

	if chkParsed["content"] != "some content" {
		t.Errorf("[TestChecksManager] Wrong check content: %s", chkParsed["content"])
	}
}
