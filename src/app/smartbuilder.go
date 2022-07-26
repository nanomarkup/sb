// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package app

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Create creates an application by generating smart application unit (.sa file).
func (b *SmartBuilder) Create(application string) error {
	return nil
}

// Generate generates smart builder unit (.sb) using smart application unit.
func (b *SmartBuilder) Generate(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("generating \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	b.logTrace(fmt.Sprintf("checking \"%s\" application", application))
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case langs.Go:
		client, raw, err := b.newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(builder)
		sources := mod.Items()
		b.logTrace(fmt.Sprintf("generating \"%s\" application using sgo plugin", application))
		if err := builder.Generate(application, &sources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

// Build builds an application using the generated items.
func (b *SmartBuilder) Build(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("building \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case langs.Go:
		client, raw, err := b.newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(builder)
		sources := mod.Items()
		if err := builder.Build(application, &sources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

// Clean removes generated/compiled files.
func (b *SmartBuilder) Clean(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("cleaning \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case langs.Go:
		client, raw, err := b.newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(builder)
		sources := mod.Items()
		if err := builder.Clean(application, &sources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

// Run runs the application.
func (b *SmartBuilder) Run(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("running \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// run an application
	folder, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	folder = filepath.Join(folder, application)
	application = fmt.Sprintf("%s.exe", application)
	if _, err = os.Stat(filepath.Join(folder, application)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("the system cannot find the \"%s\" application", application)
		} else {
			return err
		}
	}
	wd, _ := os.Getwd()
	if err = os.Chdir(folder); err != nil {
		return err
	}
	cmd := exec.Command(application)
	output, err := cmd.Output()
	if err == nil {
		fmt.Print(string(output))
	}
	os.Chdir(wd)
	return err
}

// Version displays a version of the application.
func (b *SmartBuilder) Version() string {
	return AppVersionString
}

// Init creates a main.sb module and initialize it with the main item.
// If the main item is exist then do nothing.
func (b *SmartBuilder) Init(lang string) error {
	b.logInfo(fmt.Sprintf("initializing \"%s\" language", lang))
	b.ModManager.SetLogger(b.Logger)
	if _, found := suppLangs[lang]; found {
		return b.ModManager.Init(DefaultModuleName, lang)
	} else {
		return fmt.Errorf(LanguageIsNotSupportedF, lang)
	}
}

// ReadAll loads modules.
func (b *SmartBuilder) ReadAll(lang string) (ModReader, error) {
	defer handleError()
	b.logInfo(fmt.Sprintf("reading \"%s\" language", lang))
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(lang)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

// AddItem adds an item to the module.
func (b *SmartBuilder) AddItem(module, item string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("adding \"%s\" item to \"%s\" module", item, module))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.AddItem(module, item)
}

// AddDependency adds a dependency to the item.
func (b *SmartBuilder) AddDependency(item, dependency, resolver string, update bool) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("adding \"%s\" dependency to \"%s\" item", dependency, item))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.AddDependency(item, dependency, resolver, update)
}

// DeleteItem deletes the item from the module.
func (b *SmartBuilder) DeleteItem(item string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("deleting \"%s\" item", item))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.DeleteItem(item)
}

// DeleteDependency deletes the dependency from the item.
func (b *SmartBuilder) DeleteDependency(item, dependency string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("deleting \"%s\" dependency from \"%s\" item", dependency, item))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.DeleteDependency(item, dependency)
}

func (b *SmartBuilder) newPlugin(name string) (client *plugin.Client, raw interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(ErrorMessageF, r)
			if client != nil {
				client.Kill()
			}
		}
	}()
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   name,
		Output: os.Stdout,
		Level:  hclog.Error,
	})
	pluginMap := map[string]plugin.Plugin{
		name: b.Builder.(plugin.Plugin),
	}
	client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: b.PluginHandshake,
		Plugins:         pluginMap,
		Cmd:             exec.Command(fmt.Sprintf("%s.exe", name)),
		Logger:          logger,
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, err
	}
	raw, err = rpcClient.Dispense(name)
	if err != nil {
		return nil, nil, err
	}
	return
}

func (b *SmartBuilder) checkApplication(application string, reader ModReader) (string, error) {
	// check language
	if _, found := suppLangs[reader.Lang()]; !found {
		return "", fmt.Errorf("the current \"%s\" language is not supported", reader.Lang())
	}
	// read the main item
	main, err := reader.Main()
	if err != nil {
		return "", err
	}
	// check the number of existing applications
	if len(main) == 0 {
		return "", errors.New(ApplicationIsMissing)
	}
	// read the current application if it is not specified and only one is exist
	if application == "" {
		if len(main) != 1 {
			return "", fmt.Errorf("the application is not specified")
		}
		// select the existing application
		for key := range main {
			application = key
		}
	}
	// check the application is exist
	if _, found := main[application]; !found && application != "" {
		return "", fmt.Errorf("the selected \"%s\" application is not found", application)
	}
	return application, nil
}

func (b *SmartBuilder) logTrace(message string) {
	if b.Logger != nil {
		b.Logger.Trace(message)
	}
}

func (b *SmartBuilder) logDebug(message string) {
	if b.Logger != nil {
		b.Logger.Debug(message)
	}
}

func (b *SmartBuilder) logInfo(message string) {
	if b.Logger != nil {
		b.Logger.Info(message)
	}
}

func (b *SmartBuilder) logWarn(message string) {
	if b.Logger != nil {
		b.Logger.Warn(message)
	}
}

func (b *SmartBuilder) logError(message string) {
	if b.Logger != nil {
		b.Logger.Error(message)
	}
}
