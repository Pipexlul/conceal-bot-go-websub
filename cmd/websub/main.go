package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Subscription struct {
	Topic    string
	Callback string
}

type SeenID struct {
	ID     string
	SeenTS time.Time
}

var subscriptions = []Subscription{}
var mu sync.Mutex

type USGSFeedEntry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Summary string `xml:"summary"`
	Link    struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`
}

type USGSFeed struct {
	Entries []USGSFeedEntry `xml:"entry"`
}

const (
	usgsFeedURL = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.atom"
)

func handleSubscription(w http.ResponseWriter, r *http.Request) {
	mode := r.FormValue("hub.mode")
	topic := r.FormValue("hub.topic")
	callback := r.FormValue("hub.callback")

	mu.Lock()
	defer mu.Unlock()

	if mode == "subscribe" {
		subscriptions = append(subscriptions, Subscription{Topic: topic, Callback: callback})
		fmt.Printf("New subscription for topic: '%s' with callback '%s'\n", topic, callback)
		w.WriteHeader(http.StatusAccepted)
	} else if mode == "unsubscribe" {
		for i, sub := range subscriptions {
			if sub.Callback == callback && sub.Topic == topic {
				subscriptions = append(subscriptions[:i], subscriptions[i+1:]...)
				break
			}
		}
		w.WriteHeader(http.StatusAccepted)
	} else {
		http.Error(w, "Invalid hub.mode", http.StatusBadRequest)
	}
}

func fetchAndPublishUSGSFeed() {
	seenIDs := make(map[string]SeenID)

	for {
		resp, err := http.Get(usgsFeedURL)
		if err != nil {
			fmt.Println("Error fetching USGS feed:", err)
			time.Sleep(time.Minute)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Error reading USGS feed:", err)
			continue
		}

		var feed USGSFeed
		err = xml.Unmarshal(body, &feed)
		if err != nil {
			fmt.Println("Error parsing USGS feed:", err)
			continue
		}

		now := time.Now().UTC()
		for _, entry := range feed.Entries {
			if _, found := seenIDs[entry.ID]; !found {
				fmt.Printf("New earthquake detected: %s\n", entry.Title)
				publishUpdate(&entry)
				seenIDs[entry.ID] = SeenID{ID: entry.ID, SeenTS: now}
			}
		}

		time.Sleep(time.Minute)
	}
}

func publishUpdate(entry *USGSFeedEntry) {
	mu.Lock()
	defer mu.Unlock()

	for _, sub := range subscriptions {
		if sub.Topic == "earthquakes" {
			content := fmt.Sprintf("New earthquake:\nTitle: %s\nSummary: %s\nLink: %s", entry.Title, entry.Summary, entry.Link.Href)
			fmt.Printf("Notifying subscriber: %s\n", sub.Callback)
			http.Post(sub.Callback, "text/plain", strings.NewReader(content))
		}
	}
}

func main() {
	http.HandleFunc("/hub", handleSubscription)

	go fetchAndPublishUSGSFeed()

	port := os.Getenv("PORT")
	if port == "" {
		port = "50000"
	}
	fmt.Println("WebSub listening on port", port)

	port = ":" + port

	http.ListenAndServe(port, nil)
}
