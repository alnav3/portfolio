module app

go 1.22.4

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/a-h/templ v0.2.707
	structure v0.0.0
	view v0.0.0
)

require github.com/gorilla/securecookie v1.1.2 // indirect

require (
	github.com/gorilla/sessions v1.3.0
)

replace view => ./view

replace structure => ./struct
