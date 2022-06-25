package app

import (
	"errors"
	"fmt"

	"github.com/sapplications/sbuilder/src/plugins"
)

func (b *SmartBuilder) Generate(application string) error {
	defer handleError()
	// load and check application
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
		client, raw, err := newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(plugins.Builder)
		sources := mod.Items()
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
	// load and check application
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
		client, raw, err := newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(plugins.Builder)
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
	// load and check application
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
		client, raw, err := newPlugin("sgo")
		if err != nil {
			return err
		}
		defer client.Kill()
		builder := raw.(plugins.Builder)
		sources := mod.Items()
		if err := builder.Clean(application, &sources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

func (b *SmartBuilder) Version() string {
	return AppVersionString
}

func (b *SmartBuilder) Init(lang string) error {
	if _, found := suppLangs[lang]; found {
		return b.ModManager.Init(DefaultModuleName, lang)
	} else {
		return fmt.Errorf(LanguageIsNotSupportedF, lang)
	}
}

func (b *SmartBuilder) ReadAll(lang string) (ModReader, error) {
	defer handleError()
	mod, err := b.ModManager.ReadAll(lang)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

func (b *SmartBuilder) AddItem(module, item string) error {
	defer handleError()
	return b.ModManager.AddItem(module, item)
}

func (b *SmartBuilder) AddDependency(item, dependency, resolver string, update bool) error {
	defer handleError()
	return b.ModManager.AddDependency(item, dependency, resolver, update)
}

func (b *SmartBuilder) DeleteItem(item string) error {
	defer handleError()
	return b.ModManager.DeleteItem(item)
}

func (b *SmartBuilder) DeleteDependency(item, dependency string) error {
	defer handleError()
	return b.ModManager.DeleteDependency(item, dependency)
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
