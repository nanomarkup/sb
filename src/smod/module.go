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

type ConfigFile struct {
	sb    string
	lang  string
	items map[string]map[string]string
}

func (c *ConfigFile) Init(version, language string) {
	c.sb = version
	c.lang = language
	c.items = map[string]map[string]string{
		"main": map[string]string{},
	}
}

func (c *ConfigFile) Main() (map[string]string, error) {
	main := c.items["main"]
	if main == nil {
		return nil, fmt.Errorf("The main item is not found")
	} else {
		return main, nil
	}
}

func (c *ConfigFile) Sb() string {
	return c.sb
}

func (c *ConfigFile) Lang() string {
	return c.lang
}

func (c *ConfigFile) Items() map[string]map[string]string {
	return c.items
}

func (c *ConfigFile) AddItem(item string) error {
	if _, found := c.items[item]; found {
		return fmt.Errorf("\"%s\" item already exists", item)
	}
	c.items[item] = map[string]string{}
	return nil
}

func (c *ConfigFile) AddDependency(item, dependency, resolver string, update bool) error {
	curr, found := c.items[item]
	if !found {
		return fmt.Errorf("\"%s\" item does not exist", item)
	}
	if _, found := curr[dependency]; found && !update {
		return fmt.Errorf("\"%s\" already exists for \"%s\" item", dependency, item)
	}
	curr[dependency] = resolver
	return nil
}

func (c *ConfigFile) DeleteItem(item string) error {
	delete(c.items, item)
	return nil
}

func (c *ConfigFile) DeleteDependency(item, dependency string) error {
	if curr, found := c.items[item]; found {
		delete(curr, dependency)
	}
	return nil
}

func (c *ConfigFile) LoadFromFile(filePath string) error {
	c.items = map[string]map[string]string{}

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
				c.sb = slice[1]
			} else if index == 2 {
				// check and initialize lang version
				if len(slice) != 2 {
					return fmt.Errorf("cannot parse the second token of " + filePath)
				} else if slice[0] != attrs.lang {
					return fmt.Errorf("the second token should be \"%s\"", attrs.lang)
				}
				c.lang = slice[1]
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
						c.items[item][slice[0]] = slice[1]
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
						c.items[item] = map[string]string{}
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

func (c *ConfigFile) SaveToFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, err = writer.WriteString(c.String())
	return err
}

func (c *ConfigFile) String() string {
	var res bytes.Buffer
	res.WriteString(c.Version())
	res.WriteString(c.Language())
	// sort items
	sorted := make([]string, 0, len(c.items))
	for item := range c.items {
		sorted = append(sorted, item)
	}
	sort.Strings(sorted)
	// add items
	for _, item := range sorted {
		res.WriteString("\n" + c.Item(item))
	}
	return res.String()
}

func (c *ConfigFile) Version() string {
	return fmt.Sprintf(attrs.sbFmt, c.Sb)
}

func (c *ConfigFile) Language() string {
	return fmt.Sprintf(attrs.langFmt, c.Lang)
}

func (c *ConfigFile) Item(item string) string {
	var deps = c.items[item]
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

func (c *ConfigFile) Dependency(item, dep string) string {
	if deps := c.items[item]; deps != nil {
		if res, found := deps[dep]; found {
			return fmt.Sprintf(attrs.depFmt, dep, res)
		}
	}
	return ""
}
