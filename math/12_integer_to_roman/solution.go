package integer_to_roman

// intToRoman 罗列各种情况解答
// 1. 从左往右一次拼接罗马字符
func intToRoman(num int) string {
	var (
		res     = ""
		divisor = 1000
	)
	for divisor > 0 {
		temp := num / divisor
		if temp > 0 {
			switch divisor {
			case 1000:
				for i := 0; i < temp; i++ {
					res += "M"
				}
			case 100:
				if temp < 4 {
					for i := 1; i <= temp; i++ {
						res += "C"
					}
				} else if temp == 4 {
					res += "CD"
				} else if temp == 5 {
					res += "D"
				} else if temp > 5 && temp < 9 {
					res += "DC"
					for i := 7; i <= temp; i++ {
						res += "C"
					}
				} else if temp == 9 {
					res += "CM"
				}
			case 10:
				if temp < 4 {
					for i := 1; i <= temp; i++ {
						res += "X"
					}
				} else if temp == 4 {
					res += "XL"
				} else if temp == 5 {
					res += "L"
				} else if temp > 5 && temp < 9 {
					res += "LX"
					for i := 7; i <= temp; i++ {
						res += "X"
					}
				} else if temp == 9 {
					res += "XC"
				}
			case 1:
				if temp < 4 {
					for i := 1; i <= temp; i++ {
						res += "I"
					}
				} else if temp == 4 {
					res += "IV"
				} else if temp == 5 {
					res += "V"
				} else if temp > 5 && temp < 9 {
					res += "VI"
					for i := 7; i <= temp; i++ {
						res += "I"
					}
				} else if temp == 9 {
					res += "IX"
				}
			}
		}
		num = num % divisor
		divisor /= 10
	}
	return res
}

type RomaItem struct {
	Value int
	Roma  string
}

var romaList = []RomaItem{
	{Value: 1000, Roma: "M"},
	{Value: 900, Roma: "CM"},
	{Value: 500, Roma: "D"},
	{Value: 400, Roma: "CD"},
	{Value: 100, Roma: "C"},
	{Value: 90, Roma: "XC"},
	{Value: 50, Roma: "L"},
	{Value: 40, Roma: "XL"},
	{Value: 10, Roma: "X"},
	{Value: 9, Roma: "IX"},
	{Value: 5, Roma: "V"},
	{Value: 4, Roma: "IV"},
	{Value: 1, Roma: "I"},
}

// intToRoman2 依次递减方法
func intToRoman2(num int) string {
	var roma string
	for _, v := range romaList {
		for num >= v.Value && num > 0 {
			num -= v.Value
			roma += v.Roma
		}
	}
	return roma
}
