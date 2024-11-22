package fetch

import (
	"fmt"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/fatal"
	"github.com/kingtingthegreat/top-fetch-cli/initialize"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func LocalFetch(cfg config.Config) (string, string) {
	if cfg.SpotifyClientId == "" || cfg.SpotifyClientSecret == "" {
		fatal.Fatal(cfg.Silent, "Spotify client id or client secret is not set")
	}
	if cfg.SpotifyAccessToken == "" || cfg.SpotifyRefreshToken == "" {
		fmt.Println("please visit", spotify.AuthUrl(cfg.SpotifyClientId, "http://localhost:8080"))
		var err error
		cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, err = initialize.InitSpotify(cfg)
		if err != nil {
			fatal.Fatal(cfg.Silent, err.Error())
		}
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	}
	// log.Println("getting top track")
	track, newAccessToken, err := spotify.GetUserTopTrack(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	if err != nil {
		fatal.Fatal(cfg.Silent, err.Error())
	}
	if newAccessToken != "" {
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, newAccessToken, cfg.SpotifyRefreshToken)
	}

	return track.Album.Images[0].Url, track.Name + " - " + track.Artists[0].Name
}
