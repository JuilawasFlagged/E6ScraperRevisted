package scraper

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
)

func init() {
	Mods = append(Mods, Mod{
		Name: "gelbooru",
		Func: gelbooru,
	})
}

func gelbooru(tags, page string, client *http.Client) ([]Post, error) {
	url := "https://gelbooru.com/index.php?page=dapi&s=post&q=index&limit=100&json=1"
	if tags != "" {
		url += "&tags=" + tags
	}
	if page != "" {
		url += "&pid=" + page
	}

	res, err := Request(url, client)
	if err != nil {
		return nil, err
	}

	var jsonposts gelBooruJsonPosts
	err = json.NewDecoder(res.Body).Decode(&jsonposts)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, jsonpost := range jsonposts.Posts {
		if jsonpost.FileURL == "" || jsonpost.FileMD5 == "" {
			continue
		}

		ext := filepath.Ext(jsonpost.FileURL)
		ext = strings.Replace(ext, ".", "", 1)
		posts = append(posts, Post{
			ID: jsonpost.ID,
			File: File{
				URL:       jsonpost.FileURL,
				MD5:       jsonpost.FileMD5,
				Extension: ext,
			},
		})
	}

	return posts, nil
}
