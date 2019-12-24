package main

import (
	"fmt"
)

func byte2int(b byte) uint8 {
	return uint8(b) - uint8('0')
}

func int2byte(n uint8) byte {
	return byte(n + uint8('0'))
}

func str2int(b []byte) []uint8 {
	ans := []uint8{}
	for i := 0; i < len(b); i++ {
		ans = append(ans, byte2int(b[i]))
	}
	return ans
}

func int2str(num []uint8) []byte {
	ans := []byte{}
	for i := 0; i < len(num); i++ {
		ans = append(ans, int2byte(num[i]))
	}
	return ans
}

func reverse(str []byte) []byte {
	ln := len(str)
	for i := 0; i < ln/2; i++ {
		str[i], str[ln-i-1] = str[ln-i-1], str[i]
	}
	return str
}

func trim(str []byte) []byte {
	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			return str[i:]
		}
	}
	return []byte{'0'}
}

func multiply(str1 string, str2 string) string {
	ans := []uint8{}
	nums1 := str2int([]byte(str1))
	nums2 := str2int([]byte(str2))
	var k int
	k = 0
	for i := len(nums1) - 1; i >= 0; i-- {
		var kt int
		kt = k
		var temp uint8
		temp = 0
		for j := len(nums2) - 1; j >= 0; j-- {
			var tt uint8
			tt = nums1[i]*nums2[j] + temp
			var ttt uint8
			ttt = tt % 10
			temp = tt / 10
			if len(ans) <= kt {
				ans = append(ans, ttt)
			} else {
				ans[kt] += ttt
				temp += ans[kt] / 10
				ans[kt] %= 10
			}
			kt++
		}
		if temp != 0 {
			if len(ans) <= kt {
				ans = append(ans, temp)
			} else {
				ans[kt] += temp
				temp = ans[kt] / 10
				ans[kt] %= 10
				kt++
				if temp != 0 {
					if len(ans) <= kt {
						ans = append(ans, temp)
					} else {
						ans[kt] += temp
					}
				}
			}
		}
		k++
	}
	return string(trim(reverse(int2str(ans))))
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
	num1 = "0"
	num2 = "0"
	fmt.Println(multiply(num1, num2))
	num1 = "123456"
	num2 = "0"
	fmt.Println(multiply(num1, num2))
	num1 = "9"
	num2 = "9"
	fmt.Println(multiply(num1, num2))
}
