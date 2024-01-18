package main

func isPowerOfFour(n int) bool {
	if n == 1 || n == 4 {
		return true
	}
	if n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

func main() {

}
