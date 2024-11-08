package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/Shurik12/go-yamusic-api/yamusic"
)

func main() {
	// read config
	type Configuration struct {
		Token  string
		Output string
		Log    string
		Host   string
		Port   string
	}

	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	token := configuration.Token

	// it implements yamusic.Doer interface
	client := yamusic.NewClient(
		// create default user with your auth token
		// provide user_id and access_token (needed by some methods)
		yamusic.AccessToken(0, token),
	)
	// get user by auth token
	accountStatus, _, _ := client.Account().GetUser(context.Background())
	// add userID to client
	client.SetUserID(accountStatus.Result.UID)

	// client.PrintPlaylists()
	// client.Tracks().GetLike(context.Background())

	track := client.GetTracksWithoutPlaylist()

	track_id, _ := strconv.Atoi(track[0].ID)
	uri, _ := client.Tracks().GetDownloadURL(context.Background(), track_id)
	client.Tracks().Download(context.Background(), uri)
	fmt.Println(uri)
}
