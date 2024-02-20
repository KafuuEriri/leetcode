package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var result uint32

	// 逐位反转
	for i := 0; i < 32; i++ {
		// 将 result 左移一位
		result <<= 1
		// 如果 n 的最低位是 1，则将 result 的最低位置为 1
		if num&1 == 1 {
			result |= 1
		}
		// 将 n 右移一位
		num >>= 1
	}
	return result
}

func main() {
	// 测试用例
	num := uint32(43261596)
	reversed := reverseBits(num)
	fmt.Printf("反转后的结果为：%d\n", reversed)
}
