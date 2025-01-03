package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func AuthenticateWithClientSecret(id, secret string) (string, error) {
	formBody := url.Values{}
	formBody.Add("grant_type", "client_credentials")
	formBody.Add("client_id", id)
	formBody.Add("client_secret", secret)
	resp, err := http.PostForm("https://accounts.spotify.com/api/token", formBody)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenStruct AccessTokenResponse
	if err := json.Unmarshal(body, &tokenStruct); err != nil {
		return "", err
	}
	return tokenStruct.AccessToken, nil
}

func GetPlaylistData(id, token string) (*PlaylistResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", id), nil)
	if err != nil {
		return nil, err

	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var playlist PlaylistResponse
	if err := json.Unmarshal(body, &playlist); err != nil {
		return nil, err
	}
	return &playlist, nil
}
