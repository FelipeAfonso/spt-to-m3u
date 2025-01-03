package spotify

import "fmt"

func GetSongsList(data PlaylistResponse) []SimpleSongAuthorPair {
	arr := []SimpleSongAuthorPair{}
	for _, track := range data.Tracks.Items {
		name := track.Track.Name
		author := track.Track.Artists[0].Name
		fmt.Printf("\n%s - %s", name, author)
		arr = append(arr, SimpleSongAuthorPair{
			Song:   name,
			Author: author,
		})
	}
	return arr
}
