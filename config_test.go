package config

import (
	"os"
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	config, err := Load("test.cfg")
	defer config.Free()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		t.FailNow()
	}
}

func TestLookupString(t *testing.T) {
	config, err := Load("test.cfg")
	defer config.Free()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		t.FailNow()
	}

	setting := config.Lookup("something")
	if setting.Type() != TYPE_STRING {
		t.FailNow()
	}

	if setting.Name() != "something" {
		t.FailNow()
	}

	if setting.String() != "one" {
		fmt.Println(setting.String())
		t.FailNow()
	}
}