package golang

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/smod"
)

type Builder struct {
	ModuleName    string
	Configuration string
}

func (b *Builder) Build(config *smod.ConfigFile) error {
	var err error
	if b.Configuration, err = check(b.Configuration, config); err != nil {
		return err
	}
	// check the golang file with all dependencies is exist
	wd, _ := os.Getwd()
	folderPath := filepath.Join(wd, b.Configuration)
	filePath := filepath.Join(folderPath, depsFileName)
	if _, err = os.Stat(filePath); err != nil {
		return fmt.Errorf("\"%s\" does not exist. Please use a \"generate\" command to create it.", filePath)
	}
	g := Generator{
		b.ModuleName,
		b.Configuration,
	}
	// generate a golang main file if it is missing
	filePath = filepath.Join(folderPath, mainFileName)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if err := g.generateMainFile(); err != nil {
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

func (b *Builder) Clean(config *smod.ConfigFile) error {
	var err error
	if b.Configuration, err = check(b.Configuration, config); err != nil {
		return err
	}
	defer cli.Recover()
	// check the golang file with all dependencies is exist
	wd, _ := os.Getwd()
	folderPath := filepath.Join(wd, b.Configuration)
	return goClean(folderPath)
}
