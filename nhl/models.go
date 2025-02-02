package nhl

// Schedule represents the NHL schedule response
type Schedule struct {
	NextStartDate     string        `json:"nextStartDate"`
	PreviousStartDate string        `json:"previousStartDate"`
	GameWeek          []GameDay     `json:"gameWeek"`
	Games             []Game        `json:"games"`
	Events            []interface{} `json:"events"`
}

// GameDay represents a day in the game week
type GameDay struct {
	Date  string `json:"date"`
	Games []Game `json:"games"`
}

// Game represents an NHL game
type Game struct {
	ID                int              `json:"id"`
	Season            int              `json:"season"`
	GameType          int              `json:"gameType"`
	GameDate          string           `json:"gameDate"`
	GameCenterLink    string           `json:"gameCenterLink"`
	Venue             Venue            `json:"venue"`
	StartTimeUTC      string           `json:"startTimeUTC"`
	EasternUTCOffset  string           `json:"easternUTCOffset"`
	VenueUTCOffset    string           `json:"venueUTCOffset"`
	TVBroadcasts      []TVBroadcast    `json:"tvBroadcasts"`
	GameState         string           `json:"gameState"`
	GameScheduleState string           `json:"gameScheduleState"`
	AwayTeam          Team             `json:"awayTeam"`
	HomeTeam          Team             `json:"homeTeam"`
	Period            int              `json:"period"`
	PeriodDescriptor  PeriodDescriptor `json:"periodDescriptor"`
}

// Venue represents a game venue
type Venue struct {
	Default string `json:"default"`
}

// Team represents a team in a game
type Team struct {
	ID                       int           `json:"id"`
	Name                     LanguageNames `json:"name"`
	CommonName               LanguageNames `json:"commonName"`
	PlaceNameWithPreposition LanguageNames `json:"placeNameWithPreposition"`
	Abbrev                   string        `json:"abbrev"`
	Score                    int           `json:"score"`
	Logo                     string        `json:"logo"`
}

// LanguageNames represents a team's name information
type LanguageNames struct {
	Default string `json:"default"`
	Fr      string `json:"fr,omitempty"`
}

// GameOutcome represents the outcome of a game
type GameOutcome struct {
	LastPeriodType string `json:"lastPeriodType"`
}

// PlayerInfo represents detailed information about a player
type PlayerInfo struct {
	ID           int      `json:"id"`
	FirstName    NameInfo `json:"firstName"`
	LastName     NameInfo `json:"lastName"`
	Position     string   `json:"positionCode"`
	JerseyNumber int      `json:"sweaterNumber"`
	Height       int      `json:"heightInInches"`
	Weight       int      `json:"weightInPounds"`
	BirthDate    string   `json:"birthDate"`
	BirthCity    NameInfo `json:"birthCity"`
	BirthCountry string   `json:"birthCountry"`
	BirthState   NameInfo `json:"birthStateProvince,omitempty"`
	Shoots       string   `json:"shootsCatches"`
	Headshot     string   `json:"headshot"`
}

// NameInfo represents a name with different language versions
type NameInfo struct {
	Default string `json:"default"`
}

// TeamRoster represents a team's roster
type TeamRoster struct {
	Forwards   []PlayerInfo `json:"forwards"`
	Defensemen []PlayerInfo `json:"defensemen"`
	Goalies    []PlayerInfo `json:"goalies"`
}

// Record represents a team's record
type Record struct {
	Wins     int `json:"wins"`
	Losses   int `json:"losses"`
	OtLosses int `json:"otLosses"`
	Points   int `json:"points"`
}

// TeamStats represents team statistics
type TeamStats struct {
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	OtLosses     int    `json:"otLosses"`
	Points       int    `json:"points"`
	GoalsFor     int    `json:"goalsFor"`
	GoalsAgainst int    `json:"goalsAgainst"`
	GoalDiff     int    `json:"goalDiff"`
	HomeRecord   Record `json:"homeRecord"`
	AwayRecord   Record `json:"awayRecord"`
	L10Record    Record `json:"l10Record"`
}

// ScoreboardResponse represents the NHL score response
type ScoreboardResponse struct {
	GamesByDate []GamesByDate `json:"gamesByDate"`
}

// GamesByDate represents games grouped by date
type GamesByDate struct {
	Date  string `json:"date"`
	Games []Game `json:"games"`
}

// TVBroadcast represents a TV broadcast of a game
type TVBroadcast struct {
	ID             int    `json:"id"`
	Market         string `json:"market"`
	CountryCode    string `json:"countryCode"`
	Network        string `json:"network"`
	SequenceNumber int    `json:"sequenceNumber"`
}

// PeriodDescriptor represents information about a game period
type PeriodDescriptor struct {
	Number               int    `json:"number"`
	PeriodType           string `json:"periodType"`
	MaxRegulationPeriods int    `json:"maxRegulationPeriods"`
}

// FilteredScoreboardResponse represents the NHL score response for a specific date
type FilteredScoreboardResponse struct {
	Date  string `json:"date"`
	Games []Game `json:"games"`
}

// TeamsResponse represents the response from the teams endpoint
type TeamsResponse struct {
	Teams []TeamInfo `json:"teams"`
}

// TeamInfo represents detailed team information
type TeamInfo struct {
	ID           int           `json:"id"`
	Name         LanguageNames `json:"name"`
	Abbreviation string        `json:"abbreviation"`
	City         LanguageNames `json:"city"`
	TriCode      string        `json:"triCode"`
	FranchiseID  int           `json:"franchiseId"`
	Active       bool          `json:"active"`
}

// RosterResponse represents the response from the roster endpoint
type RosterResponse struct {
	Forwards   []PlayerInfo `json:"forwards"`
	Defensemen []PlayerInfo `json:"defensemen"`
	Goalies    []PlayerInfo `json:"goalies"`
}

// PlayerSearchResult represents a player found during search
type PlayerSearchResult struct {
	FirstName    NameInfo `json:"firstName"`
	LastName     NameInfo `json:"lastName"`
	Position     string   `json:"position"`
	JerseyNumber int      `json:"jerseyNumber"`
	TeamID       int      `json:"teamId"`
	TeamAbbrev   string   `json:"teamAbbrev"`
	PlayerID     int      `json:"playerId"`
}

// SkaterStats represents statistics for a skater
type SkaterStats struct {
	Assists            int     `json:"assists"`
	EvenStrengthGoals  int     `json:"evGoals"`
	EvenStrengthPoints int     `json:"evPoints"`
	FaceoffWinPct      float64 `json:"faceoffWinPct"`
	GameWinningGoals   int     `json:"gameWinningGoals"`
	GamesPlayed        int     `json:"gamesPlayed"`
	Goals              int     `json:"goals"`
	LastName           string  `json:"lastName"`
	OvertimeGoals      int     `json:"otGoals"`
	PenaltyMinutes     int     `json:"penaltyMinutes"`
	PlayerID           int     `json:"playerId"`
	PlusMinus          int     `json:"plusMinus"`
	Points             int     `json:"points"`
	PointsPerGame      float64 `json:"pointsPerGame"`
	PositionCode       string  `json:"positionCode"`
	PowerPlayGoals     int     `json:"ppGoals"`
	PowerPlayPoints    int     `json:"ppPoints"`
	SeasonID           int     `json:"seasonId"`
	ShortHandedGoals   int     `json:"shGoals"`
	ShortHandedPoints  int     `json:"shPoints"`
	ShootingPct        float64 `json:"shootingPct"`
	ShootsCatches      string  `json:"shootsCatches"`
	Shots              int     `json:"shots"`
	FullName           string  `json:"skaterFullName"`
	TeamAbbrev         string  `json:"teamAbbrevs"`
	TimeOnIcePerGame   float64 `json:"timeOnIcePerGame"`
}

// GoalieStats represents statistics for a goalie
type GoalieStats struct {
	Assists         int     `json:"assists"`
	GamesPlayed     int     `json:"gamesPlayed"`
	GamesStarted    int     `json:"gamesStarted"`
	FullName        string  `json:"goalieFullName"`
	Goals           int     `json:"goals"`
	GoalsAgainst    int     `json:"goalsAgainst"`
	GoalsAgainstAvg float64 `json:"goalsAgainstAverage"`
	LastName        string  `json:"lastName"`
	Losses          int     `json:"losses"`
	OvertimeLosses  int     `json:"otLosses"`
	PenaltyMinutes  int     `json:"penaltyMinutes"`
	PlayerID        int     `json:"playerId"`
	Points          int     `json:"points"`
	SavePctg        float64 `json:"savePct"`
	Saves           int     `json:"saves"`
	SeasonID        int     `json:"seasonId"`
	ShootsCatches   string  `json:"shootsCatches"`
	ShotsAgainst    int     `json:"shotsAgainst"`
	Shutouts        int     `json:"shutouts"`
	TeamAbbrev      string  `json:"teamAbbrevs"`
	TimeOnIce       int     `json:"timeOnIce"`
	Wins            int     `json:"wins"`
}

// StatsResponse represents a response containing player statistics
type StatsResponse struct {
	Data  []interface{} `json:"data"`
	Total int           `json:"total"`
}

// SkaterStatsResponse represents a response containing skater statistics
type SkaterStatsResponse struct {
	Data  []SkaterStats `json:"data"`
	Total int           `json:"total"`
}

// GoalieStatsResponse represents a response containing goalie statistics
type GoalieStatsResponse struct {
	Data  []GoalieStats `json:"data"`
	Total int           `json:"total"`
}

// StatsFilter represents filters for retrieving player statistics
type StatsFilter struct {
	GameType GameType
	SeasonID int // If 0, returns all seasons
}

// PlayerLandingResponse represents the response from the player landing page API
type PlayerLandingResponse struct {
	SeasonTotals []SeasonTotal `json:"seasonTotals"`
}

// SeasonTotal represents a player's stats for a single season
type SeasonTotal struct {
	Assists            int     `json:"assists,omitempty"`
	AvgTOI             string  `json:"avgToi,omitempty"`
	FaceoffWinningPctg float64 `json:"faceoffWinningPctg,omitempty"`
	GameTypeID         int     `json:"gameTypeId"`
	GameWinningGoals   int     `json:"gameWinningGoals,omitempty"`
	GamesPlayed        int     `json:"gamesPlayed"`
	Goals              int     `json:"goals,omitempty"`
	LeagueAbbrev       string  `json:"leagueAbbrev"`
	OTGoals            int     `json:"otGoals,omitempty"`
	PenaltyMinutes     int     `json:"pim,omitempty"`
	PlusMinus          int     `json:"plusMinus,omitempty"`
	Points             int     `json:"points"`
	PowerPlayGoals     int     `json:"powerPlayGoals,omitempty"`
	PowerPlayPoints    int     `json:"powerPlayPoints,omitempty"`
	Season             int     `json:"season"`
	ShootingPctg       float64 `json:"shootingPctg,omitempty"`
	ShorthandedGoals   int     `json:"shorthandedGoals,omitempty"`
	ShorthandedPoints  int     `json:"shorthandedPoints,omitempty"`
	Shots              int     `json:"shots,omitempty"`
	TeamName           struct {
		Default string `json:"default"`
	} `json:"teamName"`
}

// TeamScheduleResponse represents a team's schedule response
type TeamScheduleResponse struct {
	Games []ScheduleGame `json:"games"`
}

// ScheduleGame represents a game in a team's schedule
type ScheduleGame struct {
	ID             int            `json:"id"`
	Season         int            `json:"season"`
	GameType       int            `json:"gameType"`
	GameDate       string         `json:"gameDate"`
	StartTimeUTC   string         `json:"startTimeUTC"`
	VenueUTCOffset string         `json:"venueUTCOffset"`
	GameState      string         `json:"gameState"`
	HomeTeam       TeamInSchedule `json:"homeTeam"`
	AwayTeam       TeamInSchedule `json:"awayTeam"`
	GameCenterLink string         `json:"gameCenterLink"`
}

// TeamInSchedule represents a team in a schedule game
type TeamInSchedule struct {
	ID           int           `json:"id"`
	Name         LanguageNames `json:"name"`
	Abbreviation string        `json:"abbrev"`
	Score        int           `json:"score"`
}

// PlayerStatsResponse represents a response containing player statistics
type PlayerStatsResponse struct {
	Splits []SeasonTotal `json:"splits"`
}

// StandingsResponse represents the NHL standings response
type StandingsResponse struct {
	Standings []StandingsTeam `json:"standings"`
}

// StandingsTeam represents a team's standings information
type StandingsTeam struct {
	TeamName          LanguageNames `json:"teamName"`
	TeamAbbrev        TeamAbbrev    `json:"teamAbbrev"`
	Conference        string        `json:"conferenceName"`
	Division          string        `json:"divisionName"`
	Wins              int           `json:"wins"`
	Losses            int           `json:"losses"`
	OtLosses          int           `json:"otLosses"`
	RegulationWins    int           `json:"regulationWins"`
	Points            int           `json:"points"`
	GamesPlayed       int           `json:"gamesPlayed"`
	GoalsFor          int           `json:"goalFor"`
	GoalsAgainst      int           `json:"goalAgainst"`
	GoalDifferential  int           `json:"goalDifferential"`
	StreakCode        string        `json:"streakCode"`
	StreakCount       int           `json:"streakCount"`
	HomeGamesPlayed   int           `json:"homeGamesPlayed"`
	HomeWins          int           `json:"homeWins"`
	HomeLosses        int           `json:"homeLosses"`
	HomeOtLosses      int           `json:"homeOtLosses"`
	HomePoints        int           `json:"homePoints"`
	L10GamesPlayed    int           `json:"l10GamesPlayed"`
	L10Wins           int           `json:"l10Wins"`
	L10Losses         int           `json:"l10Losses"`
	L10OtLosses       int           `json:"l10OtLosses"`
	L10Points         int           `json:"l10Points"`
	PlaceInLeague     int           `json:"leagueSequence"`
	PlaceInConference int           `json:"conferenceSequence"`
	PlaceInDivision   int           `json:"divisionSequence"`
	WildCardSequence  int           `json:"wildcardSequence"`
	PointsPercentage  float64       `json:"pointsPercentage"`
}

// TeamAbbrev represents a team's abbreviation in different formats
type TeamAbbrev struct {
	Default string `json:"default"`
	French  string `json:"french,omitempty"`
}
