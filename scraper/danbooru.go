package scraper

import (
	"encoding/json"
	"net/http"
)

func init() {
	Mods = append(Mods, Mod{
		Name: "danbooru",
		Func: danbooru,
	})
}

func danbooru(tags, page string, client *http.Client) ([]Post, error) {
	url := "https://danbooru.donmai.us/posts.json?limit=100"
	if tags != "" {
		url += "&tags=" + tags
	}
	if page != "" {
		url += "&page=" + page
	}

	res, err := Request(url, client)
	if err != nil {
		return nil, err
	}

	var jsonposts []danBooruJsonPost
	err = json.NewDecoder(res.Body).Decode(&jsonposts)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, jsonpost := range jsonposts {
		if jsonpost.FileURL == "" || jsonpost.FileMD5 == "" {
			continue
		}

		posts = append(posts, Post{
			ID: jsonpost.ID,
			File: File{
				URL:       jsonpost.FileURL,
				MD5:       jsonpost.FileMD5,
				Extension: jsonpost.FileExt,
			},
		})
	}

	return posts, nil
}
