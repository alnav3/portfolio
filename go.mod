module app

go 1.21.2

require (
    models v0.0.0
	view v0.0.0
	github.com/a-h/templ v0.2.543
)

replace view => ./view
replace models => ./view/models

