package models

type Author struct {
	name string
}

type Paper struct {
	title          string
	published_date string
	link           string
	publisher      Author
}
