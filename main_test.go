package main

import (
	"testing"

	"gopkg.in/gcfg.v1"
)

func TestLogin(t *testing.T) {
	err := gcfg.ReadFileInto(&config, "config.gcfg")
	check(err, "error reading config.gcfg")
	if ldapLogin("lookup", "uplook") == false {
		t.Fail()
	}
}
