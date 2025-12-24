package core

import "github.com/ChokuDoriChief/hunter_force/player"

type Action interface {
	Execute(p *player.Player) (string, error)
	Name() string
	EnergyCost() int
}

type ItemType string

const (
	FoodType     ItemType = "еда"
	ToolType     ItemType = "инструмент"
	ResourceType ItemType = "ресурс"
	MedicineType ItemType = "лекарство"
)

type TimeOfDay string

const (
	Morning TimeOfDay = "Утро"
	Day     TimeOfDay = "День"
	Evening TimeOfDay = "Вечер"
	Night   TimeOfDay = "Ночь"
)
