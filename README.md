# Spotify to M3U Playlist Converter

A command-line tool that helps you convert your Spotify playlists into M3U format, matching songs with your local music library. 

Disclaimer: this README is mostly generated with AI using the codebase, cause I like coding not writing docs.

## Features

- Convert Spotify playlists to M3U format
- Intelligent local file matching
- Support for custom library paths
- Attempt to maintain original playlist structure

## Prerequisites

Before using this tool, you'll need:

- Go installed on your system
- A Spotify Developer account
- Your local music library organized in a directory

## Getting Spotify Credentials

To use this tool, you'll need to obtain credentials from Spotify's Developer Dashboard. Here's how:

1. Visit [Spotify Developer Dashboard](https://developer.spotify.com/dashboard)
2. Log in with your Spotify account
3. Click "Create an App"
4. Fill in the app name and description
5. Once created, you'll see your Client ID on the dashboard
6. Click "Show Client Secret" to reveal your Client Secret
7. Save both the Client ID and Client Secret - you'll need these to use the tool

## Usage

The basic command structure is:

```bash
spt-to-m3u -c YOUR_CLIENT_ID -s YOUR_CLIENT_SECRET [-p PATH_TO_LIB] SPOTIFY_PLAYLIST_ID
```

### Arguments

- `SPOTIFY_PLAYLIST_ID`: The ID of your Spotify playlist (required)

### Flags

- `-c, --client_id`: Your Spotify Client ID (required)
- `-s, --client_secret`: Your Spotify Client Secret (required)
- `-p, --path`: Path to your local music library (optional, defaults to current directory)

### Finding Your Playlist ID

To get your Spotify playlist ID:

1. Open the playlist in Spotify
2. Click "Share" and select "Copy link"
3. The playlist ID is the string after the last forward slash
   - Example: In `https://open.spotify.com/playlist/39y5RxyW8k8r24onnewuNMn`
   - The playlist ID is `39y5RxyW8k8r24onnewuNMn`

### Example

```bash
spt-to-m3u -c abc123def456 -s ghijk789lmno -p /path/to/music/library 39y5RxyW8k8r24onnewuNMn
```

## Output

The tool will:
1. Search on Spotify for your playlist
2. Create an M3U file named after your Spotify playlist
3. Search your local library for matching songs
4. Generate a compatible M3U playlist file

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- Built using the [urfave/cli](https://github.com/urfave/cli) package
- Thanks to Spotify for providing their Web API
