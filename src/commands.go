package main

import (
	"fmt"
	"log/slog"
	"math"
)

func commandHandler(p *player, args []string) string {
	length := len(args)
	// for i := 0; i < length, i++
	var output string
	if length == 0 {
		slog.Info("Unexpected no arguments")
	}
	switch args[0] {
	case "stats":
		output = fmt.Sprint("User is "+p.Name+" money:", currencyFormatter(p.Money))
	case "battle":
		p.Money += uint64(Conf.MoneyRate * float64(p.Weapon))
		output = fmt.Sprint("Battle complete, earned ", currencyFormatter(uint64(Conf.MoneyRate*float64(p.Weapon))), ". You now have ", currencyFormatter(p.Money), "\n")
	case "upgrade":
		if length < 1 {
			output = fmt.Sprint("Adventurer, what do you want to upgrade? \nProper input is upgrade [weapon/armor]")
		} else {
			switch args[1] {
			case "weapon":
				var cost uint64 = uint64(math.Pow(2, float64(p.Weapon)))
				if p.Money < cost {
					output = fmt.Sprint("You cannot afford that adventurer, you need", currencyFormatter(cost), "for that upgrade")
				} else {
					p.Money -= cost
					p.Weapon += 1
					output = fmt.Sprint("Your weapon has been upgraded to level", p.Weapon, " You have", currencyFormatter(p.Money), "remaining")
				}
			}
		}
	}

	return output
}
