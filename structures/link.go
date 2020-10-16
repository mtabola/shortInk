package structures

type Link struct {
	LinkId    int    `json:"LinkId"`
	FullLink  string `json:"FullLink"`
	ShortLink string `json:"ShortLink"`
}

func NewLink(id int, fullLink string, shortLink string) *Link {
	return &Link{id, fullLink, shortLink}
}