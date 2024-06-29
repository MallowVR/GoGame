package main

import (
	"fmt"
	"math"
)

func commandHandler(p *player, args []string) {
	length := len(args)
	// for i := 0; i < length, i++
	if length == 0 {
		fmt.Println("Unexpected no arguments")
		return
	}
	switch args[0] {
	case "stats":
		fmt.Println("User is "+p.GetID()+" money:", currencyFormatter(p.Money))
	case "battle":
		p.Money += uint64(Conf.MoneyRate * float64(p.Weapon))
		fmt.Print("Battle complete, earned ", currencyFormatter(uint64(Conf.MoneyRate*float64(p.Weapon))), ". You now have ", currencyFormatter(p.Money), "\n")
	case "upgrade":
		if length < 1 {
			fmt.Println("Adventurer, what do you want to upgrade?")
			fmt.Println("Proper input is upgrade [weapon/armor]")
		} else {
			switch args[1] {
			case "weapon":
				var cost uint64 = uint64(math.Pow(2, float64(p.Weapon)))
				if p.Money < cost {
					fmt.Println("You cannot afford that adventurer, you need", currencyFormatter(cost), "for that upgrade")
				} else {
					p.Money -= cost
					p.Weapon += 1
					fmt.Println("Your weapon has been upgraded to level", p.Weapon, " You have", currencyFormatter(p.Money), "remaining")
				}
			}
		}
	}
}
