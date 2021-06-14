package smodule

import (
	"os"
	"strings"
)

func GetModuleName(name string) string {
	if strings.HasSuffix(name, modExt) {
		return name
	} else {
		return name + modExt
	}
}

func IsModuleExist(name string) bool {
	_, err := os.Stat(GetModuleName(name))
	return err == nil
}
