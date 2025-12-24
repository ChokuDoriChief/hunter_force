package actions

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
	"math/rand"
)

type GatherAction struct{}

func (a GatherAction) Execute(p *player.Player) (string, error) {
	chance := rand.Intn(100)
	switch {
	case chance < 20:
		p.AddItem("мясо", 2)
		fmt.Println("вы нашли мясо 2шт.")
	case chance < 50:
		p.AddItem("ягоды", 3)
		fmt.Println("вы нашли ягоды 3шт.")
	default:
		p.AddItem("дрова", 2)
		fmt.Println("вы нашли дрова 2шт.")
	}

	oldThirst := p.Thirst
	oldHunger := p.Hunger
	p.Thirst = player.Clamp(p.Thirst-10, 0, 100)
	p.Hunger = player.Clamp(p.Hunger-15, 0, 100)

	return fmt.Sprintf("Вы собрали ресурсы."+
		"Жажда: %d -> %d\n"+
		"Голод: %d -> %d\n", oldThirst, p.Thirst, oldHunger, p.Hunger), nil
}

func (a GatherAction) Name() string {
	return "Собирательство"
}

func (a GatherAction) EnergyCost() int {
	return 15
}
