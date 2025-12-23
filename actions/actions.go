package actions

import (
	"fmt"
	"hunter_force/player"
)

type Action interface {
	Execute(p *player.Player) (string, error)
	Name() string
	EnergyCost() int
}

func (p *player.Player) PerformAction(action Action) bool {
	if action.EnergyCost() > 0 {
		if p.Energy < action.EnergyCost() {
			fmt.Println("Недостаточно энергии!")
			return false
		}

		p.Energy -= action.EnergyCost()
	}

	result, err := action.Execute(p)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return false
	}

	if action.EnergyCost() < 0 {
		energyGain := -action.EnergyCost()
		oldEnergy := p.Energy
		p.Energy = Clamp(p.Energy+energyGain, 0, 100)
		fmt.Printf("Восстановлено энергии: %d -> %d\n", oldEnergy, p.Energy)
	}

	fmt.Println(result)
	return true
}
