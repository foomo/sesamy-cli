package utm

import (
	"context"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/foomo/sesamy-cli/pkg/provider/googleconsent"
	googleconsentvariable "github.com/foomo/sesamy-cli/pkg/provider/googleconsent/server/variable"
	containertag "github.com/foomo/sesamy-cli/pkg/provider/utm/server/tag"
	containertemplate "github.com/foomo/sesamy-cli/pkg/provider/utm/server/template"
	"github.com/foomo/sesamy-cli/pkg/provider/utm/server/trigger"
	containervariable "github.com/foomo/sesamy-cli/pkg/provider/utm/server/variable"
	"github.com/foomo/sesamy-cli/pkg/tagmanager"
	commontemplate "github.com/foomo/sesamy-cli/pkg/tagmanager/common/template"
	"github.com/pkg/errors"
)

func Server(ctx context.Context, tm *tagmanager.TagManager, cfg config.Utm) error {
	folder, err := tm.UpsertFolder(ctx, "Sesamy - "+Name)
	if err != nil {
		return err
	}

	attributionTemplate, err := commontemplate.Resolve(cfg.Templates.AttributionVariable, NameUtmAttributionVariableTemplate, containertemplate.NewUtmAttribution(NameUtmAttributionVariableTemplate))
	if err != nil {
		return err
	}

	variableTemplate, err := tm.UpsertCustomTemplate(ctx, attributionTemplate)
	if err != nil {
		return err
	}

	if _, err := tm.UpsertVariable(ctx, folder, containervariable.NewUtmAttribution(NameUtmAttributionVariable, variableTemplate)); err != nil {
		return err
	}

	cookieWriterTemplate, err := commontemplate.Resolve(cfg.Templates.CookieWriterTag, NameUtmAttributionCookieWriterTemplate, containertemplate.NewUtmAttributionCookieWriter(NameUtmAttributionCookieWriterTemplate))
	if err != nil {
		return err
	}

	tagTemplate, err := tm.UpsertCustomTemplate(ctx, cookieWriterTemplate)
	if err != nil {
		return err
	}

	var triggerOpts []trigger.UtmAllEventsOption

	if cfg.GoogleConsent.Enabled {
		if err := googleconsent.ServerEnsure(ctx, tm); err != nil {
			return err
		}

		consentVariable, err := tm.LookupVariable(ctx, googleconsentvariable.GoogleConsentModeName(cfg.GoogleConsent.Mode))
		if err != nil {
			return err
		}

		triggerOpts = append(triggerOpts, trigger.UtmAllEventsWithConsentMode(consentVariable))
	}

	allEventsTrigger, err := tm.UpsertTrigger(ctx, folder, trigger.NewUtmAllEvents(NameUtmAttributionCookieWriterTrigger, triggerOpts...))
	if err != nil {
		return errors.Wrap(err, "failed to upsert event trigger: "+NameUtmAttributionCookieWriterTrigger)
	}

	if _, err := tm.UpsertTag(ctx, folder, containertag.NewUtmAttributionCookieWriter(NameUtmAttributionCookieWriterTag, tagTemplate, allEventsTrigger)); err != nil {
		return err
	}

	return nil
}
