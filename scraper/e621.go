package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	Mods = append(Mods, Mod{
		Name: "e621",
		Func: e621,
	})
}

func e621(tags, page string, client *http.Client) ([]Post, error) {
	url := "https://e621.net/posts.json?limit=100"
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

	var jsonposts jsonPosts
	err = json.NewDecoder(res.Body).Decode(&jsonposts)

	if err != nil {
		return nil, err
	}

	var posts []Post

	// Genius blacklist bypass
	for _, post := range jsonposts.Posts {
		if post.File.URL == "" { // Is on default blacklist
			post.File.URL = fmt.Sprintf("https://static1.e621.net/data/%s/%s/%s.%s",
				post.File.MD5[0:2],
				post.File.MD5[2:4],
				post.File.MD5,
				post.File.Extension,
			)
		}

		posts = append(posts, post)
	}

	return posts, nil
}
