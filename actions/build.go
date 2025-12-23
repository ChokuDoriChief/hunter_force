package actions

import (
	"fmt"
	"hunter_force/player"
)

type BuildAction struct{}

func (a BuildAction) Execute(p *player.Player) (string, error) {
	if !p.HasItem("доски") {
		return "", fmt.Errorf("для строительства дома нужны доски")
	}
	if p.Inventory["доски"] < 5 {
		return "", fmt.Errorf("для строительства дома не хватает досок, нужно 5шт. У Вас %d", p.Inventory["доски"])
	}

	p.Shelter.Exists = true

	oldThirst := p.Thirst
	oldHunger := p.Hunger
	p.Thirst = player.Clamp(p.Thirst-20, 0, 100)
	p.Hunger = player.Clamp(p.Hunger-20, 0, 100)

	return fmt.Sprintf("Вы построили жилище. Показатели комфорта повышены."+
		"Жажда: %d -> %d\n"+
		"Голод: %d -> %d\n", oldThirst, p.Thirst, oldHunger, p.Hunger), nil
}

func (a BuildAction) Name() string {
	return "Строительство"
}

func (a BuildAction) EnergyCost() int {
	return 30
}
