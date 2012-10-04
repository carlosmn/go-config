package config

/*
#cgo pkg-config: libconfig
#include <libconfig.h>
#include <stdlib.h>
*/
import "C"

const (
	TYPE_NONE =  C.CONFIG_TYPE_NONE
	TYPE_GROUP = C.CONFIG_TYPE_GROUP
	TYPE_INT = C.CONFIG_TYPE_INT
	TYPE_INT64 = C.CONFIG_TYPE_INT64
	TYPE_FLOAT = C.CONFIG_TYPE_FLOAT
	TYPE_STRING = C.CONFIG_TYPE_STRING
	TYPE_BOOL = C.CONFIG_TYPE_BOOL
	TYPE_ARRAY = C.CONFIG_TYPE_ARRAY
	TYPE_LIST = C.CONFIG_TYPE_LIST
)

type ConfigSetting struct {
	setting *C.config_setting_t
}

func (cs *ConfigSetting) Type() int {
	return int(cs.setting._type)
}

func (cs *ConfigSetting) Length() int {
	return int(C.config_setting_length(cs.setting))
}

func (cs *ConfigSetting) String() string {
	cstr := C.config_setting_get_string(cs.setting)
	return C.GoString(cstr)
}

func (cs *ConfigSetting) Bool() bool {
	ret := C.config_setting_get_bool(cs.setting)
	if int(ret) == 0 {
		return false
	}

	return true
}

func (cs *ConfigSetting) Int() int {
	ret := C.config_setting_get_int(cs.setting)
	return int(ret)
}

func (cs *ConfigSetting) Float() float32 {
	ret := C.config_setting_get_float(cs.setting)
	return float32(ret)
}

func (cs *ConfigSetting) Slice() []*ConfigSetting {
	l := cs.Length()
	list := make([]*ConfigSetting, l)

	for i := 0; i < l; i++ {
		setting := C.config_setting_get_elem(cs.setting, C.uint(i))
		list[i] = &ConfigSetting{setting}
	}

	return list
}

func (cs *ConfigSetting) Name() string {
	return C.GoString(cs.setting.name)
}
