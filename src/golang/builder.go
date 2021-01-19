package golang

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Builder struct {
	items map[string]map[string]string
}

func (b *Builder) Init(items map[string]map[string]string) {
	b.items = items
}

func (b *Builder) Build(configuration string) error {
	if err := checkConfiguration(configuration); err != nil {
		return err
	}
	// check the golang file with all dependencies is exist
	wd, _ := os.Getwd()
	folderPath := filepath.Join(wd, configuration)
	filePath := filepath.Join(folderPath, depsFileName)
	if _, err := os.Stat(filePath); err != nil {
		return fmt.Errorf("\"%s\" does not exist. Please use a \"generate\" command to create it.", filePath)
	}
	g := Generator{
		b.items,
	}
	// generate a golang main file if it is missing
	filePath = filepath.Join(folderPath, mainFileName)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if err := g.generateMainFile(configuration); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	// build the application
	names := strings.Split(wd, "\\")
	return goBuild(folderPath, names[len(names)-1]+".exe")
}

func (b *Builder) Clean(configuration string) error {
	if err := checkConfiguration(configuration); err != nil {
		return err
	}
	// check the golang file with all dependencies is exist
	wd, _ := os.Getwd()
	folderPath := filepath.Join(wd, configuration)
	if _, err := os.Stat(folderPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return goClean(folderPath)
}
