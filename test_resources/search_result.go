package main

import (
	"encoding/json"
	"time"
)

type MigalooSearchResult struct {
	Category      string     `json:"category"`
	Content       string     `json:"content"`
	Engine        string     `json:"engine"`
	Engines       []string   `json:"engines"`
	Metadata      string     `json:"metadata"`
	ParsedURL     [][]string `json:"parsed_url"`
	Positions     []int      `json:"positions"`
	PublishedDate time.Time  `json:"published_date"`
	Score         float64    `json:"score"`
	Template      string     `json:"template"`
	Thumbnail     string     `json:"thumbnail"`
	Title         string     `json:"title"`
	URL           string     `json:"url"`
}

type MigalooSearchResponse struct {
	Answers         []interface{}         `json:"answers"`
	Corrections     []interface{}         `json:"corrections"`
	Infoboxes       []interface{}         `json:"infoboxes"`
	NumberOfResults int                   `json:"number_of_results"`
	Query           string                `json:"query"`
	Results         []MigalooSearchResult `json:"results"`
}

func (r *MigalooSearchResponse) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		MigalooSearchResponse
	}{
		*r,
	})
}

func NewSearchResponseFromJSON(data []byte) (*MigalooSearchResponse, error) {
	var response MigalooSearchResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
