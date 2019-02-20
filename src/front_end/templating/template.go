package templating

import "html/template"

var err error
var tpls *template.Template

// InitTemplating function is used to initialize the templates
func InitTemplating(pattern ...string) error {
	// initialize first and main pattern
	tpls, err = template.ParseGlob(pattern[0])
	if err != nil {
		return err
	}

	// if more patterns were given they will be added too
	if len(pattern) > 1 {
		for _, patt := range pattern[1:] {
			tpls, err = tpls.ParseGlob(patt)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// GetTemplates is used to get the parsed templates from other packages
func GetTemplates() *template.Template {
	return tpls
}
