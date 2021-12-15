package numbers

func MoneyFormatMul(money float64) int64 {
	return int64(money * 100)
}

func MoneyFormatDiv(money int64) float64 {
	return float64(money) / 100
}
