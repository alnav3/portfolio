module view

go 1.22.4

require (
	github.com/a-h/templ v0.2.543
	models v0.0.0
	scripts v0.0.0
	structure v0.0.0
)

replace structure => ../struct

replace models => ./models

replace scripts => ./scripts
