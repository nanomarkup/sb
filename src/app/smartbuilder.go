package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services/sbuilder"
	"github.com/sapplications/sbuilder/src/services/smodule"
)

type SmartBuilder struct {
	Lang        func() string
	Manager     smodule.Manager
	GoBuilder   sbuilder.Builder
	GoGenerator sbuilder.Generator
}

func (b *SmartBuilder) Generate(application string) error {
	defer common.Recover()
	// load and check application
	mod, err := b.Manager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case Langs.Go:
		b.GoGenerator.Init(mod.Items())
		if err := b.GoGenerator.Generate(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

func (b *SmartBuilder) Build(application string) error {
	defer common.Recover()
	// load and check application
	mod, err := b.Manager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case Langs.Go:
		b.GoBuilder.Init(mod.Items())
		if err := b.GoBuilder.Build(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

func (b *SmartBuilder) Clean(application string) error {
	defer common.Recover()
	// load and check application
	mod, err := b.Manager.ReadAll(b.Lang())
	if err != nil {
		return err
	}
	application, err = b.checkApplication(application, mod)
	if err != nil {
		return err
	}
	// process application
	switch mod.Lang() {
	case Langs.Go:
		// remove the built files
		b.GoBuilder.Init(mod.Items())
		if err := b.GoBuilder.Clean(application); err != nil {
			return err
		}
		// remove the generated files
		b.GoGenerator.Init(mod.Items())
		if err := b.GoGenerator.Clean(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported", mod.Lang())
	}
	return nil
}

func (b *SmartBuilder) Version() string {
	return AppVersion
}

func (b *SmartBuilder) Init(lang string) error {
	return b.Manager.Init(lang)
}

func (b *SmartBuilder) ReadAll(lang string) (smodule.Reader, error) {
	defer common.Recover()
	mod, err := b.Manager.ReadAll(lang)
	if err != nil {
		return nil, err
	}
	return mod, nil
}

func (b *SmartBuilder) AddItem(module, item string) error {
	defer common.Recover()
	return b.Manager.AddItem(module, item)
}

func (b *SmartBuilder) AddDependency(module, item, dependency, resolver string, update bool) error {
	defer common.Recover()
	return b.Manager.AddDependency(module, item, dependency, resolver, update)
}

func (b *SmartBuilder) DeleteItem(module, item string) error {
	defer common.Recover()
	return b.Manager.DeleteItem(module, item)
}

func (b *SmartBuilder) DeleteDependency(module, item, dependency string) error {
	defer common.Recover()
	return b.Manager.DeleteDependency(module, item, dependency)
}

func (b *SmartBuilder) checkApplication(application string, reader smodule.Reader) (string, error) {
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
		return "", fmt.Errorf("does not found any application in the main")
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
