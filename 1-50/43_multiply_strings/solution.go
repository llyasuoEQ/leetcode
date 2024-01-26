package multiply_strings

// multiply 字符串相乘
// 转换为手动计算俩数相乘的逻辑，字符 - 字符0 就可以用来计算
func multiply(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)
	// 创建保存每一位的数组，俩个num1Len和num2Len的数相乘最大位数就是它俩的和
	resultList := make([]int, num1Len+num2Len)
	for i := num1Len - 1; i >= 0; i-- {
		for j := num2Len - 1; j >= 0; j-- {
			// i+j就是此时每一位相乘在result中的位置
			product := int((num1[i] - '0') * (num2[j] - '0'))
			index1, index2 := i+j, i+j+1
			sum := product + resultList[index2] // 当前位 + 乘积
			resultList[index1] += sum / 10      // 当前位值
			resultList[index2] = sum % 10       // 进位值
		}
	}

	// 转换字符串
	var result string
	for _, v := range resultList {
		if !(len(result) == 0 && v == 0) {
			result = result + string('0'+rune(v))
		}
	}
	if len(result) == 0 {
		result = "0"
	}

	return result
}
