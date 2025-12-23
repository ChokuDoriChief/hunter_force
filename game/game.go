package game

import (
	"fmt"
	"hunter_force/player"
)

type Game struct {
	Player      *player.Player
	Day         int
	TimeOfDay   TimeOfDay
	Turn        int
	TurnsPerDay int
	IsRunning   bool
	EventLog    []string
}

func NewGame(playerName string) *Game {
	return &Game{Player: NewPlayer(playerName),
		Day:         1,
		TimeOfDay:   Morning,
		Turn:        0,
		TurnsPerDay: TurnsPerDayDefault,
		IsRunning:   true,
		EventLog:    make([]string, 0),
	}
}

func (g *Game) CheckGameOver() {
	if !g.Player.IsAlive {
		g.IsRunning = false
		g.AddEvent("Игрок не выжил...")
	}

	if g.Day > 7 {
		g.IsRunning = false
		g.AddEvent("Поздравляем! Вы прожили 7 дней!")
	}
}

func (g *Game) Status() string {
	turnsUntilNext := g.TurnsPerDay - (g.Turn % g.TurnsPerDay)
	if turnsUntilNext == g.TurnsPerDay {
		turnsUntilNext = 0
	}

	return fmt.Sprintf(
		"=== День %d, %s ===\n"+
			"Ход: %d | До смены времени: %d ход(ов)\n"+
			"%s\n"+
			"%s",
		g.Day, g.TimeOfDay, g.Turn, turnsUntilNext,
		g.Player.Status(),
		g.GetRecentEvents(),
	)
}

func (g *Game) ShowAvailableActions() {
	fmt.Println("\n=== Доступные действия ===")

	// Всегда доступно
	fmt.Println("1. Отдохнуть (Восст. энергию, -голод/жажда)")

	// Зависит от наличия предметов
	if g.Player.HasItem("нож") {
		fmt.Println("2. Охота (Нужен нож, дает мясо, тратит энергию)")
	} else {
		fmt.Println("2. Охота [НЕДОСТУПНО - нужен нож]")
	}

	if g.Player.HasItem("удочка") {
		fmt.Println("3. Рыбалка (Нужна удочка, дает рыбу)")
	} else {
		fmt.Println("3. Рыбалка [НЕДОСТУПНО - нужна удочка]")
	}

	fmt.Println("4. Собирательство (Случайные ресурсы)")

	if !g.Player.Shelter.Exists && g.Player.HasItem("доски") {
		fmt.Println("5. Построить убежище (Нужно 5 досок)")
	} else if g.Player.Shelter.Exists {
		fmt.Println("5. Убежище [УЖЕ ПОСТРОЕНО]")
	} else {
		fmt.Println("5. Построить убежище [НЕДОСТУПНО - нужны доски]")
	}

	if !g.Player.Campfire.IsLit && g.Player.HasItem("дрова") {
		fmt.Println("6. Разжечь костер (Нужно 3 дрова)")
	} else if g.Player.Campfire.IsLit {
		fmt.Println("6. Добавить топлива в костер (Нужно 2 дрова)")
	} else {
		fmt.Println("6. Разжечь костер [НЕДОСТУПНО - нужны дрова]")
	}

	fmt.Println("7. Использовать предмет")
	fmt.Println("8. Пропустить ход")
	fmt.Println("9. Выйти из игры")
}
