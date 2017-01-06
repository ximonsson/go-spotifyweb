package spotifyweb

// API host URL
const apiURL = "https://api.spotify.com"
const apiVersion = "v1"

// Search
const apiSearchURL = "/search"

// Artist endpoints
const (
	apiArtistsURL         = "/artists"
	apiArtistURL          = "/artists/%s"
	apiArtistAlbumURL     = "/artists/%s/albums"
	apiArtistTopTracksURL = "/artists/%s/top-tracks"
	apiRelatedArtistsURL  = "/artists/%s/related-artists"
)

// Album endpoints
const (
	apiAlbumsURL      = "/albums"
	apiAlbumURL       = apiAlbumsURL + "/%s"
	apiAlbumTracksURL = apiAlbumURL + "/tracks"
)

// Type names
const (
	TypeAlbum       = "album"
	TypeSingle      = "single"
	TypeAppearsOn   = "appears_on"
	TypeCompilation = "compilation"
	TypeArtist      = "artist"
	TypePlaylist    = "playlist"
	TypeTrack       = "track"
)

// Image media structure.
type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

// Paging response object from spotify.
// This object is supposed to extended with the correct
// Items field.
type paging struct {
	Href     string `json:"href"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Offset   int    `json:"offset"`
	Total    int    `json:"total"`
}
