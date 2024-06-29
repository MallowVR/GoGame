package main

type Config struct {
	XPRate                 float64
	MoneyRate              float64
	CrystalRate            float64
	PlayerDamageReduction  float64
	PlayerDamageMultiplier float64
	PlayerHPMultiplier     float64
	Skills                 uint16
}

var Conf Config

func LoadConfig() {
	ReadJsonFile(&Conf, "config.json")
	WriteJsonFile(&Conf, "config.json")
}
