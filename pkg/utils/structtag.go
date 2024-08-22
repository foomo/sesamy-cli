package utils

import (
	"strings"

	"github.com/fatih/structtag"
)

func ParseStructTagName(value, key string) (string, error) {
	tags, err := structtag.Parse(strings.Trim(value, "`"))
	if err != nil {
		return "", err
	}

	tag, err := tags.Get(key)
	if err != nil {
		return "", err
	}

	if tag.Value() != "" && tag.Value() != "-" {
		return strings.Split(tag.Value(), ",")[0], nil
	}

	return "", nil
}
