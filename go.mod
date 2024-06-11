module app

go 1.22.4

require (
	github.com/a-h/templ v0.2.648
	view v0.0.0
)

require (
	structure v0.0.0 // indirect
	models v0.0.0 // indirect
	scripts v0.0.0 // indirect
)

replace view => ./view
replace structure => ./struct

replace models => ./view/models

replace scripts => ./view/scripts
