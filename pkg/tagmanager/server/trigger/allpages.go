package trigger

import (
	"google.golang.org/api/tagmanager/v2"
)

const IDAllPages = "2147479574"

var AllPages = &tagmanager.Trigger{
	TriggerId: IDAllPages,
}
