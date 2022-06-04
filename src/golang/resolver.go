package golang

import (
	"reflect"
	"strings"
)

type resolver struct {
	application string
	entryPoint  string
	items       map[string]map[string]string
}

func (r *resolver) resolve() (items, []Type, error) {
	id := ""
	list := r.getItems()
	items := map[string]bool{}
	input := []Type{}
	for _, x := range list {
		// the struct and interface types are supported only
		if x.kind != itemKind.Struct {
			continue
		}
		id = strings.TrimPrefix(x.original, "*")
		// do not process the same item again
		if _, found := items[id]; found {
			continue
		} else {
			items[id] = true
		}
		input = append(input, Type{
			Id:      id,
			Kind:    reflect.Struct,
			Name:    x.name,
			PkgPath: strings.TrimPrefix(x.path+x.pkg, "*"),
		})
	}
	all := []Type{}
	for {
		curr, err := getTypeInfo(input)
		if err != nil {
			return nil, nil, err
		}
		all = append(all, curr...)
		// get the next items to process
		input = []Type{}
		for _, x := range curr {
			// process all fields
			if x.Fields != nil {
				for _, f := range x.Fields {
					// the struct and interface typers are supported only
					if (f.Kind != reflect.Struct && f.Kind != reflect.Interface) || f.Id == "." || f.PkgPath == "" {
						continue
					}
					// do not process the same item again
					if _, found := items[f.Id]; found {
						continue
					} else {
						items[f.Id] = true
					}
					input = append(input, Type{
						Id:      f.Id,
						Kind:    f.Kind,
						Name:    f.TypeName,
						PkgPath: f.PkgPath,
					})
				}
			}
			// process all methods
			if x.Methods != nil {
				for _, m := range x.Methods {
					// input params
					for _, f := range m.In {
						// the struct and interface types are supported only
						if (f.Kind != reflect.Struct && f.Kind != reflect.Interface) || f.Id == "." || f.PkgPath == "" {
							continue
						}
						// do not process the same item again
						if _, found := items[f.Id]; found {
							continue
						} else {
							items[f.Id] = true
						}
						input = append(input, Type{
							Id:      f.Id,
							Kind:    f.Kind,
							Name:    f.TypeName,
							PkgPath: f.PkgPath,
						})
					}
					// output params
					for _, f := range m.Out {
						// the struct and interface types are supported only
						if (f.Kind != reflect.Struct && f.Kind != reflect.Interface) || f.Id == "." || f.PkgPath == "" {
							continue
						}
						// do not process the same item again
						if _, found := items[f.Id]; found {
							continue
						} else {
							items[f.Id] = true
						}
						input = append(input, Type{
							Id:      f.Id,
							Kind:    f.Kind,
							Name:    f.TypeName,
							PkgPath: f.PkgPath,
						})
					}
				}
			}
		}
		if len(input) == 0 {
			break
		}
	}
	return list, all, nil
}

func (r *resolver) getItems() (list items) {
	list = make(items)
	r.getItem(r.entryPoint, list)
	return list
}

func (r *resolver) getItem(itemName string, list items) *item {
	if it, found := list[itemName]; found {
		return &it
	}
	// parse item and add it to the list
	pkg := ""
	name := ""
	kind := itemKind.Struct
	path := ""
	pathSep := "/"
	nameSep := "."
	if strings.HasPrefix(itemName, "\"") {
		kind = itemKind.String
		name = itemName
	} else {
		// get path
		data := strings.Split(itemName, pathSep)
		dataLen := len(data)
		fullName := data[dataLen-1]
		if dataLen > 1 {
			data = data[:dataLen-1]
			path = strings.Join(data, pathSep) + pathSep
		}
		// get pkg and item
		if fullName != "" {
			data = strings.Split(fullName, nameSep)
			dataLen = len(data)
			name = data[dataLen-1]
			if dataLen > 1 {
				pkg = data[0]
			}
		}
		// check and set type of func
		if name != "" && strings.HasSuffix(name, "()") {
			kind = itemKind.Func
		}
	}
	// create an item
	it := item{
		kind,
		name,
		pkg,
		path,
		itemName,
		make(items),
	}
	// process a simple item dependencies
	simpleItemName := itemName
	if itemName[0] == '*' {
		simpleItemName = itemName[1:]
	}
	var refIt *item
	deps := r.items[simpleItemName]
	for dep, res := range deps {
		refIt = r.getItem(res, list)
		if refIt != nil {
			it.deps[dep] = *refIt
		}
	}
	// add simple and ref items to the result set and return it
	list[simpleItemName] = it
	if itemName[0] == '*' {
		list[itemName] = it
	}
	return &it
}
