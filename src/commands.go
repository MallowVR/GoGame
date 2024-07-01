package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"time"
)

type cooldown struct {
	Days     uint16
	Hours    uint16
	Minutes  uint16
	Seconds  uint16
	Response string
}

type command struct {
	name     string
	cd       cooldown
	function func(*player, []string) string
}

var commands = []command{
	{
		name:     "stats",
		function: statsCommand,
	},
	{
		name: "daily",
		cd: cooldown{
			Days:     1,
			Hours:    0,
			Minutes:  0,
			Seconds:  0,
			Response: "You have already collected you Daily, come back tomorrow",
		},
		function: dailyCommand,
	},
	{
		name: "explore",
		cd: cooldown{
			Days:    0,
			Hours:   3,
			Minutes: 0,
			Seconds: 0,
		},
		function: exploreCommand,
	},
	{
		name: "train",
		cd: cooldown{
			Days:    0,
			Hours:   0,
			Minutes: 30,
			Seconds: 0,
		},
		function: trainCommand,
	},
	{
		name: "test",
		cd: cooldown{
			Days:    0,
			Hours:   0,
			Minutes: 0,
			Seconds: 0,
		},
		function: testCommand,
	},
}

func helpCommand(p *player, args []string) string {
	// return "Commands:\n- stats\n- battle\n- upgrade\n- inventory\n- store\n- ability\n- re-roll\n- ascend"
	var output string = "Commands:\n- help\n"
	if commands == nil {
		return output
	}
	for i := 0; i < len(commands); i++ {
		output = fmt.Sprint(output, "- ", commands[i].name, "\n")
	}
	return output
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
	var money = uint64(Conf.MoneyRate * (100 * math.Log2(float64(p.DailyStreak+1))))
	p.Money = p.Money + money
	return fmt.Sprint("Streak: ", p.DailyStreak, "\n Claimed Daily, got ", currencyFormatter(money))
}

func exploreCommand(p *player, args []string) string {
	var money = uint64(Conf.MoneyRate * (float64(rand.Int63n(100))))
	p.Money = p.Money + money
	return fmt.Sprint("Exploration complete, you got ", currencyFormatter(money))
}

func trainCommand(p *player, args []string) string {
	var output string
	var temp uint64
	rng := rand.Intn(100)
	if rng < 5 {
		temp = uint64(Conf.XPRate) * 10
		output = fmt.Sprint("You dragged your feet. ", temp, " xp gained")
	} else if rng > 95 {
		temp = uint64(Conf.XPRate) * 100
		output = fmt.Sprint("You underwent exceptional training. ", temp, " xp gained")
	} else {
		temp = uint64(Conf.XPRate) * 30
		output = fmt.Sprint("Training complete. ", temp, " xp gained")
	}
	if p.Experience+temp > p.Experience {
		p.Experience = p.Experience + temp
	}
	if GetLevel(p.Experience) > p.Level {
		p.Level = GetLevel(p.Experience)
		output = fmt.Sprint(output, "\nYou leveled up! You are now level ", p.Level)
	}
	return output
}

var offset = 0

func testCommand(p *player, args []string) string {
	var output string
	// var temp uint64
	for i := 0; i < 100 && i < len(ExpTable); i = i + 1 {
		output = fmt.Sprint(output, i+offset+1, " ", ExpTable[i+offset], "\n")
	}
	offset = offset + 100
	return output
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

	return output
}
