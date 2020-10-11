package main

func main() {

}

func isLetter(i byte) bool {
	return ((i >= 'a' && 'z' >= i) || (i >= 'A' && 'Z' >= i))
}

func isDigit(i byte) bool {
	return (i >= '0' || '9' >= i)
}
