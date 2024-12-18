package fetch

import (
	"fmt"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
	"github.com/kingtingthegreat/top-fetch-cli/initialize"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func LocalFetch() (string, string) {
	cfg := config.Config()
	if cfg.SpotifyClientId == "" || cfg.SpotifyClientSecret == "" {
		fatal.Fatal("Spotify client id or client secret is not set")
	}
	if cfg.SpotifyAccessToken == "" || cfg.SpotifyRefreshToken == "" {
		fmt.Println("please visit", spotify.AuthUrl(cfg.SpotifyClientId, "http://localhost:8080"))
		var err error
		cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, err = initialize.InitSpotify(cfg.SpotifyClientId, cfg.SpotifyClientSecret)
		if err != nil {
			fatal.Fatal(err.Error())
		}
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	}
	// log.Println("getting top track")
	track, newAccessToken, err := spotify.GetUserTopTrack(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	if err != nil {
		fatal.Fatal(err.Error())
	}
	if newAccessToken != "" {
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, newAccessToken, cfg.SpotifyRefreshToken)
	}

	return track.Album.Images[0].Url, track.Name + " - " + track.Artists[0].Name
}
