module app

go 1.22.4

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/a-h/templ v0.2.707
	structure v0.0.0
	view v0.0.0
)

require scripts v0.0.0 // indirect

replace view => ./view

replace structure => ./struct

replace scripts => ./view/scripts
