package template

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

func NewGoogleConsentModeCheck(name string) *tagmanager.CustomTemplate {
	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: fmt.Sprintf(GoogleConsentModeCheckData, name),
		// oogleapi: Error 400: galleryReference: This field is invalid (or unsupported).
		// GalleryReference: &tagmanager.GalleryReference{
		//	Host:       "github.com",
		//	Owner:      "analytics-engineers",
		//	Repository: "gtm-server-variable-google-consent-mode-check",
		//	Signature:  "8905ba41f72b510484a3ff9dc27dabaf09c029eb1228e2d1435b5cc2e837cc8d",
		//	Version:    "a31230ca43cdadea1b97ef7dcf76b8e9f8c04725",
		// },
	}
}
