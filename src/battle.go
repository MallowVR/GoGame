package main

type Skill struct {
	Name    string
	Healing float32
	Hits    uint16
	Damage  float32
	Block   float32
}

var Skills []Skill

func LoadSkills() {
	ReadJsonFile(&Skills, "skills.json")
	WriteJsonFile(&Skills, "skills.json") // Save the file here to automatically add any fields that were missing from the config before
}
