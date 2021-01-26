// Package smod manages smart module
//
// Copyright © 2020 Vitalii Noha vitalii.noga@gmail.com
package smod

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var attrs = struct {
	sb      string
	sbFmt   string
	lang    string
	langFmt string
	itemFmt string
	depFmt  string
}{
	"sb",
	"sb %s\n",
	"lang",
	"lang %s\n",
	"%s require (\n",
	"%s %s\n",
}

type Module struct {
	sb    string
	lang  string
	items map[string]map[string]string
}

func (m *Module) Init(version, language string) {
	m.sb = version
	m.lang = language
	m.items = map[string]map[string]string{
		"main": map[string]string{},
	}
}

func (m *Module) Main() (map[string]string, error) {
	main := m.items["main"]
	if main == nil {
		return nil, fmt.Errorf("The main item is not found")
	} else {
		return main, nil
	}
}

func (m *Module) Sb() string {
	return m.sb
}

func (m *Module) Lang() string {
	return m.lang
}

func (m *Module) Items() map[string]map[string]string {
	return m.items
}

func (m *Module) AddItem(item string) error {
	if _, found := m.items[item]; found {
		return fmt.Errorf("\"%s\" item already exists", item)
	}
	m.items[item] = map[string]string{}
	return nil
}

func (m *Module) AddDependency(item, dependency, resolver string, update bool) error {
	curr, found := m.items[item]
	if !found {
		return fmt.Errorf("\"%s\" item does not exist", item)
	}
	if _, found := curr[dependency]; found && !update {
		return fmt.Errorf("\"%s\" already exists for \"%s\" item", dependency, item)
	}
	curr[dependency] = resolver
	return nil
}

func (m *Module) DeleteItem(item string) error {
	delete(m.items, item)
	return nil
}

func (m *Module) DeleteDependency(item, dependency string) error {
	if curr, found := m.items[item]; found {
		delete(curr, dependency)
	}
	return nil
}

func (m *Module) Load() error {
	m.sb = ""
	m.lang = ""
	m.items = map[string]map[string]string{}
	// read and check all modules in the working directory

	return nil
}

func (m *Module) loadFromFile(filePath string) error {
	m.items = map[string]map[string]string{}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var item string
	var line string
	var slice []string
	var index = 1
	var length int
	var bracketOpened = false
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		// process the line
		line = strings.Trim(line, "\t \n \r")
		if line != "" {
			slice = split(line)
			if len(slice) == 0 {
				continue
			}
			for i, s := range slice {
				slice[i] = strings.Trim(s, "\t \n \r")
			}
			if index == 1 {
				// check and initialize sb version
				if len(slice) != 2 {
					return fmt.Errorf("cannot parse the first token of " + filePath)
				} else if slice[0] != attrs.sb {
					return fmt.Errorf("the first token should be \"%s\"", attrs.sb)
				}
				m.sb = slice[1]
			} else if index == 2 {
				// check and initialize lang version
				if len(slice) != 2 {
					return fmt.Errorf("cannot parse the second token of " + filePath)
				} else if slice[0] != attrs.lang {
					return fmt.Errorf("the second token should be \"%s\"", attrs.lang)
				}
				m.lang = slice[1]
			} else {
				// process items
				length = len(slice)
				if bracketOpened {
					// add new dependency item
					if length == 1 && slice[0] == ")" {
						item = ""
						bracketOpened = false
					} else if length != 2 {
						return fmt.Errorf("cannot parse the dependency token of " + filePath)
					} else {
						m.items[item][slice[0]] = slice[1]
					}
				} else {
					// add new item
					if (length == 1) && (slice[0] == "(") && (item != "") {
						bracketOpened = true
					} else if length < 2 {
						return fmt.Errorf("cannot parse the item token of " + filePath)
					} else if (slice[1] != "require") && (slice[1] != "require(") {
						return fmt.Errorf("invalid token")
					} else {
						item = slice[0]
						m.items[item] = map[string]string{}
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
	return nil
}

func (m *Module) SaveToFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, err = writer.WriteString(m.String())
	return err
}

func (m *Module) String() string {
	var res bytes.Buffer
	res.WriteString(m.Version())
	res.WriteString(m.Language())
	// sort items
	sorted := make([]string, 0, len(m.items))
	for item := range m.items {
		sorted = append(sorted, item)
	}
	sort.Strings(sorted)
	// add items
	for _, item := range sorted {
		res.WriteString("\n" + m.Item(item))
	}
	return res.String()
}

func (m *Module) Version() string {
	return fmt.Sprintf(attrs.sbFmt, m.Sb)
}

func (m *Module) Language() string {
	return fmt.Sprintf(attrs.langFmt, m.Lang)
}

func (m *Module) Item(item string) string {
	var deps = m.items[item]
	if deps == nil {
		return ""
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf(attrs.itemFmt, item))
	// sort dependency items
	depsSorted := make([]string, 0, len(deps))
	for dep := range deps {
		depsSorted = append(depsSorted, dep)
	}
	sort.Strings(depsSorted)
	// add dependency items
	for _, dep := range depsSorted {
		res.WriteString(fmt.Sprintf("\t"+attrs.depFmt, dep, deps[dep]))
	}
	res.WriteString(")\n")
	return res.String()
}

func (m *Module) Dependency(item, dep string) string {
	if deps := m.items[item]; deps != nil {
		if res, found := deps[dep]; found {
			return fmt.Sprintf(attrs.depFmt, dep, res)
		}
	}
	return ""
}
