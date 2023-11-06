package basics

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.maxHealth = player.health
	}
	fmt.Println(player.name, "Add", amount, "Health->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "Health->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.maxEnergy = player.energy
	}
	fmt.Println(player.name, "Add", amount, "Energy->", player.energy)
}

func (player *Player) consumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Damage", amount, "Energy->", player.energy)
}

func ReceiverFunc_1() {
	fmt.Println("Hello Receiver_1")

	player := Player{
		name:      "Knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(10)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(20)
	// fmt.Println()

}
