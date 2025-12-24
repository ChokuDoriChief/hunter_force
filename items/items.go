package items

type Item struct {
	Name        string
	Type        ItemType
	Value       int // На сколько восстанавливает/полезность
	Weight      int // Вес(ограничение инвентаря)
	Description string
}

var availableItems = map[string]Item{
	"ягоды": {
		"ягоды",
		FoodType,
		15, // восстанавливает голод
		1,
		"Съедобные лесные ягоды",
	},
	"вода": {
		"вода",
		FoodType,
		20, // утоляет жажду
		2,
		"Чистая вода из ручья",
	},
	"нож": {
		"нож",
		ToolType,
		5, // помогает в охоте
		3,
		"Охотничий нож",
	},
	"аптечка": {
		"аптечка",
		MedicineType,
		40, //  восстанавливает здоровье
		2,
		"Набор для лечения ран",
	},
	"дрова": {
		"дрова",
		ResourceType,
		10, // для костра
		5,
		"Сухие дрова для костра",
	},
	"мясо": {
		"мясо",
		FoodType,
		25,
		2,
		"свежее мясо",
	},
	"рыба": {
		"рыба",
		FoodType,
		10,
		2,
		"свежевыловленная сырая рыба",
	},
	"готовая рыба": {
		"готовая рыба",
		FoodType,
		25,
		2,
		"жаренная рыба",
	},
	"удочка": {
		"удочка",
		ToolType,
		5,
		4,
		"обычная удочка рыбака",
	},
	"доски": {
		"доски",
		ResourceType,
		12,
		6,
		"доски для создания убежища",
	},
	// TODO: добавить дополнительные предметы
}

func Get(name string) (Item, bool) {
	item, exists := availableItems[name]
	return item, exists
}

func GetAllItems() map[string]Item {
	return availableItems
}
