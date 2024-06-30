package main

import (
	"fmt"
	"time"
)

type Times struct {
	Name     string
	LastTime time.Time
}

type player struct {
	ID          string
	Name        string
	Money       uint64
	Crystals    uint64
	Level       uint16
	Armor       uint16
	Weapon      uint16
	Warding     uint16
	Skill       uint16
	DailyStreak uint64
	Times       []Times
}

type playerInterface interface {
	GetID()
	setID()
	Initialize()
	loadPlayer()
}

func (p *player) GetTime(_in string) time.Time {
	for i := 0; i < len(p.Times); i++ {
		if p.Times[i].Name == _in {
			return p.Times[i].LastTime
		}
	}
	p.Times = append(p.Times, Times{
		Name: _in,
	})
	return time.Time{}
}

func (p *player) SetTime(_in string, _time time.Time) {
	for i := 0; i < len(p.Times); i++ {
		if p.Times[i].Name == _in {
			p.Times[i].LastTime = _time
			return
		}
	}
	p.Times = append(p.Times, Times{
		Name:     _in,
		LastTime: _time,
	})
	return
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
