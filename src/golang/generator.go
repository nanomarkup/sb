package golang

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sapplications/sbuilder/src/smod"
)

type Generator struct {
	Configuration string
}

func (g *Generator) Generate(config *smod.ConfigFile) error {
	var err error
	useCurrentConfig := g.Configuration == ""
	if g.Configuration, err = check(g.Configuration, config); err != nil {
		return err
	}
	if _, err := g.entryPoint(config); err != nil {
		return err
	}
	if !useCurrentConfig {
		// delete the configuration golang file
		if _, err := os.Stat(configFileName); err == nil {
			if err := os.Remove(configFileName); err != nil {
				return err
			}
		}
	}
	// generate a golang file and save all dependencies
	if err := g.generateDepsFile(); err != nil {
		return err
	}
	// generate a golang file and save current configuration
	return g.generateConfigFile()
}

func (g *Generator) entryPoint(config *smod.ConfigFile) (string, error) {
	// read the main item
	main := config.Items["main"]
	if main == nil {
		return "", fmt.Errorf("The main item is not found")
	}
	// read the configuration
	entry, found := main[g.Configuration]
	if !found {
		return "", fmt.Errorf("The selected \"%s\" configuration is not found", g.Configuration)
	}
	return entry, nil
}

func (g *Generator) generateMainFile() error {
	file, err := os.Create(mainFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString("package main\n\n")
	writer.WriteString("func main() {\n")
	writer.WriteString("\tExecute()\n")
	writer.WriteString("}\n")
	return nil
}

func (g *Generator) generateDepsFile() error {
	wd, _ := os.Getwd()
	root := filepath.Join(wd, g.Configuration)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		os.Mkdir(root, os.ModePerm)
	}
	file, err := os.Create(filepath.Join(root, depsFileName))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString(fmt.Sprintf("package %s\n\n", g.Configuration))
	writer.WriteString("func Execute() {\n")
	writer.WriteString("}\n")
	return nil
}

func (g *Generator) generateConfigFile() error {
	// get package paths
	var out bytes.Buffer
	var paths []string
	cmd := exec.Command("go", "list")
	cmd.Stdout = &out
	if err := cmd.Run(); err == nil {
		paths = strings.Split(string(out.Bytes()), "\n")
	}
	if len(paths) > 0 && strings.HasPrefix(paths[0], "_") {
		// use a local import path
		paths = []string{"."}
	}
	// create a golang config file
	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString("package main\n\n")
	if len(paths) > 0 {
		writer.WriteString(fmt.Sprintf("import \"%s/%s\"\n\n", paths[0], g.Configuration))
	}
	writer.WriteString(fmt.Sprintf("const Configuration = \"%s\"\n\n", g.Configuration))
	writer.WriteString("func Execute() {\n")
	writer.WriteString(fmt.Sprintf("\t%s.Execute()\n", g.Configuration))
	writer.WriteString("}\n")
	return nil
}
