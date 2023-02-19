// Copyright 2023 Vitalii Noha vitalii.noga@gmail.com. All rights reserved.

package sb

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func (h *ModHelper) Apps() ([]string, error) {
	mod, err := h.Manager.ReadAll()
	if err != nil {
		return nil, err
	}
	var item = mod.Items()[AppsItemName]
	if item == nil {
		return nil, fmt.Errorf(ItemIsMissingF, AppsItemName)
	} else {
		return maps.Keys(item), nil
	}
}
