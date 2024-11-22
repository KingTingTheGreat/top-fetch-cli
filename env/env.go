package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var envFile = ""

func GetBasePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("cannot get executable path")
	}

	execDir := filepath.Dir(execPath)
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current working directory")
	}

	if !strings.HasPrefix(execDir, os.TempDir()) {
		return execDir, nil
	}
	return cwd, nil
}

func LoadEnv() error {
	basePath, err := GetBasePath()
	if err != nil {
		return err
	}

	envFile = filepath.Join(basePath, ".env")

	godotenv.Load(envFile)

	return nil
}

func SaveEnv(clientId, clientSecret, accessToken, refreshToken string) {
	envVars := map[string]string{
		"SPOTIFY_CLIENT_ID":     clientId,
		"SPOTIFY_CLIENT_SECRET": clientSecret,
		"SPOTIFY_ACCESS_TOKEN":  accessToken,
		"SPOTIFY_REFRESH_TOKEN": refreshToken,
	}

	godotenv.Write(envVars, envFile)
}
