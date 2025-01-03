package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"spt-to-m3u/spotify"
)

func main() {
	// 	cli.RootCommandHelpTemplate = fmt.Sprintf(`%s
	// WEBSITE: http://awesometown.example.com
	// `, cli.RootCommandHelpTemplate)

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
		},
		Usage:     "Import your spotify playlists into your local library.",
		ArgsUsage: "SPOTIFY_PLAYLIST_ID",
		UsageText: "spt-to-m3u -c YOUR_CLIENT_ID -s YOUR_CLIENT_SECRET SPOTIFY_PLAYLIST_ID",
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
			fmt.Printf("\nAuthenticating to look for Playlist %s...", id)
			token, err := spotify.AuthenticateWithClientSecret(client_id, client_secret)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\nAuthenticated with user credentials %s...", token)

			body, err := spotify.GetPlaylistData(id, token)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n\nPlaylist found %s...", body)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
