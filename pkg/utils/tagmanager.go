package utils

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func TemplateType(template *tagmanager.CustomTemplate) string {
	if template.GalleryReference != nil && template.GalleryReference.GalleryTemplateId != "" {
		return fmt.Sprintf("cvt_%s", template.GalleryReference.GalleryTemplateId)
	}
	return fmt.Sprintf("cvt_%s_%s", template.ContainerId, template.TemplateId)
}

func TriggerIDs(triggers []*tagmanager.Trigger) []string {
	ret := make([]string, len(triggers))
	for i, trigger := range triggers {
		ret[i] = trigger.TriggerId
	}
	return ret
}
