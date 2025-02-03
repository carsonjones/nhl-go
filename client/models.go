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
	Clock             GameClock        `json:"clock"`
	Situation         *GameSituation   `json:"situation,omitempty"`
}

// Venue represents a game venue
type Venue struct {
	Default string `json:"default"`
}

// VenueLocation represents a game venue location
type VenueLocation struct {
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
	ShotsOnGoal              int           `json:"sog"`
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
	Goals       int            `json:"goals"`
	ShotsOnGoal int            `json:"sog"`
	FaceoffPct  float64        `json:"faceoffPct"`
	Hits        int            `json:"hits"`
	PIM         int            `json:"pim"`
	PowerPlay   PowerPlayStats `json:"powerPlay"`
}

// ScoreboardResponse represents the NHL score response
type ScoreboardResponse struct {
	FocusedDate      string        `json:"focusedDate"`
	FocusedDateCount int           `json:"focusedDateCount"`
	GamesByDate      []GamesByDate `json:"gamesByDate"`
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

// GameDetails represents detailed information about a specific game
type GameDetails struct {
	ID           int           `json:"id"`
	GameType     int           `json:"gameType"`
	Season       int           `json:"season"`
	GameDate     string        `json:"gameDate"`
	StartTimeUTC string        `json:"startTimeUTC"`
	Venue        Venue         `json:"venue"`
	GameState    string        `json:"gameState"`
	HomeTeam     DetailedTeam  `json:"homeTeam"`
	AwayTeam     DetailedTeam  `json:"awayTeam"`
	Clock        GameClock     `json:"clock"`
	TVBroadcasts []TVBroadcast `json:"tvBroadcasts"`
	Summary      GameSummary   `json:"summary"`
	ThreeStars   []StarPlayer  `json:"threeStars"`
}

// DetailedTeam represents a team with detailed game information
type DetailedTeam struct {
	ID                       int           `json:"id"`
	CommonName               LanguageNames `json:"commonName"`
	Abbrev                   string        `json:"abbrev"`
	PlaceName                LanguageNames `json:"placeName"`
	PlaceNameWithPreposition LanguageNames `json:"placeNameWithPreposition"`
	Score                    int           `json:"score"`
	ShotsOnGoal              int           `json:"sog"`
	Logo                     string        `json:"logo"`
	DarkLogo                 string        `json:"darkLogo"`
}

// GameSummary represents the scoring and penalty summary
type GameSummary struct {
	Scoring   []PeriodSummary   `json:"scoring"`
	Shootout  []interface{}     `json:"shootout"`
	Penalties []PeriodPenalties `json:"penalties"`
}

// PeriodSummary represents scoring information for a period
type PeriodSummary struct {
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	Goals            []GoalEvent      `json:"goals"`
}

// PeriodPenalties represents penalties for a period
type PeriodPenalties struct {
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	Penalties        []PenaltyEvent   `json:"penalties"`
}

// GoalEvent represents a goal scored in the game
type GoalEvent struct {
	SituationCode     string         `json:"situationCode"`
	Strength          string         `json:"strength"`
	PlayerID          int            `json:"playerId"`
	FirstName         LanguageNames  `json:"firstName"`
	LastName          LanguageNames  `json:"lastName"`
	Name              LanguageNames  `json:"name"`
	TeamAbbrev        LanguageNames  `json:"teamAbbrev"`
	TimeInPeriod      string         `json:"timeInPeriod"`
	ShotType          string         `json:"shotType"`
	GoalModifier      string         `json:"goalModifier"`
	AwayScore         int            `json:"awayScore"`
	HomeScore         int            `json:"homeScore"`
	LeadingTeamAbbrev *LanguageNames `json:"leadingTeamAbbrev,omitempty"`
	Assists           []AssistEvent  `json:"assists"`
}

// AssistEvent represents an assist on a goal
type AssistEvent struct {
	PlayerID      int           `json:"playerId"`
	FirstName     LanguageNames `json:"firstName"`
	LastName      LanguageNames `json:"lastName"`
	Name          LanguageNames `json:"name"`
	AssistsToDate int           `json:"assistsToDate"`
	SweaterNumber int           `json:"sweaterNumber"`
}

// PenaltyEvent represents a penalty called in the game
type PenaltyEvent struct {
	TimeInPeriod      string        `json:"timeInPeriod"`
	Type              string        `json:"type"`
	Duration          int           `json:"duration"`
	CommittedByPlayer string        `json:"committedByPlayer"`
	TeamAbbrev        LanguageNames `json:"teamAbbrev"`
	DrawnBy           string        `json:"drawnBy"`
	DescKey           string        `json:"descKey"`
}

// StarPlayer represents a player selected as one of the three stars
type StarPlayer struct {
	Star       int           `json:"star"`
	PlayerID   int           `json:"playerId"`
	TeamAbbrev string        `json:"teamAbbrev"`
	Headshot   string        `json:"headshot"`
	Name       LanguageNames `json:"name"`
	SweaterNo  int           `json:"sweaterNo"`
	Position   string        `json:"position"`
	Goals      int           `json:"goals,omitempty"`
	Assists    int           `json:"assists,omitempty"`
	Points     int           `json:"points,omitempty"`
	SavePctg   float64       `json:"savePctg,omitempty"`
}

// TeamGameStats represents a team's stats for a specific game
type TeamGameStats struct {
	ID           int             `json:"id"`
	Name         LanguageNames   `json:"name"`
	Abbreviation string          `json:"abbrev"`
	Score        int             `json:"score"`
	ShotsOnGoal  int             `json:"sog"`
	FaceoffPct   float64         `json:"faceoffPct"`
	PowerPlay    PowerPlayStats  `json:"powerPlay"`
	Scratches    []PlayerBrief   `json:"scratches"`
	Leaders      GameTeamLeaders `json:"leaders"`
}

// PowerPlayStats represents power play statistics
type PowerPlayStats struct {
	Goals         int     `json:"goals"`
	Opportunities int     `json:"opportunities"`
	Percentage    float64 `json:"percentage"`
}

// GameTeamLeaders represents team leaders in various categories for a game
type GameTeamLeaders struct {
	Goals   []PlayerBrief `json:"goals"`
	Assists []PlayerBrief `json:"assists"`
	Points  []PlayerBrief `json:"points"`
}

// PlayerBrief represents basic player information
type PlayerBrief struct {
	ID       int           `json:"id"`
	Name     LanguageNames `json:"name"`
	Position string        `json:"position"`
	Number   string        `json:"sweaterNumber"`
}

// GameClock represents the game clock information
type GameClock struct {
	TimeRemaining    string `json:"timeRemaining"`
	SecondsRemaining int    `json:"secondsRemaining"`
	Running          bool   `json:"running"`
	InIntermission   bool   `json:"inIntermission"`
}

// BoxscoreResponse represents the boxscore data for a game
type BoxscoreResponse struct {
	ID                int             `json:"id"`
	Season            int             `json:"season"`
	GameType          int             `json:"gameType"`
	GameDate          string          `json:"gameDate"`
	StartTimeUTC      string          `json:"startTimeUTC"`
	Venue             Venue           `json:"venue"`
	GameState         string          `json:"gameState"`
	HomeTeam          DetailedTeam    `json:"homeTeam"`
	AwayTeam          DetailedTeam    `json:"awayTeam"`
	PlayerByGameStats PlayerGameStats `json:"playerByGameStats"`
}

// PlayerGameStats represents player statistics for both teams
type PlayerGameStats struct {
	HomeTeam TeamPlayerStats `json:"homeTeam"`
	AwayTeam TeamPlayerStats `json:"awayTeam"`
}

// TeamPlayerStats represents player statistics for a team
type TeamPlayerStats struct {
	Forwards []PlayerStats     `json:"forwards"`
	Defense  []PlayerStats     `json:"defense"`
	Goalies  []GoalieGameStats `json:"goalies"`
}

// PlayerStats represents statistics for a skater
type PlayerStats struct {
	PlayerID          int           `json:"playerId"`
	SweaterNumber     int           `json:"sweaterNumber"`
	Name              LanguageNames `json:"name"`
	Position          string        `json:"position"`
	Goals             int           `json:"goals"`
	Assists           int           `json:"assists"`
	Points            int           `json:"points"`
	PlusMinus         int           `json:"plusMinus"`
	PIM               int           `json:"pim"`
	Hits              int           `json:"hits"`
	PowerPlayGoals    int           `json:"powerPlayGoals"`
	SOG               int           `json:"sog"`
	FaceoffWinningPct float64       `json:"faceoffWinningPctg"`
	TOI               string        `json:"toi"`
	BlockedShots      int           `json:"blockedShots"`
	Shifts            int           `json:"shifts"`
	Giveaways         int           `json:"giveaways"`
	Takeaways         int           `json:"takeaways"`
}

// GoalieGameStats represents statistics for a goalie
type GoalieGameStats struct {
	PlayerID                 int           `json:"playerId"`
	SweaterNumber            int           `json:"sweaterNumber"`
	Name                     LanguageNames `json:"name"`
	Position                 string        `json:"position"`
	EvenStrengthShotsAgainst string        `json:"evenStrengthShotsAgainst"`
	PowerPlayShotsAgainst    string        `json:"powerPlayShotsAgainst"`
	ShorthandedShotsAgainst  string        `json:"shorthandedShotsAgainst"`
	SaveShotsAgainst         string        `json:"saveShotsAgainst"`
	SavePctg                 float64       `json:"savePctg"`
	EvenStrengthGoalsAgainst int           `json:"evenStrengthGoalsAgainst"`
	PowerPlayGoalsAgainst    int           `json:"powerPlayGoalsAgainst"`
	ShorthandedGoalsAgainst  int           `json:"shorthandedGoalsAgainst"`
	PIM                      int           `json:"pim"`
	GoalsAgainst             int           `json:"goalsAgainst"`
	TOI                      string        `json:"toi"`
	Starter                  bool          `json:"starter"`
	Decision                 string        `json:"decision,omitempty"`
	ShotsAgainst             int           `json:"shotsAgainst"`
	Saves                    int           `json:"saves"`
}

// PeriodStats represents statistics for a game period
type PeriodStats struct {
	PeriodNumber int              `json:"periodNumber"`
	HomeScore    int              `json:"homeScore"`
	AwayScore    int              `json:"awayScore"`
	Goals        []GoalSummary    `json:"goals"`
	Penalties    []PenaltySummary `json:"penalties"`
}

// GoalSummary represents a goal scored in the game
type GoalSummary struct {
	Period       int           `json:"period"`
	TimeInPeriod string        `json:"timeInPeriod"`
	GoalType     string        `json:"goalType"`
	ScoredBy     PlayerBrief   `json:"scoredBy"`
	AssistedBy   []PlayerBrief `json:"assistedBy"`
}

// PenaltySummary represents a penalty called in the game
type PenaltySummary struct {
	Period       int         `json:"period"`
	TimeInPeriod string      `json:"timeInPeriod"`
	Type         string      `json:"type"`
	Minutes      int         `json:"minutes"`
	Player       PlayerBrief `json:"player"`
}

// PenaltyBoxItem represents a player currently in the penalty box
type PenaltyBoxItem struct {
	Player        PlayerBrief `json:"player"`
	TimeRemaining string      `json:"timeRemaining"`
	Type          string      `json:"type"`
}

// Official represents a game official
type Official struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// Coach represents a team coach
type Coach struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

// PlayByPlayResponse represents play-by-play data for a game
type PlayByPlayResponse struct {
	Plays       []PlayEvent  `json:"plays"`
	RosterSpots []RosterSpot `json:"rosterSpots"`
}

// RosterSpot represents a player in the game roster
type RosterSpot struct {
	TeamID        int           `json:"teamId"`
	PlayerID      int           `json:"playerId"`
	FirstName     LanguageNames `json:"firstName"`
	LastName      LanguageNames `json:"lastName"`
	SweaterNumber int           `json:"sweaterNumber"`
	PositionCode  string        `json:"positionCode"`
	Headshot      string        `json:"headshot"`
}

// PlayEvent represents a single event in the game
type PlayEvent struct {
	EventID          int              `json:"eventId"`
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	TimeInPeriod     string           `json:"timeInPeriod"`
	TimeRemaining    string           `json:"timeRemaining"`
	SituationCode    string           `json:"situationCode"`
	TypeCode         int              `json:"typeCode"`
	TypeDescKey      string           `json:"typeDescKey"`
	Details          EventDetails     `json:"details"`
}

// EventDetails represents details about a play event
type EventDetails struct {
	EventOwnerTeamID    int     `json:"eventOwnerTeamId,omitempty"`
	XCoord              float64 `json:"xCoord,omitempty"`
	YCoord              float64 `json:"yCoord,omitempty"`
	ZoneCode            string  `json:"zoneCode,omitempty"`
	ShotType            string  `json:"shotType,omitempty"`
	ShootingPlayerID    int     `json:"shootingPlayerId,omitempty"`
	GoalieInNetID       int     `json:"goalieInNetId,omitempty"`
	BlockingPlayerID    int     `json:"blockingPlayerId,omitempty"`
	HittingPlayerID     int     `json:"hittingPlayerId,omitempty"`
	HitteePlayerID      int     `json:"hitteePlayerId,omitempty"`
	WinningPlayerID     int     `json:"winningPlayerId,omitempty"`
	LosingPlayerID      int     `json:"losingPlayerId,omitempty"`
	Reason              string  `json:"reason,omitempty"`
	TypeCode            string  `json:"typeCode,omitempty"`
	DescKey             string  `json:"descKey,omitempty"`
	Duration            int     `json:"duration,omitempty"`
	CommittedByPlayerID int     `json:"committedByPlayerId,omitempty"`
	DrawnByPlayerID     int     `json:"drawnByPlayerId,omitempty"`
	AwaySOG             int     `json:"awaySOG,omitempty"`
	HomeSOG             int     `json:"homeSOG,omitempty"`
	ScoringPlayerID     int     `json:"scoringPlayerId,omitempty"`
	ScoringPlayerTotal  int     `json:"scoringPlayerTotal,omitempty"`
	Assist1PlayerID     int     `json:"assist1PlayerId,omitempty"`
	Assist1PlayerTotal  int     `json:"assist1PlayerTotal,omitempty"`
	Assist2PlayerID     int     `json:"assist2PlayerId,omitempty"`
	Assist2PlayerTotal  int     `json:"assist2PlayerTotal,omitempty"`
}

// GameStoryResponse represents the game story/narrative
type GameStoryResponse struct {
	GameID            int              `json:"id"`
	Season            int              `json:"season"`
	GameType          int              `json:"gameType"`
	GameDate          string           `json:"gameDate"`
	Venue             Venue            `json:"venue"`
	VenueLocation     VenueLocation    `json:"venueLocation"`
	StartTimeUTC      string           `json:"startTimeUTC"`
	EasternUTCOffset  string           `json:"easternUTCOffset"`
	VenueUTCOffset    string           `json:"venueUTCOffset"`
	VenueTimezone     string           `json:"venueTimezone"`
	TVBroadcasts      []TVBroadcast    `json:"tvBroadcasts"`
	GameState         string           `json:"gameState"`
	GameScheduleState string           `json:"gameScheduleState"`
	HomeTeam          Team             `json:"homeTeam"`
	AwayTeam          Team             `json:"awayTeam"`
	ShootoutInUse     bool             `json:"shootoutInUse"`
	MaxPeriods        int              `json:"maxPeriods"`
	RegPeriods        int              `json:"regPeriods"`
	OtInUse           bool             `json:"otInUse"`
	TiesInUse         bool             `json:"tiesInUse"`
	Summary           GameSummary      `json:"summary"`
	PeriodDescriptor  PeriodDescriptor `json:"periodDescriptor"`
	Clock             GameClock        `json:"clock"`
}

// GameSituation represents the current game situation
type GameSituation struct {
	HomeTeam struct {
		Abbrev                string   `json:"abbrev"`
		SituationDescriptions []string `json:"situationDescriptions"`
		Strength              int      `json:"strength"`
	} `json:"homeTeam"`
	AwayTeam struct {
		Abbrev   string `json:"abbrev"`
		Strength int    `json:"strength"`
	} `json:"awayTeam"`
	SituationCode    string `json:"situationCode"`
	TimeRemaining    string `json:"timeRemaining"`
	SecondsRemaining int    `json:"secondsRemaining"`
}
