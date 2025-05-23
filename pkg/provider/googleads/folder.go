package googleads

import (
	"context"

	pkgtagmanager "github.com/foomo/sesamy-cli/pkg/tagmanager"
	"google.golang.org/api/tagmanager/v2"
)

func Folder(ctx context.Context, tm *pkgtagmanager.TagManager) (*tagmanager.Folder, error) {
	return tm.UpsertFolder(ctx, "Sesamy - "+Name)
}
