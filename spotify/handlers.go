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
	playlist := strings.Trim(data.Name, " ")
	playlist = strings.ReplaceAll(playlist, " ", "_")
	playlist = strings.ToLower(playlist)
	playlist = strings.ReplaceAll(playlist, ".", "")
	playlist = strings.ReplaceAll(playlist, ",", "_")
	playlist = strings.ReplaceAll(playlist, "/", "_")
	playlist = strings.ReplaceAll(playlist, "__", "_")
	return fmt.Sprintf("%s.m3u", playlist)
}
