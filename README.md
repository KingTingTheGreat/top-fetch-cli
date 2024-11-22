# Top Fetch CLI

## Set Up

### Web
1. Go to Spotify for developers and create a new application with http://localhost:8080 as a callback uri
2. Create an *.env* file and add entries for SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET
3. Run ```go run main.go``` and follow the one-time instructions

### Local
1. Go to [Top Fetch](https://top-fetch.vercel.app) and sign in with Spotify
2. Copy your ID
3. Run ```go run main.go *your-id*``` or create a *.env* file and add your id as TOP_FETCH_ID
