package actions

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
)

type RestAction struct{}

func (a RestAction) Execute(p *player.Player) (string, error) {
	if p.Energy >= 100 {
		fmt.Println("Вы достаточно бодры!")
	}

	oldHunger := p.Hunger
	oldThirst := p.Thirst

	p.Hunger = player.Clamp(p.Hunger-10, 0, 100)
	p.Thirst = player.Clamp(p.Thirst-15, 0, 100)

	return fmt.Sprintf("Вы отдохнули."+
		"Энергия: %d\n"+
		"Голод: %d -> %d\n"+
		"Жажда: %d -> %d\n",
		p.Energy, oldHunger, p.Hunger, oldThirst, p.Thirst), nil
}

func (a RestAction) Name() string {
	return "Отдых"
}

func (a RestAction) EnergyCost() int {
	return -30
}
