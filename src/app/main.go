package app

import (
	"fmt"

	"github.com/sapplications/sbuilder/src/common"
	"github.com/sapplications/sbuilder/src/services"
)

type SmartBuilder struct {
	Module      services.IModule
	GoBuilder   services.IBuilder
	GoGenerator services.IGenerator
}

func (sb *SmartBuilder) Generate(configuration string) error {
	defer common.Recover()
	// load and check configuration
	common.Check(sb.Module.LoadFromFile(ModFileName))
	configuration, err := sb.checkConfiguration(configuration)
	if err != nil {
		return err
	}
	// process configuration
	switch sb.Module.Lang() {
	case Langs.Go:
		sb.GoGenerator.Init(sb.Module.Items())
		if err := sb.Generate(configuration); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported")
	}
	return nil
}

func (sb *SmartBuilder) Build(configuration string) error {
	defer common.Recover()
	// load and check configuration
	common.Check(sb.Module.LoadFromFile(ModFileName))
	configuration, err := sb.checkConfiguration(configuration)
	if err != nil {
		return err
	}
	// process configuration
	switch sb.Module.Lang() {
	case Langs.Go:
		sb.GoBuilder.Init(sb.Module.Items())
		if err := sb.GoBuilder.Build(configuration); err != nil {
			return err
		}
	default:
		return fmt.Errorf("\"%s\" language is not supported")
	}
	return nil
}

func (sb *SmartBuilder) Clean(configuration string) error {
	defer common.Recover()
	// load and check configuration
	common.Check(sb.Module.LoadFromFile(ModFileName))
	configuration, err := sb.checkConfiguration(configuration)
	if err != nil {
		return err
	}
	// process configuration
	switch sb.Module.Lang() {
	case Langs.Go:
		// remove the generated files
		sb.GoGenerator.Init(sb.Module.Items())
		if err := sb.GoGenerator.Clean(configuration); err != nil {
			return err
		}
		// remove the built files
		sb.GoBuilder.Init(sb.Module.Items())
		if err := sb.GoBuilder.Clean(configuration); err != nil {
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

func (sb *SmartBuilder) checkConfiguration(configuration string) (string, error) {
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
	// check the number of existing configurations
	if len(main) == 0 {
		return "", fmt.Errorf("Does not found any configuration in the main")
	}
	// read the current configuration if it is not specified and only one is exist
	if configuration == "" {
		if len(main) != 1 {
			return "", fmt.Errorf("The configuration is not specified")
		}
		// select the existing configuration
		for key := range main {
			configuration = key
		}
	}
	// check the configuration is exist
	if _, found := main[configuration]; !found && configuration != "" {
		return "", fmt.Errorf("The selected \"%s\" configuration is not found", configuration)
	}
	return configuration, nil
}
