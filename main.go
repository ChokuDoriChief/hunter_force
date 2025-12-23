package main

import (
	"bufio"
	"fmt"
	game2 "hunter_force/game"
	"os"
	"strings"
)

func main() {
	game := game2.NewGame("Иван")

	// Стартовые предметы
	game.Player.AddItem("ягоды", 5)
	game.Player.AddItem("вода", 3)
	game.Player.AddItem("дрова", 10)
	game.Player.AddItem("нож", 1)
	game.Player.AddItem("удочка", 1)
	game.Player.AddItem("доски", 8)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== ВЫЖИВАНИЕ В ЛЕСУ ===")
	fmt.Println("Ваша цель: прожить 7 дней в лесу")
	fmt.Println("Начинаем игру...\n")

	for game.IsRunning {
		// Очистка экрана (простая)
		fmt.Print("\033[H\033[2J")

		// Статус
		fmt.Println(game.Status())

		// Действия
		game.ShowAvailableActions()

		// Выбор
		fmt.Print("\nВыберите действие (1-9): ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			game.Player.PerformAction(game2.RestAction{})
			game.NewTurn()
		case "2":
			if game.Player.HasItem("нож") {
				game.Player.PerformAction(game2.HuntAction{})
				game.NewTurn()
			} else {
				fmt.Println("Для охоты нужен нож!")
			}
		case "3":
			if game.Player.HasItem("удочка") {
				game.Player.PerformAction(game2.FishAction{})
				game.NewTurn()
			} else {
				fmt.Println("Для рыбалки нужна удочка!")
			}
		case "4":
			game.Player.PerformAction(game2.GatherAction{})
			game.NewTurn()
		case "5":
			if !game.Player.Shelter.Exists && game.Player.HasItem("доски") {
				game.Player.PerformAction(game2.BuildAction{})
				game.NewTurn()
			} else if game.Player.Shelter.Exists {
				fmt.Println("Убежище уже построено!")
			} else {
				fmt.Println("Нужно 5 досок для строительства!")
			}
		case "6":
			if !game.Player.Campfire.IsLit && game.Player.HasItem("дрова") {
				game.Player.PerformAction(game2.MakeCampfireAction{})
				game.NewTurn()
			} else if game.Player.Campfire.IsLit && game.Player.HasItem("дрова") {
				game.Player.PerformAction(game2.AddFuelCampfireAction{})
				game.NewTurn()
			} else {
				fmt.Println("Нужны дрова!")
			}
		case "7":
			// Меню предметов
			fmt.Println("\nВаши предметы:")
			for item, count := range game.Player.Inventory {
				fmt.Printf("- %s: %d шт.\n", item, count)
			}
			fmt.Print("Введите название предмета для использования (или 'назад'): ")
			scanner.Scan()
			itemName := strings.TrimSpace(scanner.Text())
			if itemName != "назад" {
				if game.Player.UseItem(itemName) {
					game.NewTurn()
				}
			}
		case "8":
			fmt.Println("Вы бездействуете...")
			game.NewTurn()
		case "9":
			game.IsRunning = false
			fmt.Println("Игра завершена.")
		default:
			fmt.Println("Неверный выбор.")
		}

		// Пауза
		if game.IsRunning {
			fmt.Println("\nНажмите Enter для продолжения...")
			scanner.Scan()
		}
	}

	// Итог
	fmt.Printf("\n=== ИГРА ОКОНЧЕНА ===\n")
	fmt.Printf("Вы прожили: %d дней\n", game.Day)
	if game.Player.IsAlive && game.Day >= 7 {
		fmt.Println("🏆 Вы выжили! Поздравляем!")
	} else if game.Player.IsAlive {
		fmt.Println("Вы сдались...")
	} else {
		fmt.Println("💀 Вы не выжили...")
	}
}
