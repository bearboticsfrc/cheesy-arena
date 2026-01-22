// Copyright 2022 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)

package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreSummaryDetermineMatchStatus(t *testing.T) {
	redScoreSummary := &ScoreSummary{Score: 10}
	blueScoreSummary := &ScoreSummary{Score: 10}
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	redScoreSummary.Score = 11
	assert.Equal(t, RedWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, RedWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	blueScoreSummary.Score = 12
	assert.Equal(t, BlueWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, BlueWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	redScoreSummary.Score = 175
	redScoreSummary.NumOpponentMajorFouls = 0
	redScoreSummary.AutoFuelPoints = 40
	redScoreSummary.AutoClimbPoints = 15
	redScoreSummary.ActiveFuelPoints = 80
	redScoreSummary.TeleopClimbPoints = 40
	redScoreSummary.FoulPoints = 0
	redScoreSummary.EnergizedRankingPoint = true
	redScoreSummary.SuperchargedRankingPoint = false
	redScoreSummary.TraversalRankingPoint = true

	blueScoreSummary.Score = 150
	blueScoreSummary.NumOpponentMajorFouls = 0
	blueScoreSummary.AutoFuelPoints = 20
	blueScoreSummary.AutoClimbPoints = 15
	blueScoreSummary.ActiveFuelPoints = 50
	blueScoreSummary.TeleopClimbPoints = 10
	blueScoreSummary.FoulPoints = 0
	blueScoreSummary.EnergizedRankingPoint = false
	blueScoreSummary.SuperchargedRankingPoint = false
	blueScoreSummary.TraversalRankingPoint = false

	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, RedWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	blueScoreSummary.NumOpponentMajorFouls = 12
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, BlueWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	redScoreSummary.NumOpponentMajorFouls = 12
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, RedWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	blueScoreSummary.AutoPoints = 12
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, BlueWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))

	redScoreSummary.AutoPoints = 12
	assert.Equal(t, TieMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, false))
	assert.Equal(t, RedWonMatch, DetermineMatchStatus(redScoreSummary, blueScoreSummary, true))
}
