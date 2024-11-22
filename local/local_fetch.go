package local

import (
	"fmt"
	"log"

	"github.com/kingtingthegreat/top-fetch-cli/config"
	"github.com/kingtingthegreat/top-fetch-cli/env"
	"github.com/kingtingthegreat/top-fetch-cli/initialize"
	"github.com/kingtingthegreat/top-fetch/spotify"
)

func LocalFetch(cfg config.Config) (string, string) {
	if cfg.SpotifyClientId == "" || cfg.SpotifyClientSecret == "" {
		log.Fatal("Spotify client id or client secret is not set")
	}
	if cfg.SpotifyAccessToken == "" || cfg.SpotifyRefreshToken == "" {
		fmt.Println("please visit", spotify.AuthUrl(cfg.SpotifyClientId, "http://localhost:8080"))
		var err error
		cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken, err = initialize.InitSpotify(cfg)
		if err != nil {
			log.Fatal(err.Error())
		}
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	}
	// log.Println("getting top track")
	track, newAccessToken, err := spotify.GetUserTopTrack(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyAccessToken, cfg.SpotifyRefreshToken)
	if err != nil {
		log.Fatal(err.Error())
	}
	if newAccessToken != "" {
		go env.SaveEnv(cfg.SpotifyClientId, cfg.SpotifyClientSecret, newAccessToken, cfg.SpotifyRefreshToken)
	}

	return track.Name + " - " + track.Artists[0].Name, track.Album.Images[0].Url

}
