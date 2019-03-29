package choose_your_own_adventure

import (
	"encoding/json"
	"io"
)

func JsonStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
