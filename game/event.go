package game

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
	"math/rand"
)

func (g *Game) TriggerRandomEvent() {
	events := []func(*Game){
		func(g *Game) {
			g.AddEvent("Вы нашли заброшенную хижину с припасами!")
			g.Player.AddItem("аптечка", 1)
			g.Player.AddItem("вода", 2)
		},
		func(g *Game) {
			g.AddEvent("На вас напали дикие звери!")
			if g.Player.Campfire.IsLit {
				g.AddEvent("Но они испугались огня и убежали.")
			} else {
				g.Player.Health = player.Clamp(g.Player.Health-20, 0, 100)
				g.AddEvent("Вы потеряли 20 здоровья!")
			}
		},
		func(g *Game) {
			g.AddEvent("Пошел сильный дождь.")
			if !g.Player.Shelter.Exists {
				g.Player.Health = player.Clamp(g.Player.Health-10, 0, 100)
				g.AddEvent("Вы промокли и простудились (-10 здоровья)")
			} else {
				g.AddEvent("Укрытие защитило вас от дождя.")
			}
		},
	}

	event := events[rand.Intn(len(events))]
	event(g)
}

func (g *Game) AddEvent(event string) {
	g.EventLog = append(g.EventLog, event)

	if len(g.EventLog) > 10 {
		g.EventLog = g.EventLog[1:]
	}
}

func (g *Game) GetRecentEvents() string {
	if len(g.EventLog) == 0 {
		return "Событий пока нет"
	}

	result := "Последние события:\n"
	for i, event := range g.EventLog {
		result += fmt.Sprintf("%d. %s\n", i+1, event)
	}

	return result
}
