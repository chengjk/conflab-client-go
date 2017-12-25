package conflab

import (
	"strconv"
	"os"
)

func Get(key string) string {
	return GetWithDefault(key, "")
}
func GetWithDefault(key, def string) string {
	//env
	if s, b := os.LookupEnv(key);b{
		return s
	}

	//runtime
	for _,item:=range cache{
		if v := item[key]; v != "" {
			return v
		}
	}
	return def
}

func GetBoolean(k string) {
	GetBooleanDefault(k, false)
}

func GetBooleanDefault(k string, def bool) bool {
	v := Get(k)
	if bool, e := strconv.ParseBool(v); e == nil {
		return bool
	} else {
		return def
	}
}

func GetInt(key string) int {
	return GetIntDefault(key, -1)
}
func GetIntDefault(key string, def int) int {
	v := Get(key)
	if i, err := strconv.Atoi(v); err == nil {
		return i
	} else {
		return def
	}
}
