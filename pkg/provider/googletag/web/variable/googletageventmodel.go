package variable

import (
	"github.com/foomo/sesamy-cli/pkg/tagmanager/web/variable"
	"google.golang.org/api/tagmanager/v2"
)

func GoogleTagEventModelName(v string) string {
	return variable.DataLayerVariableName("eventModel." + v)
}

func NewGoogleTagEventModel(v string) *tagmanager.Variable {
	return variable.NewDataLayerVariable("eventModel." + v)
}
