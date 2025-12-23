package game

import (
	"fmt"
	"math/rand"
)

type TimeOfDay string

const (
	Morning TimeOfDay = "Утро"
	Day     TimeOfDay = "День"
	Evening TimeOfDay = "Вечер"
	Night   TimeOfDay = "Ночь"
)

const (
	TurnsPerDayDefault = 4
	TurnsAtMorning     = 0
	TurnsAtDay         = 1
	TurnsAtEvening     = 2
	TurnsAtNight       = 3
)

func (g *Game) NewTurn() {
	g.Turn++

	switch g.Turn % g.TurnsPerDay {
	case TurnsAtMorning:
		g.TimeOfDay = Morning
	case TurnsAtDay:
		g.TimeOfDay = Day
	case TurnsAtEvening:
		g.TimeOfDay = Evening
	case TurnsAtNight:
		g.TimeOfDay = Night
	}

	if g.Turn%g.TurnsPerDay == TurnsAtMorning && g.Turn > 0 {
		g.Day++
		g.Player.Day = g.Day
	}

	g.Player.UpdateByTime(g.TimeOfDay)

	g.AddEvent(fmt.Sprintf("Ход: %d: %s, День %d",
		g.Turn, g.TimeOfDay, g.Day))

	if rand.Intn(100) < 25 {
		g.TriggerRandomEvent()
	}

	g.CheckGameOver()
}
