# NHL API Client Roadmap

## Current Features

### Schedule and Scores
- [x] Get Current Day's Schedule
- [x] Get Schedule by Date
- [x] Get Team Schedule
- [ ] Get Game Details
- [ ] Get Game Stats/Boxscore
- [ ] Get Live Game Updates
- [ ] Get Play-by-Play Data

### Teams
- [x] Get Team Details
- [x] Get Team Roster
- [ ] Get Team Stats
- [ ] Get Team History
- [ ] Get Head-to-Head Records
- [ ] Get Team Trends

### Players
- [x] Search Players
- [x] Get Player Stats (Regular Season/Playoffs)
- [x] Filter Stats by Season
- [ ] Get Player Game Logs
- [ ] Get Player Career Milestones
- [ ] Get Player Awards/Achievements
- [ ] Get Player Draft Information
- [ ] Get Player Advanced Stats

### Standings
- [ ] League Standings
- [ ] Conference Standings
- [ ] Division Standings
- [ ] Wild Card Standings
- [ ] Playoff Picture/Race

### League Information
- [ ] Get League Schedule (Key Dates)
- [ ] Get Draft Information
- [ ] Get League Leaders
- [ ] Get League Records
- [ ] Get Historical Data
- [ ] Get Recent Transactions
- [ ] Get Injury Reports
- [ ] Get News/Updates

### Statistical Analysis
- [ ] Team Advanced Stats
- [ ] Player Advanced Stats
- [ ] Situational Stats (PP, PK, etc.)
- [ ] Statistical Trends
- [ ] Custom Stat Filters

## Priority Queue

### High Priority
1. Team Schedule - Essential for following specific teams
2. Standings - Core feature for understanding team performance
3. Game Details/Boxscore - Detailed game information
4. League Leaders - Important for player performance context

### Medium Priority
1. Player Game Logs - Detailed player performance tracking
2. Team Stats - Comprehensive team statistics
3. Live Game Updates - Real-time game information
4. Recent Transactions - Team roster changes

### Low Priority
1. Historical Data - Past seasons and records
2. Advanced Stats - Detailed statistical analysis
3. Player Draft Information - Background information
4. News/Updates - Supplementary information

## Implementation Notes

### API Endpoints
- Base URL: `https://api.nhle.com/stats/rest/en`
- Schedule endpoint: `/schedule`
- Player endpoint: `/player`
- Team endpoint: `/team`
- Game endpoint: `/game`

### Data Models
- Need to implement models for:
  - Game Details
  - Standings
  - Team Stats
  - Advanced Stats

### Future Considerations
- Caching strategy for frequently accessed data
- Rate limiting implementation
- Error handling improvements
- Data validation enhancements
- Documentation updates 