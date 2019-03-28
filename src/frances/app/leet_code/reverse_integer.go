/*Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321
*/

package main

func reverse(x int) int {
	sign := 1
	//负数，转正， 并标识
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x = x / 10
	}

	res = res * sign

	return res
}

func main() {
	println(reverse(-41213))
}
