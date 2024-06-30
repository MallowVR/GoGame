package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"time"
)

type cooldown struct {
	Days    uint16
	Hours   uint16
	Minutes uint16
	Seconds uint16
}

type command struct {
	name     string
	cd       cooldown
	function func(*player, []string) string
}

var commands = []command{
	command{
		name:     "help",
		function: helpCommand,
	},
	command{
		name:     "stats",
		function: statsCommand,
	},
	command{
		name: "daily",
		cd: cooldown{
			Days:    1,
			Hours:   0,
			Minutes: 0,
			Seconds: 0,
		},
		function: dailyCommand,
	},
	command{
		name: "explore",
		cd: cooldown{
			Days:    0,
			Hours:   0,
			Minutes: 20,
			Seconds: 0,
		},
		function: exploreCommand,
	},
	command{
		name: "test",
		cd: cooldown{
			Days:    0,
			Hours:   0,
			Minutes: 0,
			Seconds: 30,
		},
		function: testCommand,
	},
}

func helpCommand(p *player, args []string) string {
	return "Commands:\n- stats\n- battle\n- upgrade\n- inventory\n- store\n- ability\n- re-roll\n- ascend"
}

func statsCommand(p *player, args []string) string {
	return fmt.Sprint(
		"Level: ", p.Level,
		"\nCoin purse: ", currencyFormatter(p.Money),
		"\nSkill: ", Skills[p.Skill].Name,
	)
}

func battleCommand(p *player, args []string) string {
	return ""
}

func dailyCommand(p *player, args []string) string {
	// currentTime := time.Now().UTC()
	// if p.Times.Daily.Year() == currentTime.Year() && p.Times.Daily.YearDay() == currentTime.YearDay() {
	// 	return "You have already claimed this today!\nPlease come back tomorrow!"
	// }
	// if p.Times.Daily.Year() != currentTime.Year() || p.Times.Daily.YearDay() != (currentTime.YearDay()-1) {
	// 	p.DailyStreak = 1
	// } else {
	// 	p.DailyStreak = p.DailyStreak + 1
	// }
	// p.Times.Daily = currentTime
	var money = uint64(Conf.MoneyRate * (100 * math.Log2(float64(p.DailyStreak+1))))
	p.Money = p.Money + money
	return fmt.Sprint("Streak: ", p.DailyStreak, "\n Claimed Daily, got ", currencyFormatter(money))
}

func exploreCommand(p *player, args []string) string {
	// currentTime := time.Now().UTC()
	// if currentTime.Sub(p.Times.Explore).Seconds() < 1200 {
	// 	return fmt.Sprint("You cannot explore yet, come back in ", (20 - currentTime.Sub(p.Times.Explore).Minutes()))
	// }
	// p.Times.Explore = currentTime
	var money = uint64(Conf.MoneyRate * (float64(rand.Int63n(100))))
	p.Money = p.Money + money
	return fmt.Sprint("Exploration complete, you got ", currencyFormatter(money))
}

func testCommand(p *player, args []string) string {
	return "feedback"
}

func (cmd *command) coolDownCheck(p *player) bool {

	if cmd.cd.Days == 0 && cmd.cd.Hours == 0 && cmd.cd.Minutes == 0 && cmd.cd.Seconds == 0 {
		return true
	}
	currentTime := time.Now().UTC()
	lastTime := p.GetTime(cmd.name)
	if cmd.cd.Hours == 0 && cmd.cd.Minutes == 0 && cmd.cd.Seconds == 0 && (currentTime.Year() != lastTime.Year() || (currentTime.YearDay()-lastTime.YearDay() >= int(cmd.cd.Days))) {
		p.SetTime(cmd.name, currentTime)
		return true
	}
	var cd = (((cmd.cd.Days*24)+cmd.cd.Hours)*60+cmd.cd.Minutes)*60 + cmd.cd.Seconds
	if currentTime.Unix()-lastTime.Unix() > int64(cd) {
		p.SetTime(cmd.name, currentTime)
		return true
	}
	return false
}

func commandHandler(p *player, args []string) string {
	length := len(args)
	var output string = ""
	if length == 0 {
		slog.Info("Unexpected no arguments")
		return output
	}
	for i := 0; i < len(commands); i++ {
		if commands[i].name == args[0] {
			if commands[i].coolDownCheck(p) {
				output = commands[i].function(p, args)
			} else {
				output = "This is on cooldown, please check back later"
			}
		}
	}
	if output == "" {
		output = helpCommand(p, args)
	}
	// switch args[0] {
	// case "stats":
	// 	output = fmt.Sprint("User is "+p.Name+" money: ", currencyFormatter(p.Money))
	// case "battle":
	// 	p.Money += uint64(Conf.MoneyRate * float64(p.Weapon))
	// 	output = fmt.Sprint("Battle complete, earned ", currencyFormatter(uint64(Conf.MoneyRate*float64(p.Weapon))), ". You now have ", currencyFormatter(p.Money), "\n")
	// case "upgrade":
	// 	if length < 2 {
	// 		output = fmt.Sprint("Adventurer, what do you want to upgrade? \nProper input is upgrade [weapon/armor]")
	// 	} else {
	// 		switch args[1] {
	// 		case "weapon":
	// 			var cost uint64 = uint64(math.Pow(2, float64(p.Weapon)))
	// 			if p.Money < cost {
	// 				output = fmt.Sprint("You cannot afford that adventurer, you need ", currencyFormatter(cost), " for that upgrade")
	// 			} else {
	// 				p.Money -= cost
	// 				p.Weapon += 1
	// 				output = fmt.Sprint("Your weapon has been upgraded to level ", p.Weapon, " You have ", currencyFormatter(p.Money), " remaining")
	// 			}
	// 		}
	// 	}
	// default:

	// }

	return output
}
