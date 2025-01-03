package spotify

import (
	"fmt"
	"strings"
)

func GetSongsList(data PlaylistResponse) []SimpleSongAuthorPair {
	arr := []SimpleSongAuthorPair{}
	for _, track := range data.Tracks.Items {
		name := track.Track.Name
		author := track.Track.Artists[0].Name
		arr = append(arr, SimpleSongAuthorPair{
			Song:   name,
			Author: author,
		})
	}
	return arr
}

func SafePlaylistName(data PlaylistResponse) string {
	trimmed := strings.Trim(data.Name, " ")
	spaceless := strings.ReplaceAll(trimmed, " ", "_")
	lowercase := strings.ToLower(spaceless)
	safe := strings.ReplaceAll(lowercase, ".", "")
	return fmt.Sprintf("%s.m3u", safe)
}
