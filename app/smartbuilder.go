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
	mod, err := b.ModManager.ReadAll(ModKind.SB)
	if err != nil {
		return err
	}
	b.logTrace(fmt.Sprintf("checking \"%s\" application", application))
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	info, err := mod.App(application)
	if err != nil {
		return err
	}
	coder, found := info[coderAttrName]
	if !found {
		return fmt.Errorf(AttrIsMissingF, coderAttrName, application)
	}
	// process application
	client, raw, err := b.newPlugin(coder)
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
	return nil
}

// Build builds an application using the generated items.
func (b *SmartBuilder) Build(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("building \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(ModKind.SB)
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	info, err := mod.App(application)
	if err != nil {
		return err
	}
	coder, found := info[coderAttrName]
	if !found {
		return fmt.Errorf(AttrIsMissingF, coderAttrName, application)
	}
	// process application
	client, raw, err := b.newPlugin(coder)
	if err != nil {
		return err
	}
	defer client.Kill()
	builder := raw.(builder)
	b.logTrace(fmt.Sprintf("generating \"%s\" application using sgo plugin", application))
	if err := builder.Build(application); err != nil {
		return err
	}
	return nil
}

// Clean removes generated/compiled files.
func (b *SmartBuilder) Clean(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("cleaning \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(ModKind.SB)
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	info, err := mod.App(application)
	if err != nil {
		return err
	}
	coder, found := info[coderAttrName]
	if !found {
		return fmt.Errorf(AttrIsMissingF, coderAttrName, application)
	}
	// process application
	client, raw, err := b.newPlugin(coder)
	if err != nil {
		return err
	}
	defer client.Kill()
	builder := raw.(builder)
	sources := mod.Items()
	if err := builder.Clean(application, &sources); err != nil {
		return err
	}
	return nil
}

// Run runs the application.
func (b *SmartBuilder) Run(application string) error {
	defer handleError()
	b.logInfo(fmt.Sprintf("running \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(ModKind.SB)
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
			return fmt.Errorf(AppIsMissingInSystemF, application)
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

// Init creates a apps.sb module and initialize it with the apps item.
// If the apps item is exist then do nothing.
func (b *SmartBuilder) Init() error {
	b.logInfo("initializing module")
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.Init(DefaultModuleName, ModKind.SB)
}

// ReadAll loads modules.
func (b *SmartBuilder) ReadAll(kind string) (ModReader, error) {
	defer handleError()
	b.logInfo(fmt.Sprintf("reading \"%s\" modules", kind))
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(kind)
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
	// read the current application if it is not specified and only one is exist
	if application == "" {
		apps, err := reader.Apps()
		if err != nil {
			return "", err
		}
		// check the number of existing applications
		if len(apps) == 0 {
			return "", errors.New(AppIsMissing)
		}
		if len(apps) != 1 {
			return "", fmt.Errorf(AppIsNotSpecified)
		}
		// select the existing application
		for key := range apps {
			application = key
		}
	}
	return application, nil
}

//lint:ignore U1000 Ignore unused function
func (b *SmartBuilder) logTrace(message string) {
	if b.Logger != nil {
		b.Logger.Trace(message)
	}
}

//lint:ignore U1000 Ignore unused function
func (b *SmartBuilder) logDebug(message string) {
	if b.Logger != nil {
		b.Logger.Debug(message)
	}
}

//lint:ignore U1000 Ignore unused function
func (b *SmartBuilder) logInfo(message string) {
	if b.Logger != nil {
		b.Logger.Info(message)
	}
}

//lint:ignore U1000 Ignore unused function
func (b *SmartBuilder) logWarn(message string) {
	if b.Logger != nil {
		b.Logger.Warn(message)
	}
}

//lint:ignore U1000 Ignore unused function
func (b *SmartBuilder) logError(message string) {
	if b.Logger != nil {
		b.Logger.Error(message)
	}
}
