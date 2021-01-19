package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services"
)

type SmartBuilder struct {
	Module      services.Module
	GoBuilder   services.Builder
	GoGenerator services.Generator
}

func (sb *SmartBuilder) Generate(application string) error {
	defer common.Recover()
	// load and check application
	common.Check(sb.Module.LoadFromFile(ModFileName))
	application, err := sb.checkApplication(application)
	if err != nil {
		return err
	}
	// process application
	switch sb.Module.Lang() {
	case Langs.Go:
		sb.GoGenerator.Init(sb.Module.Items())
		if err := sb.GoGenerator.Generate(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported")
	}
	return nil
}

func (sb *SmartBuilder) Build(application string) error {
	defer common.Recover()
	// load and check application
	common.Check(sb.Module.LoadFromFile(ModFileName))
	application, err := sb.checkApplication(application)
	if err != nil {
		return err
	}
	// process application
	switch sb.Module.Lang() {
	case Langs.Go:
		sb.GoBuilder.Init(sb.Module.Items())
		if err := sb.GoBuilder.Build(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported")
	}
	return nil
}

func (sb *SmartBuilder) Clean(application string) error {
	defer common.Recover()
	// load and check application
	common.Check(sb.Module.LoadFromFile(ModFileName))
	application, err := sb.checkApplication(application)
	if err != nil {
		return err
	}
	// process application
	switch sb.Module.Lang() {
	case Langs.Go:
		// remove the built files
		sb.GoBuilder.Init(sb.Module.Items())
		if err := sb.GoBuilder.Clean(application); err != nil {
			return err
		}
		// remove the generated files
		sb.GoGenerator.Init(sb.Module.Items())
		if err := sb.GoGenerator.Clean(application); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported")
	}
	return nil
}

func (sb *SmartBuilder) Version() string {
	return AppVersion
}

func (sb *SmartBuilder) checkApplication(application string) (string, error) {
	// check version
	if _, found := versions[sb.Module.Sb()]; !found {
		return "", fmt.Errorf("The current \"%s\" version is not supported", sb.Module.Sb())
	}
	// check language
	if _, found := suppLangs[sb.Module.Lang()]; !found {
		return "", fmt.Errorf("The current \"%s\" language is not supported", sb.Module.Lang())
	}
	// read the main item
	main, err := sb.Module.Main()
	if err != nil {
		return "", err
	}
	// check the number of existing applications
	if len(main) == 0 {
		return "", fmt.Errorf("Does not found any application in the main")
	}
	// read the current application if it is not specified and only one is exist
	if application == "" {
		if len(main) != 1 {
			return "", fmt.Errorf("The application is not specified")
		}
		// select the existing application
		for key := range main {
			application = key
		}
	}
	// check the application is exist
	if _, found := main[application]; !found && application != "" {
		return "", fmt.Errorf("The selected \"%s\" application is not found", application)
	}
	return application, nil
}
