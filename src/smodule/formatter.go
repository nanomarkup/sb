package smodule

import (
	"bytes"
	"fmt"
	"sort"
)

func (f *Formatter) Item(name string, deps map[string]string) string {
	if deps == nil {
		return ""
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf(attrs.itemFmt, name))
	// sort dependency items
	depsSorted := make([]string, 0, len(deps))
	for dep := range deps {
		depsSorted = append(depsSorted, dep)
	}
	sort.Strings(depsSorted)
	// add dependency items
	for _, dep := range depsSorted {
		res.WriteString(fmt.Sprintf("\t"+attrs.depFmt, dep, deps[dep]))
	}
	res.WriteString(")\n")
	return res.String()
}

func (f *Formatter) String(module Reader) string {
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf(attrs.useFmt, module.Lang()))
	// sort items
	items := module.Items()
	sorted := make([]string, 0, len(items))
	for item := range items {
		sorted = append(sorted, item)
	}
	sort.Strings(sorted)
	// add items
	for _, item := range sorted {
		res.WriteString("\n" + f.Item(item, items[item]))
	}
	return res.String()
}
