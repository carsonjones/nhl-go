# NHL API Client Roadmap

## Current Features

### Schedule and Scores
- [x] Get Current Day's Schedule
- [x] Get Schedule by Date
- [x] Get Team Schedule
- [x] Get Game Details
- [x] Get Game Stats/Boxscore
- [x] Get Play-by-Play Data
- [ ] Get Live Game Updates

### Teams
- [x] Get Team Details
- [x] Get Team Roster
- [x] Get Team Stats
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
- [x] League Standings
- [x] Conference Standings
- [x] Division Standings
- [x] Wild Card Standings
- [ ] Add playoff indicators to standings
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
1. Live Game Updates - Real-time game information
2. Team History and Head-to-Head Records
3. League Leaders - Important for player performance context
4. Player Game Logs - Detailed player performance tracking

### Medium Priority
1. Team Trends - Performance analysis over time
2. Recent Transactions - Team roster changes
3. Playoff Picture/Race - Playoff implications
4. Advanced Stats - Detailed statistical analysis

### Low Priority
1. Historical Data - Past seasons and records
2. Player Draft Information - Background information
3. News/Updates - Supplementary information
4. Injury Reports - Player availability

## Implementation Notes

### API Endpoints
- Base URL: `https://api-web.nhle.com/v1`
- Schedule endpoint: `/schedule`
- Player endpoint: `/player`
- Team endpoint: `/team`
- Game endpoint: `/gamecenter`

### Data Models
- Need to implement models for:
  - Advanced Stats
  - Historical Data
  - League Leaders
  - Live Updates

### Future Considerations
- Caching strategy for frequently accessed data
- Rate limiting implementation
- Error handling improvements
- Data validation enhancements
- Documentation updates 