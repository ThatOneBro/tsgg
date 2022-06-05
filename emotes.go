package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const emoteEndpoint = "https://cdn.destiny.gg/2.43.4/emotes/emotes.json"

type ImageData struct {
	Height int32  `json:"height"`
	Width  int32  `json:"width"`
	Mime   string `json:"mime"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}
type Emote struct {
	Prefix string      `json:"prefix"`
	Twitch bool        `json:"twitch"`
	Theme  string      `json:"theme"`
	Image  []ImageData `json:"image"`
}

func getEmoteNames(e []Emote) []string {
	s := make([]string, len(e))
	for i, em := range e {
		s[i] = em.Prefix
	}
	return s
}

func getEmotes() ([]string, error) {
	emotes := make([]string, 0)

	resp, err := http.Get(emoteEndpoint)
	if err != nil {
		return emotes, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emotes, fmt.Errorf("emote endpoint status code %d", resp.StatusCode)
	}

	// var em []Emote = make([]Emote, 0)
	var em []Emote

	err = json.NewDecoder(resp.Body).Decode(&em)
	if err != nil {
		return emotes, err
	}

	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return emotes, err
	// }

	// err = json.Unmarshal(b, &em)
	// if err != nil {
	// 	return emotes, err
	// }
	// for _, item := range em {
	// 	log.Println(item)
	// }

	emotes = append(emotes, getEmoteNames(em)...)
	return emotes, nil
}
