package main

import (
	"fmt"
)

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
	if coppers != 0 || (golds == 0 && silvers == 0) {
		out = fmt.Sprint(out, coppers, " coppers")
	}
	return out
}
