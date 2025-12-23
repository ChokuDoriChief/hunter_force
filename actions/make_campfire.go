package actions

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
)

type MakeCampfireAction struct{}

func (a MakeCampfireAction) Execute(p *player.Player) (string, error) {
	if !p.HasItem("дрова") {
		return "", fmt.Errorf("для костра нужны дрова")
	}
	if p.Inventory["дрова"] < 3 {
		return "", fmt.Errorf("для создания костра нужно 3шт. дров. У вас %d", p.Inventory["дрова"])
	}
	p.RemoveItem("дрова", 3)
	p.Campfire.IsLit = true
	p.Campfire.Fuel = 100
	//todo: нужно доделать чтобы со временем костер сгорал время крч надо реализовать
	return "Вы разожгли костер. Показатели улучшены!", nil

}

func (a MakeCampfireAction) Name() string {
	return "Создание костра"
}

func (a MakeCampfireAction) EnergyCost() int {
	return 10
}
