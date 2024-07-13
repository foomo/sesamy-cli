package cmd

import (
	"slices"

	"github.com/spf13/viper"
)

func Tag(name string) bool {
	tags := viper.GetStringSlice("tags")
	if len(tags) == 0 {
		return true
	}
	if len(tags) > 0 && !slices.Contains(tags, "-"+name) && slices.Contains(tags, name) {
		return true
	}
	return false
}
