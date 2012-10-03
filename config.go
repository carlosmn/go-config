package config

/*
#cgo pkg-config: libconfig
#include <libconfig.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
	"runtime"
)

type Config struct {
	config C.config_t
}

type ConfigError struct {
	config *Config
}

func (e ConfigError) Error() string {
	return fmt.Sprintf("Error at %s:%d: %s",
		e.config.ErrorFile(), e.config.ErrorLine(), e.config.ErrorText())
}

const (
	FALSE = 0
	TRUE = 1
)

func NewConfig() *Config {
	c := new(Config)
	C.config_init(&c.config)
	runtime.SetFinalizer(c, freeConfig)
	return c
}

func freeConfig(c *Config) {
	C.config_destroy(&c.config)
}
	
func (c *Config) ErrorLine() int {
	return int(c.config.error_line)
}

func (c *Config) ErrorText() string {
	return C.GoString(c.config.error_text)
}

func (c *Config) ErrorFile() string {
	return C.GoString(c.config.error_file)
}

func Load (path string) (*Config, error) {
	c := NewConfig()
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	ret := C.config_read_file(&c.config, cpath)
	if ret == FALSE {
		return nil, ConfigError{c}
	}

	return c, nil
}

func (c *Config) Lookup(key string) *ConfigSetting {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	setting := C.config_lookup(&c.config, ckey)

	return &ConfigSetting{setting}
}
