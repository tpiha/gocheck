package main

import "testing"

// TestChecksManager checks the CheckManager object loading and parsing capabilities for checks JSON file
func TestChecksManager(t *testing.T) {
	cm := ChecksManager{}
	err := cm.Load("checks_test.json")

	if err != nil {
		t.Errorf("[TestChecksManager] ChecksManger.Load failed: %s", err)
	}

	chkParsed := cm.Checks["check_etc_hosts_has_8888"].(map[string]interface{})

	if chkParsed["type"] != "file_contains" {
		t.Errorf("[TestChecksManager] Wrong check type: %s", chkParsed["type"])
	}

	if chkParsed["path"] != "/home/synkee/test" {
		t.Errorf("[TestChecksManager] Wrong check path: %s", chkParsed["path"])
	}

	if chkParsed["check"] != "bla" {
		t.Errorf("[TestChecksManager] Wrong check content: %s", chkParsed["check"])
	}
}
