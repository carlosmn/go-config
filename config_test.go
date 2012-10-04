package config

import (
	"testing"
)

var config *Config

func loadConfig(s string, t *testing.T) *Config {
	config, err := Load("test.cfg")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	return config
}

func assertType(t *testing.T, setting *ConfigSetting, expected int) {
	if setting.Type() != expected {
		t.Errorf("Types %d and %d don't match", setting.Type(), expected)
	}
}

func assertBool(t *testing.T, setting *ConfigSetting, expected bool) {
	if v := setting.Bool(); v != expected {
		t.Errorf("Bool value mismatch, %v != %v", v, expected)
	}
}

func assertString(t *testing.T, got string, expected string) {
	if  got != expected {
		t.Errorf("String value mismatch, %v != %v", got, expected)
	}
}

func TestLookupString(t *testing.T) {
	config := loadConfig("test.cfg", t)

	setting := config.Lookup("something")
	assertType(t, setting, TYPE_STRING)

	assertString(t, setting.Name(), "something")
	assertString(t, setting.String(), "one")
}

func TestLookupBool(t *testing.T) {
	config := loadConfig("test.cfg", t)

	setting := config.Lookup("bool")
	assertType(t, setting, TYPE_BOOL)

	assertBool(t, setting, false)

	setting = config.Lookup("truthiness")
	assertType(t, setting, TYPE_BOOL)
	assertBool(t, setting, true)
}

func TestList(t *testing.T) {
	config := loadConfig("test.cfg", t)

	setting := config.Lookup("list")
	assertType(t, setting, TYPE_LIST)

	elems := setting.Slice()
	assertType(t, elems[0], TYPE_INT)
	assertType(t, elems[1], TYPE_STRING)
	assertType(t, elems[2], TYPE_BOOL)
}
