package config

import "text/template"

//we should not allow to import anything in this package
//we should import this package everywhere that we need

//we can put here anything that we need "globally" in our application
type AppConfig struct {
	UseCache 	  bool 							//we use this to allow us to test in dev mode
	TemplateCache map[string]*template.Template
}