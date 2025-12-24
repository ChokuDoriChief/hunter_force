package player

import "fmt"

func (p *Player) GetInventoryWeight() int {
	items := GetAvailableItems()
	currentWeight := 0
	currentItem := 0
	for k, v := range p.Inventory {
		currentItem = items[k].Weight    // Вес одной единицы груза
		currentWeight += currentItem * v // Общий вес текущей позиции в инвентаре
	}

	return currentWeight
}

func (p *Player) AddItem(itemName string, quantity int) bool {
	items := GetAvailableItems() // Общий список предметов
	item, ok := items[itemName]  // Проверка на существование
	if !ok {
		return false

	}

	addItemWeight := item.Weight * quantity                 // Вес добавляемого предмета
	if p.MaxWeight-p.GetInventoryWeight() < addItemWeight { // Проверка на перевес
		fmt.Println("Не хватает места в инвентаре")
		return false
	}
	p.Inventory[itemName] += quantity
	return true
}

func (p *Player) RemoveItem(itemName string, quantity int) bool {
	items := GetAvailableItems() // Общий список предметов
	_, ok := items[itemName]     // Проверка на существование
	if !ok {
		return false
	}

	countItem, ok := p.Inventory[itemName]
	if !ok {
		return false
	}
	if countItem < quantity {
		fmt.Printf("Недостаточное количество %s для удаления", itemName)
		return false
	}

	p.Inventory[itemName] = countItem - quantity // Меняем значение в мапе после удаления
	if p.Inventory[itemName] == 0 {
		delete(p.Inventory, itemName)
	}

	return true
}

func (p *Player) HasItem(itemName string) bool {
	_, ok := p.Inventory[itemName]
	if !ok {
		return false
	}
	return true
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
