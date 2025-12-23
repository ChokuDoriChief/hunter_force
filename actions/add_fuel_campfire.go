package actions

import (
	"fmt"
	"hunter_force/player"
)

type AddFuelCampfireAction struct{}

func (a AddFuelCampfireAction) Execute(p *player.Player) (string, error) {
	if !p.HasItem("дрова") {
		return "", fmt.Errorf("для пополнения топлива костра нужны дрова")
	}
	if p.Inventory["дрова"] < 2 {
		return "", fmt.Errorf("чтобы пополнить топливо костра нужно 2шт. дров. У вас %d", p.Inventory["дрова"])
	}

	p.Campfire.Fuel = 100
	return "Вы добавили дров в костер.", nil

}

func (a AddFuelCampfireAction) Name() string {
	return "Топливо для костра"
}

func (a AddFuelCampfireAction) EnergyCost() int {
	return 5
}
