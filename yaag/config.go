package yaag

import "github.com/manGoweb/yaag/yaag/models"

type DocFn func(spec *models.Spec)

type Config struct {
	On bool

	BaseUrls map[string]string

	DocTitle string
	DocPath  string

	PostProcessor DocFn
}
