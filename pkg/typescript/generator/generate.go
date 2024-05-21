package generator

import (
	"log/slog"
	"strings"

	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/code"
)

type (
	Options struct {
		PackageNameReplacer *strings.Replacer
	}
	Option func(*Options)
)

func Generate(l *slog.Logger, ctpl *contemplate.Contemplate, opts ...Option) (map[string]*code.File, error) {
	l.Info("👷 generating typescript code")

	o := Options{
		PackageNameReplacer: strings.NewReplacer(".", "_", "/", "_", "-", "_"),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&o)
		}
	}

	ret := make(map[string]*code.File, len(ctpl.Packages))
	for path, pkg := range ctpl.Packages {
		l.Debug("👷 adding package", "name", pkg.Name(), "path", pkg.Path())
		inst := NewPackage(l, ctpl, pkg, o.PackageNameReplacer)
		if err := inst.Generate(); err != nil {
			return nil, err
		}
		if len(inst.Code().Body().Parts) > 0 {
			ret[o.PackageNameReplacer.Replace(path)+".ts"] = inst.Code()
		}
	}
	return ret, nil
}
