package golang

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	items map[string]map[string]string
}

func (g *Generator) Init(items map[string]map[string]string) {
	g.items = items
}

func (g *Generator) Generate(application string) error {
	if err := checkApplication(application); err != nil {
		return err
	}
	// generate a golang file and save all dependencies
	entry, err := g.entryPoint(application)
	if err != nil {
		return err
	} else {
		return g.generateDepsFile(application, entry)
	}
}

func (g *Generator) Clean(application string) error {
	if err := checkApplication(application); err != nil {
		return err
	}
	// get current application if it is missing
	if application == "" {
		return fmt.Errorf("The application is not specified")
	}
	if main, err := readMain(g.items); err == nil {
		if _, found := main[application]; found {
			if dir, err := os.Getwd(); err == nil {
				folderPath := filepath.Join(dir, application)
				// remove the main file
				filePath := filepath.Join(folderPath, mainFileName)
				if _, err := os.Stat(filePath); err == nil {
					os.Remove(filePath)
				}
				// remove the deps file
				filePath = filepath.Join(folderPath, depsFileName)
				if _, err := os.Stat(filePath); err == nil {
					os.Remove(filePath)
				}
				// remove the application folder if it is empty
				if empty, _ := isDirEmpty(folderPath); empty {
					os.Remove(folderPath)
				}
			}
		}
	}
	return nil
}

func (g *Generator) entryPoint(application string) (string, error) {
	// read the main item
	main, err := readMain(g.items)
	if err != nil {
		return "", err
	}
	// read the application
	entry, found := main[application]
	if !found {
		return "", fmt.Errorf("The selected \"%s\" application is not found", application)
	}
	return entry, nil
}

func (g *Generator) generateMainFile(application string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(wd, application, mainFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.WriteString("package main\n\n")
	writer.WriteString(fmt.Sprintf("const AppName = \"%s\"\n\n", application))
	writer.WriteString("func main() {\n")
	writer.WriteString("\tExecute()\n")
	writer.WriteString("}\n")
	return nil
}

func (g *Generator) generateDepsFile(application, entryPoint string) error {
	// check and get info about all dependencies
	r := resolver{
		application,
		entryPoint,
		g.items,
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
	root := filepath.Join(wd, application)
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
	funcName := ""
	fullNameDefine := ""
	fullNameReturn := ""
	for i := range its {
		if it, found := list[i]; found {
			switch it.kind {
			case itemKind.Func:
				imports[it.path+it.pkg] = true
			case itemKind.Struct:
				funcName = fmt.Sprintf("Use%s%s", strings.Title(it.pkg), it.name)
				fullNameDefine = it.name
				fullNameReturn = it.name
				if len(it.path) > 0 {
					if it.path[0] == '*' {
						funcName = funcName + "Ref"
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
				code = append(code, fmt.Sprintf("func %s() %s {\n", funcName, fullNameDefine))
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
							funcName = fmt.Sprintf("Use%s%s", strings.Title(v.pkg), v.name)
							if len(v.path) > 0 && v.path[0] == '*' {
								funcName = funcName + "Ref"
							}
							code = append(code, fmt.Sprintf("\tv.%s = %s()\n", k, funcName))
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
