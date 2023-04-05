package blog

type Blog struct {
	Id       int    `json:"id"`
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}
