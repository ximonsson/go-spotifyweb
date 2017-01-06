package spotifyweb

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// An artist, full object
type Artist struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Href         string            `json:"href"`
	Type         string            `json:"type"`
	URI          string            `json:"uri"`
	Popularity   int               `json:"popularity"`
	ExternalURLs map[string]string `json:"external_urls"`
	Genres       []string          `json:"genres"`
	Images       []Image           `json:"images"`
}

// Track, full object
// @NOTE: I am scared of circular references with the Album field, but I am sure the
// API responses will not make it come to a problem
type Track struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Disc         int               `json:"disc_number"`
	Href         string            `json:"href"`
	Popularity   int               `json:"popularity"`
	TrackNumber  int               `json:"track_number"`
	URI          string            `json:"uri"`
	ExternalURLs map[string]string `json:"external_urls"`
	Artists      []Artist          `json:"artists"`
	Duration     int               `json:"duration_ms"`
	Album        Album             `json:"album"`
}

// Album, full object
type Album struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Release      string            `json:"string"`
	Genres       []string          `json:"genres"`
	Type         string            `json:"album_type"`
	Images       []Image           `json:"images"`
	Artists      []Artist          `json:"artists"`
	Href         string            `json:"href"`
	Popularity   int               `json:"popularity"`
	ExternalURLs map[string]string `json:"external_urls"`
	Tracks       []struct {
		paging
		Items []Track `json:"items"`
	} `json:"tracks"`
}

// Search result for a search on artist name
type ArtistSearchResult struct {
	Artists struct {
		paging
		Items []Artist `json:"items"`
	} `json:"artists"`
}

// SearchResult struct of recieved data for a search query towards the Spotify API.
type SearchResult struct {
	// Artists matching search query
	Artists struct {
		paging
		Items []Artist `json:"items"`
	} `json:"artists"`
	// Albums matching search query
	Albums struct {
		paging
		Items []Album `json:"items"`
	} `json:"albums"`
	// Tracks matching search query
	Tracks struct {
		paging
		Items []Track `json:"items"`
	} `json:"tracks"`
}

// Get Spotify catalog information about artists, albums, tracks or playlists that match a keyword string.
func Search(q string, types []string, limit, offset int) (SearchResult, error) {
	params := url.Values{}
	params.Set("q", q)
	params.Set("type", strings.Join(types, ","))
	if limit != -1 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if offset != -1 {
		params.Set("offset", strconv.Itoa(offset))
	}
	var res SearchResult
	e := doRequest(apiSearchURL, params, "GET", &res)
	return res, e
}

// GetArtist fetches artist data for a single artist by id.
func GetArtist(id string) (Artist, error) {
	var artist Artist
	e := doRequest(fmt.Sprintf(apiArtistURL, id), nil, "GET", &artist)
	return artist, e
}

// Get artists by list of ids.
func GetArtists(id ...string) ([]Artist, error) {
	res := struct {
		Artists []Artist `json:"artists"`
	}{
		Artists: make([]Artist, 0),
	}
	params := url.Values{}
	params.Set("ids", strings.Join(id, ","))
	e := doRequest(apiArtistsURL, params, "GET", &res)
	return res.Artists, e
}

// Get Spotify catalog information about artists similar to a given artist.
// Similarity is based on analysis of the Spotify community’s listening history
func GetRelatedArtists(id string) ([]Artist, error) {
	res := struct {
		Artists []Artist `json:"artists"`
	}{
		Artists: make([]Artist, 0),
	}
	if e := doRequest(fmt.Sprintf(apiRelatedArtistsURL, id), nil, "GET", &res); e != nil {
		return nil, e
	}
	return res.Artists, nil
}

// GetArtistAlbums fetches albums of specified types for an artist.
// limit and offset are optional parameters, set the to -1 for default values.
// Returns a list of albums and the total amount of albums found in the query.
func GetArtistAlbums(id string, types []string, limit int, offset int) ([]Album, int, error) {
	path := fmt.Sprintf(apiArtistAlbumURL, id)
	// params
	params := url.Values{}
	params.Set("album_type", strings.Join(types, ","))
	if limit != -1 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if offset != -1 {
		params.Set("offset", strconv.Itoa(offset))
	}

	// make request
	res := struct {
		paging
		Items []Album `json:"items"`
	}{}
	if e := doRequest(path, params, "GET", &res); e != nil {
		return nil, 0, e
	}
	return res.Items, res.Total, nil
}

// Get artist top tracks by country.
func GetArtistTopTracks(id, country string) ([]Track, error) {
	path := fmt.Sprintf(apiArtistTopTracksURL, id)
	res := struct {
		Tracks []Track `json:"tracks"`
	}{}
	params := url.Values{}
	params.Set("country", country)
	if e := doRequest(path, params, "GET", &res); e != nil {
		return nil, e
	}
	return res.Tracks, nil
}

func GetAlbum(id string) (Album, error) {
	var album Album
	e := doRequest(fmt.Sprintf(apiAlbumURL, id), nil, "GET", &album)
	return album, e
}

// Get albums by ids.
func GetAlbums(id ...string) ([]Album, error) {
	res := struct {
		Albums []Album `json:"albums"`
	}{}
	params := url.Values{}
	params.Set("ids", strings.Join(id, ","))
	e := doRequest(apiAlbumsURL, params, "GET", &res)
	return res.Albums, e
}

// Get album tracks.
func GetAlbumTracks(id string) ([]Track, error) {
	// response structure
	res := struct {
		paging
		Items []Track `json:"items"`
	}{}
	path := fmt.Sprintf(apiAlbumTracksURL, id)
	// params
	params := url.Values{}
	// make request
	if e := doRequest(path, params, "GET", &res); e != nil {
		return nil, e
	}
	return res.Items, nil
}
