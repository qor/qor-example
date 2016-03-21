package utils

import (
	"fmt"
	"strconv"
)

// Возвращает номер заполняя в начале нулями определенной длины
// @return true/false, format 1234567 -> 0001234567
func GetFullNumber(card string, limit int) (res string) {
	res = ""
	cl := len(card)
	if cl >= limit {
		res = card
	} else if cl < limit {
		lenCard := limit - cl
		for i := 0; i < lenCard; i++ {
			res = fmt.Sprintf("%s0", res)
		}
		res += card
	}
	return res
}

// Возвращает номер без нулей в переди
// @return Number format 0001234567 -> 1234567
func GetNumber(card string) (res string) {
	v, _ := strconv.ParseInt(card, 10, 32)
	return fmt.Sprintf("%d", v)
}
