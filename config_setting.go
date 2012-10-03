package config

/*
#cgo pkg-config: libconfig
#include <libconfig.h>
#include <stdlib.h>
*/
import "C"

const (
	TYPE_NONE =  C.CONFIG_TYPE_NONE
	TYPE_STRING = C.CONFIG_TYPE_STRING
	TYPE_ARRAY = C.CONFIG_TYPE_ARRAY
)

type ConfigSetting struct {
	setting *C.config_setting_t
}

func (cs *ConfigSetting) Type() int {
	return int(cs.setting._type)
}

func (cs *ConfigSetting) String() string {
	cstr := C.config_setting_get_string(cs.setting)
	return C.GoString(cstr)
}


func (cs *ConfigSetting) Name() string {
	return C.GoString(cs.setting.name)
}
