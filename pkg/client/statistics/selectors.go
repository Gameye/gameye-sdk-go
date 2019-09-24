package statistics

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
)

func copyStats(in *map[string]interface{}) (out map[string]float64) {
	out = make(map[string]float64)
	for k, v := range *in {
		out[k] = v.(float64)
	}
	return out
}

func copyPlayers(in *map[string]interface{}) (out map[string]bool) {
	out = make(map[string]bool)
	for k, v := range *in {
		out[k] = v.(bool)
	}
	return out
}

func findPlayer(players []Player, playerKey string) (player Player, err error) {
	for _, v := range players {
		if v.PlayerKey == playerKey {
			return v, nil
		}
	}

	return Player{}, errors.New("could not find player with key")
}

// SelectRawStatistics returns a json string containing all the unprocessed statistics
func SelectRawStatistics(state State) (statisticsJson string, err error) {
	bytes, err := json.MarshalIndent(state.Statistics, "", "    ")
	if err != nil {
		err := errors.New("error marshaling statistics into Json")
		return "", err
	}
	return string(bytes), nil
}

// SelectPlayerList returns all the players
func SelectPlayerList(state State) (players []Player) {
	players = []Player{}
	rawStats := state.Statistics["statistic"].(map[string]interface{})

	for _, v := range rawStats["player"].(map[string]interface{}) {
		player := v.(map[string]interface{})
		playerStats := player["statistic"].(map[string]interface{})
		players = append(players, Player{
			PlayerKey: player["playerKey"].(string),
			Uid:       player["uid"].(string),
			Connected: player["connected"].(bool),
			Name:      player["name"].(string),
			Statistic: copyStats(&playerStats),
		})
	}

	sort.Slice(players, func(i, j int) bool {
		return strings.Compare(players[j].PlayerKey, players[i].PlayerKey) > 0
	})

	return players
}

// SelectTeamList returns all the teams
func SelectTeamList(state State) (teams []Team) {
	teams = []Team{}
	rawStats := state.Statistics["statistic"].(map[string]interface{})

	for _, v := range rawStats["team"].(map[string]interface{}) {
		team := v.(map[string]interface{})
		teamStats := team["statistic"].(map[string]interface{})
		teamPlayers := team["player"].(map[string]interface{})
		teams = append(teams, Team{
			TeamKey:   team["teamKey"].(string),
			Name:      team["name"].(string),
			Player:    copyPlayers(&teamPlayers),
			Statistic: copyStats(&teamStats),
		})
	}

	return teams
}

// SelectTeam returns a team for the given teamKey
func SelectTeam(state State, teamKey string) (team Team, err error) {
	teams := SelectTeamList(state)
	for _, v := range teams {
		if v.TeamKey == teamKey {
			return v, nil
		}
	}

	return Team{}, errors.New("could not find team with that key")
}

// SelectPlayerListForTeam returns the a slice of the players on the team matching the given teamKey
func SelectPlayerListForTeam(state State, teamKey string) (players []Player, err error) {
	players = []Player{}

	team, err := SelectTeam(state, teamKey)
	if err != nil {
		return players, err
	}

	allPlayers := SelectPlayerList(state)

	for playerKey, connected := range team.Player {
		if connected {
			found, err := findPlayer(allPlayers, playerKey)
			if err == nil {
				players = append(players, found)
			}
		}
	}

	sort.Slice(players, func(i, j int) bool {
		return strings.Compare(players[j].PlayerKey, players[i].PlayerKey) > 0
	})

	return players, err
}

// SelectPlayer returns a player with the given playerKey
func SelectPlayer(state State, playerKey string) (player Player, err error) {
	players := SelectPlayerList(state)
	return findPlayer(players, playerKey)
}

// SelectRounds returns the number of started rounds
func SelectRounds(state State) (rounds int) {
	rawStats := state.Statistics["statistic"].(map[string]interface{})
	rounds = int(rawStats["startedRounds"].(float64))
	return rounds
}
