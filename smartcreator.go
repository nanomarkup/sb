// Copyright 2023 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package app

import (
	"fmt"
	"strings"
)

// Create creates an application by generating smart application unit (.sa file).
func (c *SmartCreator) Create(application string) error {
	if application == "" {
		return fmt.Errorf(AppIsNotSpecified)
	}
	defer handleError()
	logInfo(c.Logger, fmt.Sprintf("creating \"%s\" application", application))
	reader, err := c.ModManager.ReadAll()
	if err != nil && !strings.HasPrefix(err.Error(), fmt.Sprintf(ModuleFilesMissingF, ModKind.SA)) {
		return err
	}
	items := reader.Items()
	apps, err := getApps(items)
	if err != nil && !strings.HasPrefix(err.Error(), fmt.Sprintf(ItemIsMissingF, AppsItemName)) {
		return err
	}
	if apps != nil {
		if _, ok := apps[application]; ok {
			return fmt.Errorf(AppIsExistF, application)
		}
	}
	if _, ok := items[AppsItemName]; !ok {
		err = c.ModManager.AddItem(fmt.Sprintf("%s.%s", DefaultModuleName, ModKind.SA), AppsItemName)
		if err != nil {
			return err
		}
	}
	return c.ModManager.AddDependency(AppsItemName, application, "", false)
}
