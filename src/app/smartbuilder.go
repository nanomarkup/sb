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

func (b *SmartBuilder) Generate(application string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("generating \"%s\" application", application))
	// load and check application
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	b.Logger.Trace(fmt.Sprintf("checking \"%s\" application", application))
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
		b.Logger.Trace(fmt.Sprintf("generating \"%s\" application using sgo plugin", application))
		if err := builder.Generate(application, &sources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

func (b *SmartBuilder) Build(application string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("building \"%s\" application", application))
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

func (b *SmartBuilder) Clean(application string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("cleaning \"%s\" application", application))
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

func (b *SmartBuilder) Run(application string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("running \"%s\" application", application))
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

func (b *SmartBuilder) Version() string {
	return AppVersionString
}

func (b *SmartBuilder) Init(lang string) error {
	b.Logger.Info(fmt.Sprintf("initializing \"%s\" language", lang))
	b.ModManager.SetLogger(b.Logger)
	if _, found := suppLangs[lang]; found {
		return b.ModManager.Init(DefaultModuleName, lang)
	} else {
		return fmt.Errorf(LanguageIsNotSupportedF, lang)
	}
}

func (b *SmartBuilder) ReadAll(lang string) (ModReader, error) {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("reading \"%s\" language", lang))
	b.ModManager.SetLogger(b.Logger)
	mod, err := b.ModManager.ReadAll(lang)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

func (b *SmartBuilder) AddItem(module, item string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("adding \"%s\" item to \"%s\" module", item, module))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.AddItem(module, item)
}

func (b *SmartBuilder) AddDependency(item, dependency, resolver string, update bool) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("adding \"%s\" dependency to \"%s\" item", dependency, item))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.AddDependency(item, dependency, resolver, update)
}

func (b *SmartBuilder) DeleteItem(item string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("deleting \"%s\" item", item))
	b.ModManager.SetLogger(b.Logger)
	return b.ModManager.DeleteItem(item)
}

func (b *SmartBuilder) DeleteDependency(item, dependency string) error {
	defer handleError()
	b.Logger.Info(fmt.Sprintf("deleting \"%s\" dependency from \"%s\" item", dependency, item))
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
		HandshakeConfig: handshakeConfig,
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
