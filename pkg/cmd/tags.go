package cmd

import (
	"slices"
)

func Tag(name string, tags []string) bool {
	if len(tags) == 0 {
		return true
	}
	if len(tags) > 0 && !slices.Contains(tags, "-"+name) && slices.Contains(tags, name) {
		return true
	}
	return false
}
