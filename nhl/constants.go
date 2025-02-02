package nhl

// GameType represents different types of NHL games
type GameType int

const (
	GameTypeRegularSeason GameType = 2
	GameTypePlayoffs      GameType = 3
	GameTypeAllStar       GameType = 4
)

// SortOrder represents different ways to sort game schedules
type SortOrder string

const (
	SortByDateAsc  SortOrder = "asc"
	SortByDateDesc SortOrder = "desc"
)

// Base URLs for NHL API
const (
	BaseURLWeb = "https://api-web.nhle.com/v1"
)
