package actions

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/player"
	"math/rand"
)

type FishAction struct{}

func (a FishAction) Execute(p *player.Player) (string, error) {
	if !p.HasItem("удочка") {
		return "", fmt.Errorf("для рыбалки нужна удочка")
	}

	if rand.Intn(100) < 30 {
		p.RemoveItem("удочка", 1)
		fmt.Println("удочка сломалась")
	}

	p.AddItem("рыба", 3)
	oldMorale := p.Morale
	oldThirst := p.Thirst
	oldHunger := p.Hunger
	p.Morale = player.Clamp(p.Morale+5, 0, 100)
	p.Thirst = player.Clamp(p.Thirst-12, 0, 100)
	p.Hunger = player.Clamp(p.Hunger-7, 0, 100)

	return fmt.Sprintf("Вы удачно порыбачили. Добыто 3 рыбы."+
		"Мораль: %d -> %d\n"+
		"Жажда: %d -> %d\n"+
		"Голод: %d -> %d\n", oldMorale, p.Morale, oldThirst, p.Thirst, oldHunger, p.Hunger), nil
}

func (a FishAction) Name() string {
	return "Рыбалка"
}

func (a FishAction) EnergyCost() int {
	return 7
}
