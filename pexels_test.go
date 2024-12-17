package main

import (
	"log"
	"testing"

	"github.com/kumarsgoyal/pexels-go/client"
	"github.com/kumarsgoyal/pexels-go/config"
	"github.com/kumarsgoyal/pexels-go/types"
)

var testClient *client.PexelsClient

// Set up the test client before each test case
func setup() {
	// Load the configuration
	cfg, err := config.LoadConfig(".apiConfig")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// Initialize the client
	testClient = client.NewClient(cfg.PexelAPIKey)
}

// Test photo search functionality
func TestPhotoSearch(t *testing.T) {
	setup()
	searchParams := &types.PhotoSearchParams{
		Query:       "elephant",
		Orientation: "landscape",
		Size:        "large",
		Page:        1,
		PerPage:     5,
	}
	photosResponse, err := testClient.Photos.Search(searchParams)
	if err != nil {
		t.Fatalf("Error searching photos: %v", err)
	}

	if len(photosResponse.Photos) == 0 {
		t.Fatal("No photos found")
	}

	for _, photo := range photosResponse.Photos {
		log.Printf("Photo ID: %d, Photographer: %s, URL: %s", photo.ID, photo.Photographer, photo.URL)
	}
}

// Test for curated photos
func TestPhotoCurated(t *testing.T) {
	setup()

	searchParams := &types.PaginationParams{
		Page:    1,
		PerPage: 10,
	}
	photosResponse, err := testClient.Photos.Curated(searchParams)
	if err != nil {
		t.Fatalf("Error fetching curated photos: %v", err)
	}

	if len(photosResponse.Photos) == 0 {
		t.Fatal("No curated photos found")
	}

	for _, photo := range photosResponse.Photos {
		t.Logf("Photo ID: %d, Photographer: %s, URL: %s", photo.ID, photo.Photographer, photo.URL)
	}
}

// Test fetching a single photo by ID
func TestGetPhotoByID(t *testing.T) {
	setup()
	photoID := 2014422
	photosResponse, err := testClient.Photos.GetPhoto(photoID)
	if err != nil {
		t.Fatalf("Error fetching photo: %v", err)
	}

	log.Printf("Fetched Photo: %+v", photosResponse)
}

// Test video search functionality
func TestVideoSearch(t *testing.T) {
	setup()
	searchParams := &types.VideoSearchParams{
		Query:       "elephant",
		Orientation: "landscape",
		Size:        "small",
		Page:        1,
		PerPage:     5,
	}

	videosResponse, err := testClient.Videos.Search(searchParams)
	if err != nil {
		t.Fatalf("Error searching videos: %v", err)
	}

	if len(videosResponse.Videos) == 0 {
		t.Fatal("No videos found")
	}

	for _, video := range videosResponse.Videos {
		t.Logf("Video ID: %d", video.ID)
		t.Logf("Photographer: %s, URL: %s", video.User.Name, video.User.URL)
		t.Logf("Video URL: %s", video.URL)
		t.Logf("Duration: %d seconds", video.Duration)
		t.Logf("Image: %s", video.Image)

		for _, videoFile := range video.VideoFiles {
			t.Logf("Video File (Quality: %s, Resolution: %dx%d): %s", videoFile.Quality, videoFile.Width, videoFile.Height, videoFile.Link)
		}

		for _, picture := range video.VideoPictures {
			t.Logf("Picture ID: %d, Picture URL: %s", picture.ID, picture.Picture)
		}
	}
}

func TestVideosPopular(t *testing.T) {
	setup()

	filterParams := &types.VideoFilterParams{
		MinWidth:    640,
		MinHeight:   480,
		MinDuration: 30,
		MaxDuration: 300,
		Page:        1,
		PerPage:     5,
	}

	videosResponse, err := testClient.Videos.Popular(filterParams)
	if err != nil {
		t.Fatalf("Error fetching popular videos: %v", err)
	}

	if len(videosResponse.Videos) == 0 {
		t.Fatal("No popular videos found")
	}

	for _, video := range videosResponse.Videos {
		t.Logf("Video ID: %d", video.ID)
		t.Logf("Photographer: %s, URL: %s", video.User.Name, video.User.URL)
		t.Logf("Video URL: %s", video.URL)
		t.Logf("Duration: %d seconds", video.Duration)
		t.Logf("Image: %s", video.Image)

		for _, videoFile := range video.VideoFiles {
			t.Logf("Video File (Quality: %s, Resolution: %dx%d): %s", videoFile.Quality, videoFile.Width, videoFile.Height, videoFile.Link)
		}

		for _, picture := range video.VideoPictures {
			t.Logf("Picture ID: %d, Picture URL: %s", picture.ID, picture.Picture)
		}
	}
}

// Test fetching a video by ID
func TestGetVideoByID(t *testing.T) {
	setup()
	videoID := 2499611 // Replace with an actual video ID
	video, err := testClient.Videos.GetVideo(videoID)
	if err != nil {
		t.Fatalf("Error fetching video details: %v", err)
	}

	t.Logf("Video ID: %d", video.ID)
	t.Logf("Photographer: %s, URL: %s", video.User.Name, video.User.URL)
	t.Logf("Video URL: %s", video.URL)
	t.Logf("Duration: %d seconds", video.Duration)
	t.Logf("Image: %s", video.Image)

	for _, videoFile := range video.VideoFiles {
		t.Logf("Video File (Quality: %s, Resolution: %dx%d): %s", videoFile.Quality, videoFile.Width, videoFile.Height, videoFile.Link)
	}

	for _, picture := range video.VideoPictures {
		t.Logf("Picture ID: %d, Picture URL: %s", picture.ID, picture.Picture)
	}
}

// Test fetching all collections with pagination
func TestGetAllCollections(t *testing.T) {
	setup()

	// Define pagination parameters
	paginationParams := types.PaginationParams{
		Page:    1,
		PerPage: 1,
	}

	// Fetch all collections
	collectionsResponse, err := testClient.Collections.All(paginationParams)
	if err != nil {
		t.Fatalf("Error fetching collections: %v", err)
	}

	// Log the collections
	for _, collection := range collectionsResponse.Collections {
		t.Logf("Collection ID: %s, Title: %s, Photos Count: %d, Videos Count: %d",
			collection.ID, collection.Title, collection.PhotosCount, collection.VideosCount)
	}

	// Check pagination links for next/previous pages
	t.Logf("Next Page: %s", collectionsResponse.NextPage)
	t.Logf("Previous Page: %s", collectionsResponse.PrevPage)
}

// Test fetching featured collections
func TestGetFeaturedCollections(t *testing.T) {
	setup()

	// Define pagination parameters
	paginationParams := types.PaginationParams{
		Page:    1,
		PerPage: 1,
	}

	// Fetch featured collections
	featuredCollectionsResponse, err := testClient.Collections.Featured(paginationParams)
	if err != nil {
		t.Fatalf("Error fetching featured collections: %v", err)
	}

	if len(featuredCollectionsResponse.Collections) == 0 {
		t.Fatal("No featured collections found")
	}

	// Log the featured collections
	for _, collection := range featuredCollectionsResponse.Collections {
		t.Logf("Featured Collection ID: %s, Title: %s, Photos Count: %d, Videos Count: %d",
			collection.ID, collection.Title, collection.PhotosCount, collection.VideosCount)
	}
}

// Test fetching media from a specific collection (photos or videos)
func TestGetMediaFromCollection(t *testing.T) {
	setup()

	// Define media parameters
	mediaParams := types.MediaParams{
		CollectionID: "5qa21sj", // Replace with a valid collection ID
		MediaType:    "photos",  // Specify the media type (photos or videos)
		Sort:         "desc",    // Specify sort order (asc or desc)
		Pagination: types.PaginationParams{
			Page:    1,
			PerPage: 5,
		},
	}

	// Fetch media (photos or videos) from the collection
	mediaResponse, err := testClient.Collections.Media(mediaParams)
	if err != nil {
		t.Fatalf("Error fetching media for collection: %v", err)
	}

	// Log the media details
	for _, media := range mediaResponse.Media {
		t.Logf("Media ID: %d, URL: %s", media.ID, media.URL)
	}
}
