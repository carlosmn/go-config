package config

/*
#cgo pkg-config: libconfig
#include <libconfig.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type Config struct {
	config C.config_t
}

const (
	FALSE = 0
	TRUE = 1
)
	
func (c *Config) ErrorLine() int {
	return int(c.config.error_line)
}

func (c *Config) ErrorText() string {
	return C.GoString(c.config.error_text)
}

func Load (path string) (*Config, error) {
	c := new(Config)
	C.config_init(&c.config)
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	ret := C.config_read_file(&c.config, cpath)
	if ret == FALSE {
		s := fmt.Sprintf("Error loading '%s' at line %d: %s", path, c.ErrorLine(), c.ErrorText())
		c.Free()

		return nil, errors.New(s)
	}

	return c, nil
}

func (c *Config) Lookup(key string) *ConfigSetting {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	setting := C.config_lookup(&c.config, ckey)

	return &ConfigSetting{setting}
}

func (c *Config) Free() {
	C.config_destroy(&c.config)
}
