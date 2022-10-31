package main

import (
	"fmt"
	"os"

	apivideosdk "github.com/apivideo/api.video-go-client"
)

func main() {
	//Connect to production environment
	client := apivideosdk.SandboxClientBuilder("9zqu212GCXkAs8BAMWhkCALQOdnvEuB6tHQAorbU2me").Build()

	//List Videos
	//First create the url options for searching
	opts := apivideosdk.VideosApiListRequest{}.
		CurrentPage(1).
		PageSize(25).
		SortBy("publishedAt").
		SortOrder("desc")

	//Then call the List endpoint with the options
	result, err := client.Videos.List(opts)

	if err != nil {
		fmt.Println(err)
	}

	for _, video := range result.Data {
		fmt.Printf("%s\n", video.VideoId)
		fmt.Printf("%s\n", *video.Title)
	}

	//Upload a video
	//First create a container
	create, err := client.Videos.Create(apivideosdk.VideoCreationPayload{Title: "My video title"})

	if err != nil {
		fmt.Println(err)
	}

	//Then open the video file
	videoFile, err := os.Open("/home/tomat/Videos/waves_4k.mp4")

	if err != nil {
		fmt.Println(err)
	}

	//Finally upload your video to the container with the videoId
	uploadedVideo, err := client.Videos.UploadFile(create.VideoId, videoFile)
	if err != nil {
		fmt.Println(err)
	}

	//And get the assets
	fmt.Printf("%s\n", *uploadedVideo.Assets.Hls)
	fmt.Printf("%s\n", *uploadedVideo.Assets.Iframe)
}
