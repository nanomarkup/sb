package smod

import "strings"

func split(line string) []string {
	var res []string
	its := strings.Split(line, " ")
	add := true
	ind := -1
	for _, it := range its {
		if add {
			res = append(res, it)
			ind++
			if len(it) > 0 && it[0] == '"' {
				add = false
			}
		} else {
			res[ind] = res[ind] + " " + it
			if len(it) > 0 && it[len(it)-1] == '"' {
				add = true
			}
		}
	}
	return res
}
