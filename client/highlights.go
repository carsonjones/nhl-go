package nhl

import (
	"fmt"
)

// VideoItem represents a video/highlight from NHL
type VideoItem struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	BrightcoveID string `json:"brightcoveId"`
	URL         string `json:"url"`
	Thumbnail   string `json:"thumbnail"`
}

// HighlightsResponse contains video highlights
type HighlightsResponse struct {
	Items []VideoItem `json:"items"`
}

// GetGameHighlights returns video highlights for a specific game
func (c *Client) GetGameHighlights(gameID int) (*HighlightsResponse, error) {
	url := fmt.Sprintf("https://forge-dapi.d3.nhle.com/v2/content/en-us/videos?tags.slug=gameid-%d", gameID)
	
	var rawResponse struct {
		Items []struct {
			Title     string `json:"title"`
			Slug      string `json:"slug"`
			Thumbnail struct {
				TemplateURL string `json:"templateUrl"`
			} `json:"thumbnail"`
			Fields struct {
				Description string `json:"description"`
				Duration    string `json:"duration"`
				BrightcoveID string `json:"brightcoveId"`
			} `json:"fields"`
		} `json:"items"`
	}
	
	err := c.get(url, &rawResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to get highlights: %v", err)
	}
	
	response := &HighlightsResponse{
		Items: make([]VideoItem, 0, len(rawResponse.Items)),
	}
	
	for _, item := range rawResponse.Items {
		video := VideoItem{
			Title:       item.Title,
			Slug:        item.Slug,
			Description: item.Fields.Description,
			Duration:    item.Fields.Duration,
			BrightcoveID: item.Fields.BrightcoveID,
			URL:         fmt.Sprintf("https://www.nhl.com/video/%s", item.Slug),
			Thumbnail:   item.Thumbnail.TemplateURL,
		}
		response.Items = append(response.Items, video)
	}
	
	return response, nil
}
