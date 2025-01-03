package library

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"spt-to-m3u/spotify"
	"strings"
)

func FindArtistFolder(localPath, author string) (string, error) {
	fileSystem := os.DirFS(localPath)
	var firstMatch string

	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if firstMatch != "" {
			return nil
		}
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		match := strings.Contains(safeMatcherPath(path), safeMatcherPath(author))
		if match {
			firstMatch = fmt.Sprintf("%s/%s", localPath, path)
		}
		return nil
	})

	if err != nil {
		return firstMatch, err
	}
	if firstMatch == "" {
		return firstMatch, errors.New("Failed to find a matching directory")
	}
	return firstMatch, nil
}

func FindSong(localPath string, pair spotify.SimpleSongAuthorPair) (string, error) {
	artistFolder, err := FindArtistFolder(localPath, pair.Author)
	if err == nil {
		localPath = artistFolder
	}
	// fmt.Printf("\n\n path: %s, song: %s", localPath, safeMatcherPath(pair.Song))
	fileSystem := os.DirFS(localPath)
	var firstMatch string

	err = fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if firstMatch != "" {
			return nil
		}
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		file := safeMatcherPath(d.Name())
		song := safeMatcherPath(pair.Song)
		// fmt.Printf("\n\n file: %s   parsed: %s   song: %s", d.Name(), file, song)
		match := strings.Contains(file, song)
		if match {
			firstMatch = fmt.Sprintf("%s/%s", localPath, path)
		}
		return nil
	})

	if err != nil {
		return firstMatch, err
	}
	return firstMatch, nil
}

func safeMatcherPath(p string) string {
	reg := regexp.MustCompile(`[0-9]`)
	p = reg.ReplaceAllString(p, "")

	p = strings.ToLower(p)
	p = strings.ReplaceAll(p, "remaster", "")
	p = strings.ReplaceAll(p, "live", "")
	// p = strings.Split(p, "-")[0]

	p = strings.ReplaceAll(p, " ", "")
	p = strings.ReplaceAll(p, "_", "")
	p = strings.ReplaceAll(p, "-", "")
	p = strings.ReplaceAll(p, ".", "")
	p = strings.ReplaceAll(p, ",", "")

	p = strings.ReplaceAll(p, "ä", "a")
	p = strings.ReplaceAll(p, "ë", "e")
	p = strings.ReplaceAll(p, "ï", "i")
	p = strings.ReplaceAll(p, "ö", "o")
	p = strings.ReplaceAll(p, "ü", "u")

	p = strings.ReplaceAll(p, "á", "a")
	p = strings.ReplaceAll(p, "é", "e")
	p = strings.ReplaceAll(p, "í", "i")
	p = strings.ReplaceAll(p, "ó", "o")
	p = strings.ReplaceAll(p, "ú", "u")

	return p
}
