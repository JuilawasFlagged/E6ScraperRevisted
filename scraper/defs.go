package scraper

type jsonPosts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID   uint64 `json:"id"`
	File File   `json:"file"`
}

type File struct {
	URL       string `json:"url"`
	Extension string `json:"ext"`
	MD5       string `json:"md5"`
}

// Transformed into a Post and a File eventually
type r34JsonPost struct {
	ID      uint64 `json:"id"`
	FileURL string `json:"file_url"`
	FileMD5 string `json:"hash"`
}

// Gelbooru's API response is the same as danbooru's response, except without extension.
type gelBooruJsonPosts struct {
	Posts []danBooruJsonPost `json:"post"`
}

// Transformed into a Post and a File eventually
type danBooruJsonPost struct {
	ID      uint64 `json:"id"`
	FileExt string `json:"file_ext"`
	FileURL string `json:"file_url"`
	FileMD5 string `json:"md5"`
}
