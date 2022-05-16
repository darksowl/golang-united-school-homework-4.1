package reverse_string
//package main
import ("fmt")
func ReverseString(input string) (output string) {
	// solution goes here
	s := []rune(input)
	r := []rune(input)
	for i:=0; i<len(r);i++ {
		r[len(r)-i-1] = s[i]
	}
	return string(r)
}

//func main() {
//	s := "Jetbrains系"
//	r:= ReverseString(s)
//	fmt.Println(r)
//}