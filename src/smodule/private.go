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

func loadModule(name string) (*Module, error) {
	mod := Module{}
	mod.name = name
	mod.items = Items{}

	fileName := GetModuleFileName(name)
	file, err := os.Open(fileName)
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
					return nil, fmt.Errorf("cannot parse the first token of " + fileName)
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
						return nil, fmt.Errorf("cannot parse the dependency token of " + fileName)
					} else {
						mod.items[item][slice[0]] = slice[1]
					}
				} else {
					// add new item
					if (length == 1) && (slice[0] == "(") && (item != "") {
						bracketOpened = true
					} else if length < 2 {
						return nil, fmt.Errorf("cannot parse the item token of " + fileName)
					} else if (slice[1] != "require") && (slice[1] != "require(") {
						return nil, fmt.Errorf("invalid token")
					} else {
						item = slice[0]
						mod.items[item] = Item{}
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

func loadModules(lang string) (modules, error) {
	// read and check all modules in the working directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}
	mods := modules{}
	modLang := ""
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
		if lang != "" && lang != mod.lang {
			// skip the loaded module if the language is not the selected language
			continue
		}
		if modLang == "" {
			modLang = mod.lang
		}
		if modLang != mod.lang {
			return nil, fmt.Errorf("the language of \"%s\" module do not match other modules", fname)
		}
		// add module
		mods = append(mods, Module{name: GetModuleName(fname), lang: mod.lang, items: mod.items})
	}
	if modFound {
		return mods, nil
	} else {
		wd, _ := os.Getwd()
		return nil, fmt.Errorf(ModuleFilesMissingF, wd)
	}
}

func loadItems(mods modules) (smodule.ReadWriter, error) {
	all := Items{}
	lang := ""
	if len(mods) > 0 {
		lang = mods[0].lang
	}
	for _, m := range mods {
		// read all items and validate them
		for name, data := range m.items {
			if _, found := all[name]; found {
				return nil, fmt.Errorf("\"%s\" item of \"%s\" module already exists", name, m.name)
			}
			all[name] = data
		}
	}
	return &Module{name: "", lang: lang, items: all}, nil
}

func saveModule(module *Module) error {
	fileName := GetModuleFileName(module.name)
	exists := IsModuleExists(fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// notify about a new module has been created
	defer func() {
		if !exists {
			fmt.Printf(ModuleIsCreatedF, fileName)
		}
	}()
	// save the module
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := Formatter{}
	_, err = writer.WriteString(f.String(module))
	return err
}

func addItem(module, lang, item string) error {
	// check the item is exist
	if found, modName := IsItemExists(lang, item); found {
		return fmt.Errorf(ItemExistsF, item, modName)
	}
	// load the existing module or create a new one
	var mod *Module
	var err error
	if IsModuleExists(module) {
		if mod, err = loadModule(module); err != nil {
			return err
		}
		// check language of the selected module
		if mod.lang != lang {
			return fmt.Errorf(ModuleLanguageMismatchF, mod.lang, mod.name, lang)
		}
	} else {
		mod = &Module{name: module, lang: lang, items: Items{}}
	}
	// add the item to the selected module
	if err = mod.AddItem(item); err != nil {
		return err
	} else {
		return saveModule(mod)
	}
}

func findItem(lang, item string) (*Module, error) {
	wd, _ := os.Getwd()
	mods, err := loadModules(lang)
	if (err != nil) && (err.Error() != fmt.Sprintf(ModuleFilesMissingF, wd)) {
		return nil, err
	}
	for _, m := range mods {
		if _, found := m.items[item]; found {
			return &m, nil
		}
	}
	return nil, nil
}
