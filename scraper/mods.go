package scraper

import (
	"net/http"
)

type ModFunc func(tags, page string, client *http.Client) ([]Post, error)
type Mod struct {
	Name string
	Func ModFunc
}

var Mods []Mod
