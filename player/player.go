package player

import (
	"fmt"
	"github.com/ChokuDoriChief/hunter_force/core"
	"hunter_force/game"
	"hunter_force/items"
)

type Player struct {
	Name      string
	Health    int
	Hunger    int
	Thirst    int
	Energy    int
	Morale    int
	Day       int
	Inventory map[string]int
	MaxWeight int
	IsAlive   bool
	Shelter   Shelter
	Campfire  Campfire
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, Health: 100, Hunger: 80, Thirst: 70, Energy: 90, Morale: 85, Day: 1, Inventory: make(map[string]int), MaxWeight: 20, IsAlive: true, Campfire: Campfire{false, 0}, Shelter: Shelter{false}}
}

type Shelter struct {
	Exists bool
}

type Campfire struct {
	IsLit bool
	Fuel  int
}

func (p *Player) PerformAction(action core.Action) bool {
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

func (p *Player) UpdateByTime(timeOfDay game.TimeOfDay) {
	p.Thirst = clamp(p.Thirst-5, 0, 100)
	p.Hunger = clamp(p.Hunger-5, 0, 100)
	p.Energy = clamp(p.Energy-5, 0, 100)

	switch timeOfDay {
	case game.Night:
		if !p.Campfire.IsLit && !p.Shelter.Exists {
			p.Morale = clamp(p.Morale-15, 0, 100)
			p.Health = clamp(p.Health-10, 0, 100)
			fmt.Println("Холодная ночь без укрытия подорвала ваше здоровье.")
		} else if p.Campfire.IsLit {
			p.Morale = clamp(p.Morale+10, 0, 100)
			fmt.Println("Тепло костра поднимает настроение.")
		}
	case game.Morning:
		p.Energy = clamp(p.Energy+15, 0, 100)
		if p.Energy > 50 {
			fmt.Println("Утренняя свежесть бодрит.")
		}
	}

	if p.Hunger < 20 {
		p.Health = clamp(p.Health-5, 0, 100)
		if p.Hunger < 10 {
			fmt.Println("Сильный голод сказывается на здоровье.")
		}
	}

	if p.Thirst < 20 {
		p.Health = clamp(p.Health-5, 0, 100)
		if p.Thirst < 10 {
			fmt.Println("Сильная жажда сказывается на здоровье.")
		}
	}

	if p.Campfire.IsLit {
		p.Campfire.Fuel = clamp(p.Campfire.Fuel-10, 0, 100)
		if p.Campfire.Fuel <= 0 {
			p.Campfire.IsLit = false
			fmt.Println("Костер потух")
		}
	}

	if p.Health <= 0 {
		p.IsAlive = false
	}
}

func (p *Player) Status() string {
	inventoryInfo := "Инвентарь пуст"
	if len(p.Inventory) > 0 {
		inventoryInfo = "Инвентарь:\n"
		items := GetAvailableItems()
		for itemName, quantity := range p.Inventory {
			item := items[itemName]
			inventoryInfo += fmt.Sprintf("- %s: %d шт. (вес %d)\n", item.Name, quantity, item.Weight*quantity)
		}
	}
	return fmt.Sprintf(
		"Игрок %s (День %d)\n"+
			"♥ Здоровье: %d/100\n"+
			"🍖 Голод: %d/100\n"+
			"💧 Жажда: %d/100\n"+
			"⚡ Энергия: %d/100\n"+
			"😊 Мораль: %d/100\n"+
			"📦 %s (%d/%d кг)\n",
		p.Name, p.Day,
		p.Health, p.Hunger, p.Thirst, p.Energy, p.Morale,
		inventoryInfo, p.GetInventoryWeight(), p.MaxWeight)
}

func (p *Player) UseItem(itemName string) bool {
	if !p.HasItem(itemName) {
		return false
	}

	availableItems := GetAvailableItems()
	item, exists := availableItems[itemName]
	if !exists {
		fmt.Printf("Предмет %s не существует в игре\n", itemName)
		return false
	}

	success := false
	switch item.Type {
	case items.FoodType:
		oldHunger := p.Hunger
		p.Hunger = clamp(p.Hunger+item.Value, 0, 100)
		fmt.Printf("Съел %s. Голод: %d -> %d\n",
			itemName, oldHunger, p.Hunger)
		success = true

	case items.ToolType:
		fmt.Printf("Использовал %s: %s\n", itemName, item.Description)
		// TODO: написать реализацию кейса
		success = true

	case items.MedicineType:
		if p.Health >= 100 {
			fmt.Println("Лечение не требуется!")
			return false
		}
		oldHealth := p.Health
		p.Health = Clamp(p.Health+item.Value, 0, 100)
		fmt.Printf("Использовал %s. Здоровье %d -> %d\n",
			itemName, oldHealth, p.Health)
		success = true

	case items.ResourceType:
		fmt.Printf("%s нельзя использовать напрямую\n", itemName)
		return false

	default:
		fmt.Printf("Неизвестный тип предмета %s\n", item.Type)
	}

	if success {
		p.RemoveItem(itemName, 1)
	}

	return success
}

func Clamp(value, min, max int) int { // вспомогательная функция-ограничитель значения
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
