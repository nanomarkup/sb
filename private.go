// Copyright 2022 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package sb

import (
	"fmt"
)

const (
	coderAttrName string = "coder"
)

type builder interface {
	Build(app string) error
	Clean(app string, sources *map[string]map[string]string) error
	Generate(app string, sources *map[string]map[string]string) error
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}

func getApp(name string, items map[string]map[string]string) (map[string]string, error) {
	apps, err := getApps(items)
	if err != nil {
		return nil, err
	}
	// check the applicatin is exist
	if _, found := apps[name]; !found {
		return nil, fmt.Errorf(AppIsMissingF, name)
	}
	// read application data
	info, found := items[name]
	if !found {
		return nil, fmt.Errorf(ItemIsMissingF, name)
	}
	return info, nil
}

func getApps(items map[string]map[string]string) (map[string]string, error) {
	apps := items[AppsItemName]
	if apps == nil {
		return nil, fmt.Errorf(ItemIsMissingF, AppsItemName)
	} else {
		return apps, nil
	}
}

func logTrace(logger Logger, message string) {
	if logger != nil {
		logger.Trace(message)
	}
}

func logDebug(logger Logger, message string) {
	if logger != nil {
		logger.Debug(message)
	}
}

func logInfo(logger Logger, message string) {
	if logger != nil {
		logger.Info(message)
	}
}

func logWarn(logger Logger, message string) {
	if logger != nil {
		logger.Warn(message)
	}
}

func logError(logger Logger, message string) {
	if logger != nil {
		logger.Error(message)
	}
}
