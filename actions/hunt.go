package actions

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
	"math/rand"
)

type HuntAction struct{}

func (h HuntAction) Execute(p *player.Player) (string, error) {
	if !p.HasItem("нож") {
		return "", fmt.Errorf("для охоты нужен нож")
	}

	if rand.Intn(100) < 30 {
		p.RemoveItem("нож", 1)
		p.Morale = player.Clamp(p.Morale-5, 0, 100)
		return "Охота не удалась: добыча убежала", nil
	}

	p.AddItem("мясо", 2)
	oldMorale := p.Morale
	oldThirst := p.Thirst
	oldHunger := p.Hunger
	p.Morale = player.Clamp(p.Morale+10, 0, 100)
	p.Thirst = player.Clamp(p.Thirst-10, 0, 100)
	p.Hunger = player.Clamp(p.Hunger-10, 0, 100)

	return fmt.Sprintf(
		"Вы удачно поохотились. Добыто 2 ед. мяса"+
			"Мораль: %d -> %d\n"+
			"Жажда: %d -> %d\n"+
			"Голод: %d -> %d\n", oldMorale, p.Morale, oldThirst, p.Thirst, oldHunger, p.Hunger), nil
}

func (h HuntAction) Name() string {
	return "Охота"
}

func (h HuntAction) EnergyCost() int {
	return 20
}
