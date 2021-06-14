package smodule

import (
	"fmt"
	"os"
	"strings"

	"github.com/sapplications/sbuilder/src/services/smodule"
)

func GetModuleName(fileName string) string {
	if strings.HasSuffix(fileName, moduleExt) {
		return fileName[0 : len(fileName)-len(moduleExt)]
	} else {
		return fileName
	}
}

func GetModuleFileName(name string) string {
	if strings.HasSuffix(name, moduleExt) {
		return name
	} else {
		return name + moduleExt
	}
}

func IsModuleExists(name string) bool {
	_, err := os.Stat(GetModuleFileName(name))
	return err == nil
}

func IsItemExists(lang, item string) (bool, smodule.ModuleName) {
	wd, _ := os.Getwd()
	mods, err := loadModules(lang)
	if (err != nil) && (err.Error() != fmt.Sprintf(ModuleFilesMissingF, wd)) {
		return false, ""
	}
	for _, m := range mods {
		if _, found := m.items[item]; found {
			return true, m.name
		}
	}
	return false, ""
}
