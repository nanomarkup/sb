package golang

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sapplications/sbuilder/src/cli"
	"github.com/sapplications/sbuilder/src/smod"
)

type Generator struct {
	ModuleName    string
	Configuration string
}

func (g *Generator) Generate(config *smod.ConfigFile) error {
	var err error
	if g.Configuration, err = check(g.Configuration, config); err != nil {
		return err
	}
	// generate a golang file and save all dependencies
	entry, err := g.entryPoint(config)
	if err != nil {
		return err
	} else {
		return g.generateDepsFile(entry, config)
	}
}

func (g *Generator) Clean() error {
	// get current configuration if it is missing
	defer cli.Recover()
	var c smod.ConfigFile
	cli.Check(c.LoadFromFile(g.ModuleName))
	if g.Configuration == "" {
		g.Configuration, _ = check(g.Configuration, &c)
	}
	if g.Configuration == "" {
		return nil
	}
	if main := c.Items["main"]; main != nil {
		if _, found := main[g.Configuration]; found {
			if dir, err := os.Getwd(); err == nil {
				folderPath := filepath.Join(dir, g.Configuration)
				// remove the main file
				filePath := filepath.Join(folderPath, mainFileName)
				if _, err := os.Stat(filePath); err == nil {
					cli.Check(os.Remove(filePath))
				}
				// remove the deps file
				filePath = filepath.Join(folderPath, depsFileName)
				if _, err := os.Stat(filePath); err == nil {
					cli.Check(os.Remove(filePath))
				}
				// remove the configuration folder if it is empty
				// if empty, _ := cli.IsDirEmpty(folderPath); empty {
				// 	cli.Check(os.Remove(folderPath))
				// }
			}
		}
	}
	return nil
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
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(wd, g.Configuration, mainFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString("package main\n\n")
	writer.WriteString(fmt.Sprintf("const Configuration = \"%s\"\n\n", g.Configuration))
	writer.WriteString("func main() {\n")
	writer.WriteString("\tExecute()\n")
	writer.WriteString("}\n")
	return nil
}

func (g *Generator) generateDepsFile(entryPoint string, config *smod.ConfigFile) error {
	// check and get info about all dependencies
	r := resolver{
		g.Configuration,
		entryPoint,
		config,
	}
	list, err := r.resolve()
	if err != nil {
		return err
	}
	code, imports := g.generateItems(entryPoint, list)
	entry, found := list[entryPoint]
	if found && entry.kind == itemKind.String {
		imports = append(imports, "fmt")
	}
	// save dependencies to a file
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
	writer.WriteString(fmt.Sprintf("package main\n\n"))
	// write the import section
	if len(imports) > 0 {
		writer.WriteString("import (\n")
		for _, v := range imports {
			writer.WriteString(fmt.Sprintf("\t\"%s\"\n", v))
		}
		writer.WriteString(")\n\n")
	}
	// write entry point
	writer.WriteString("func Execute() {\n")
	if found {
		switch entry.kind {
		case itemKind.Func:
			writer.WriteString(fmt.Sprintf("\t%s.%s\n", entry.pkg, entry.name))
		case itemKind.Struct:
			writer.WriteString(fmt.Sprintf("\tapp := Use%s%s()\n", strings.Title(entry.pkg), entry.name))
			writer.WriteString(fmt.Sprintf("\tapp.Execute()\n"))
		case itemKind.String:
			writer.WriteString(fmt.Sprintf("\tfmt.Println(%s)\n", entry.original))
		}
	}
	writer.WriteString("}\n\n")
	// write items
	if len(code) > 0 {
		for _, v := range code {
			writer.WriteString(fmt.Sprintf("%s", v))
		}
	}
	return nil
}

func (g *Generator) generateItems(entryPoint string, list items) ([]string, []string) {
	code := []string{}
	imports := map[string]bool{}
	// get all type of struct items to process
	its := map[string]bool{}
	g.getStructItems(entryPoint, list, its)
	// generate code for all type of struct items
	fullNameDefine := ""
	fullNameReturn := ""
	for i := range its {
		if it, found := list[i]; found {
			switch it.kind {
			case itemKind.Func:
				imports[it.path+it.pkg] = true
			case itemKind.Struct:
				fullNameDefine = it.name
				fullNameReturn = it.name
				if len(it.path) > 0 {
					if it.path[0] == '*' {
						fullNameDefine = fmt.Sprintf("*%s.%s", it.pkg, it.name)
						fullNameReturn = fmt.Sprintf("&%s.%s", it.pkg, it.name)
						imports[it.path[1:]+it.pkg] = true
					} else {
						fullNameDefine = fmt.Sprintf("%s.%s", it.pkg, it.name)
						fullNameReturn = fullNameDefine
						imports[it.path+it.pkg] = true
					}
				}
				// create a new item and initialize it
				code = append(code, fmt.Sprintf("func Use%s%s() %s {\n", strings.Title(it.pkg), it.name, fullNameDefine))
				if len(it.deps) == 0 {
					code = append(code, fmt.Sprintf("\treturn %s{}\n", fullNameReturn))
				} else {
					code = append(code, fmt.Sprintf("\tv := %s{}\n", fullNameReturn))
					for k, v := range it.deps {
						switch v.kind {
						case itemKind.Func:
							imports[v.path+v.pkg] = true
							code = append(code, fmt.Sprintf("\tv.%s = %s.%s\n", k, v.pkg, strings.Replace(v.name, "()", "", 1)))
						case itemKind.Struct:
							code = append(code, fmt.Sprintf("\tv.%s = Use%s%s()\n", k, strings.Title(v.pkg), v.name))
						case itemKind.String:
							code = append(code, fmt.Sprintf("\tv.%s = %s\n", k, v.original))
						}
					}
					code = append(code, fmt.Sprintf("\treturn v\n"))
				}
				code = append(code, "}\n\n")
			}
		}
	}
	// map -> slice
	imp := []string{}
	for key := range imports {
		imp = append(imp, key)
	}
	return code, imp
}

func (g *Generator) getStructItems(original string, list items, result map[string]bool) {
	if result[original] {
		return
	}
	if it, found := list[original]; found && it.kind == itemKind.Struct {
		result[original] = true
		for _, v := range it.deps {
			if v.kind == itemKind.Struct {
				g.getStructItems(v.original, list, result)
			}
		}
	}
}
