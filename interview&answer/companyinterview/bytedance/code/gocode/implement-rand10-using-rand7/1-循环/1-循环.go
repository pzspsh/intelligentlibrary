package main

func main() {

}

// 用Rand7()实现 Rand10()
func rand7() int {
	for {
		a := rand7()
		b := rand7()
		target := a + (b-1)*7
		if target <= 40 {
			return target%10 + 1
		}
	}
}
