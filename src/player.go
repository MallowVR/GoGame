package main

import (
	"fmt"
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
	Name     string
	Money    uint64
	Crystals uint64
	Armor    uint16
	Weapon   uint16
	Warding  uint16
	Skill    uint16
	Times    times
}

type playerInterface interface {
	GetID()
	setID()
	Initialize()
	loadPlayer()
}

func (p *player) Initialize() {
	p.Money = 10
	p.Crystals = 0
	p.Skill = 0
	p.Weapon = 1
	p.Armor = 1
	p.Warding = 1
}

func (p *player) GetID() string {
	return p.ID
}

func (p *player) setID(_in string) {
	p.ID = _in
	return
}

func (p *player) loadPlayer(_userName string) {
	ReadJsonFile(&p, "players/"+_userName)

	if p.GetID() == "" {
		fmt.Println("found", p.ID, "initializing player", _userName)
		p.setID(_userName)
		p.Initialize()
	}
}

func (p *player) savePlayer() {
	WriteJsonFile(&p, "players/"+p.ID)
}
