package smodule

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sapplications/sbuilder/src/services/smodule"
)

func split(line string) []string {
	var res []string
	its := strings.Split(line, " ")
	add := true
	ind := -1
	for _, it := range its {
		if add {
			res = append(res, it)
			ind++
			if len(it) > 0 && it[0] == '"' {
				add = false
			}
		} else {
			res[ind] = res[ind] + " " + it
			if len(it) > 0 && it[len(it)-1] == '"' {
				add = true
			}
		}
	}
	return res
}

func loadModule(filePath string) (*Module, error) {
	mod := Module{}
	mod.items = map[string]map[string]string{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var item string
	var line string
	var slice []string
	var index = 1
	var length int
	var bracketOpened = false
	trimChars := "\t \n \r"
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		// process the line
		line = strings.Trim(line, trimChars)
		if line != "" {
			slice = split(line)
			if len(slice) == 0 {
				continue
			}
			for i, s := range slice {
				slice[i] = strings.Trim(s, trimChars)
			}
			if index == 1 {
				// check and initialize language
				if len(slice) != 2 {
					return nil, fmt.Errorf("cannot parse the first token of " + filePath)
				} else if slice[0] != attrs.module {
					return nil, fmt.Errorf("the first token should be \"%s\"", attrs.module)
				}
				mod.lang = slice[1]
			} else {
				// process items
				length = len(slice)
				if bracketOpened {
					// add new dependency item
					if length == 1 && slice[0] == ")" {
						item = ""
						bracketOpened = false
					} else if length != 2 {
						return nil, fmt.Errorf("cannot parse the dependency token of " + filePath)
					} else {
						mod.items[item][slice[0]] = slice[1]
					}
				} else {
					// add new item
					if (length == 1) && (slice[0] == "(") && (item != "") {
						bracketOpened = true
					} else if length < 2 {
						return nil, fmt.Errorf("cannot parse the item token of " + filePath)
					} else if (slice[1] != "require") && (slice[1] != "require(") {
						return nil, fmt.Errorf("invalid token")
					} else {
						item = slice[0]
						mod.items[item] = map[string]string{}
						if (slice[1] == "require(") || (length > 2 && strings.TrimSuffix(slice[2], "\n") == "(") {
							bracketOpened = true
						}
					}
				}
			}
			index++
		}
		// check the EOF
		if err != nil {
			break
		}
	}
	return &mod, nil
}

func loadAll(language string) (smodule.ReadWriter, error) {
	// read and check all modules in the working directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}
	modLang := ""
	modItems := map[string]map[string]string{}
	modFound := false
	var mod *Module
	for _, f := range files {
		fname := f.Name()
		if filepath.Ext(fname) != ".sb" {
			continue
		}
		modFound = true
		// load module
		if mod, err = loadModule(fname); err != nil {
			return nil, err
		}
		// validate the loaded module
		if language != "" && language != mod.lang {
			// skip the loaded module if the language is not right
			continue
		}
		if modLang == "" {
			modLang = mod.lang
		}
		if modLang != mod.lang {
			return nil, fmt.Errorf("the language of \"%s\" module do not match other modules", fname)
		}
		// populate items
		for name, data := range mod.items {
			if _, found := modItems[name]; found {
				return nil, fmt.Errorf("\"%s\" item of \"%s\" module already exists", name, fname)
			}
			modItems[name] = data
		}
	}
	if modFound {
		return &Module{modLang, modItems}, nil
	} else {
		wd, _ := os.Getwd()
		return &Module{modLang, modItems}, fmt.Errorf(ModuleFilesMissingF, wd)
	}
}

func readAll(language string) (smodule.Reader, error) {
	return loadAll(language)
}

func saveModule(module string, info smodule.Reader) error {
	file, err := os.Create(module)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := Formatter{}
	_, err = writer.WriteString(f.String(info))
	return err
}
