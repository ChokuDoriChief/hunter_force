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
