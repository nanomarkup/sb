package smodule

import (
	"os"
	"strings"
)

func GetFileName(name string) string {
	if strings.HasSuffix(name, modExt) {
		return name
	} else {
		return name + modExt
	}
}

func IsExist(name string) bool {
	_, err := os.Stat(GetFileName(name))
	return err == nil
}
