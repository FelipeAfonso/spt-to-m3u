package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"spt-to-m3u/library"
	"spt-to-m3u/spotify"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "client_id",
				Aliases:  []string{"c"},
				Required: true,
				Usage:    "Client ID taken from your developer dashboard",
			},
			&cli.StringFlag{
				Name:     "client_secret",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "Client Secret taken from your developer dashboard",
			},
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   "./",
				Usage:   "Path to local library if not CWD",
			},
		},
		Usage:     "Import your spotify playlists into your local library.",
		ArgsUsage: "SPOTIFY_PLAYLIST_ID",
		UsageText: "spt-to-m3u -c YOUR_CLIENT_ID -s YOUR_CLIENT_SECRET [-p PATH_TO_LIB] SPOTIFY_PLAYLIST_ID",
		Description: `You can get your SPOTIFY_PLAYLIST_ID by inspecting 
the playlist URL (Either by opening in your browser, 
or checking the share link's URL). It should be the 
last bit of text after the last '/'. i.e.:

  The playlist's URL
https://open.spotify.com/playlist/39y5RxyW8k8r24onnewuNMn

  The SPOTIFY_PLAYLIST_ID
39y5RxyW8k8r24onnewuNMn`,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			id := cmd.Args().First()
			client_id := cmd.String("client_id")
			client_secret := cmd.String("client_secret")
			path := cmd.String("path")

			fmt.Printf("\nAuthenticating to look for Playlist %s...", id)
			token, err := spotify.AuthenticateWithClientSecret(client_id, client_secret)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\nAuthenticated with user credentials Successfully...")

			body, err := spotify.GetPlaylistData(id, token)
			if err != nil {
				log.Fatal(err)
			}
			songs := spotify.GetSongsList(*body)
			fmt.Printf("\nPlaylist found: %s\nStarted Matching Songs:\n", body.Name)

			for i, pair := range songs {
				matched, err := library.FindSong(path, pair)
				if err != nil {
					log.Fatal(err)
				}
				status := ""
				if matched != "" {
					status = "~ MATCHED"
				}
				fmt.Printf("\n[%d] %s - %s %s", i+1, pair.Author, pair.Song, status)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
