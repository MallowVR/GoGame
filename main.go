package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/user"
	"time"
)

type times struct {
	Daily  time.Time
	Train  time.Time
	Battle time.Time
	Boss   time.Time
}

type player struct {
	ID       string
	Money    uint64
	Crystals uint64
	Armor    uint16
	Weapon   uint16
	Warding  uint16
	Skill    uint16
	Times    times
}

type Config struct {
	XPRate                 float64
	MoneyRate              float64
	CrystalRate            float64
	PlayerDamageReduction  float64
	PlayerDamageMultiplier float64
	PlayerHPMultiplier     float64
	Skills                 uint16
}

type Skill struct {
	Name    string
	Healing float32
	Hits    uint16
	Damage  float32
	Block   float32
}

var Skills []Skill
var Conf Config

type playerInterface interface {
	GetID()
	setID()
	Initialize()
}

func (p *player) Initialize() {
	p.Money = 10
	p.Crystals = 0
	p.Skill = 0
	p.Weapon = 0
	p.Armor = 0
	p.Warding = 0
}

func (p *player) GetID() string {
	return p.ID
}

func (p *player) setID(_in string) {
	p.ID = _in
	return
}

func ReadJsonFile(_in any, _fileName string) bool {
	file, err := os.OpenFile(_fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
		return false
	}

	fileContent, err := os.ReadFile(_fileName)
	if err != nil {
		panic(err)
		return false
	}

	file.Close()

	json.Unmarshal(fileContent, _in)
	return true
}

func WriteJsonFile(_in any, _fileName string) {
	file, err := os.OpenFile(_fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(_in, "", "  ")
	if err != nil {
		panic(err)
	}

	if file != nil {
		file.Write(json)
	}

	file.Sync()

	file.Close()
}

func LoadConfig() {
	ReadJsonFile(&Conf, "config.json")
	WriteJsonFile(&Conf, "config.json")
}

func LoadSkills() {
	ReadJsonFile(&Skills, "skills.json")
}

func currencyFormatter(_in uint64) string {
	var temp = _in
	var coppers = temp % 20
	temp = temp / 20
	var silvers = temp % 15
	temp = temp / 15
	var golds = temp
	var out string = ""
	if golds != 0 {
		out = fmt.Sprint(out, golds, " Gold ")
	}
	if silvers != 0 {
		out = fmt.Sprint(out, silvers, " silvers ")
	}
	if coppers != 0 {
		out = fmt.Sprint(out, coppers, " coppers")
	}
	return out
}

func main() {
	LoadConfig()
	// Skills = make([]Skill, Conf.Skills)
	LoadSkills()
	println(len(Skills))
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	var a player

	test := ReadJsonFile(&a, "players/"+user.Name)
	if !test {
		return
	}

	if a.GetID() == "" {
		fmt.Println("found", a.ID, "initializing player", user.Name)
		a.setID(user.Name)
	}

	var command string = "stats"

	if len(os.Args) > 0 {
		command = os.Args[1]
	}

	switch command {
	case "stats":
		fmt.Println("User is "+a.GetID()+" money:", currencyFormatter(a.Money))
		break
	case "battle":
		a.Money += uint64(Conf.MoneyRate * float64(a.Weapon))
		fmt.Print("Battle complete, earned ", currencyFormatter(uint64(Conf.MoneyRate*float64(a.Weapon))), ". You now have ", currencyFormatter(a.Money), "\n")
		break
	case "upgrade":
		if len(os.Args) < 3 {
			fmt.Println("Adventurer, what do you want to upgrade?")
			fmt.Println("Proper input is upgrade [weapon/armor]")
		} else {
			switch os.Args[2] {
			case "weapon":
				var cost uint64 = uint64(math.Pow(2, float64(a.Weapon)))
				if a.Money < cost {
					fmt.Println("You cannot afford that adventurer, you need", currencyFormatter(cost), "for that upgrade")
				} else {
					a.Money -= cost
					a.Weapon += 1
					fmt.Println("Your weapon has been upgraded to level", a.Weapon, " You have", currencyFormatter(a.Money), "remaining")
				}
			}
		}
	}

	WriteJsonFile(&a, "players/"+user.Name)

}
