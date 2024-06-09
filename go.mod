module app

go 1.22.4

require (
	github.com/a-h/templ v0.2.648
	models v0.0.0
    scripts v0.0.0
	view v0.0.0
)

replace view => ./view

replace models => ./view/models
replace scripts => ./view/scripts
