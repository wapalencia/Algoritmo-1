<<<<<<< HEAD
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("ana"))
	fmt.Println(isPalindrome("raro"))
}
=======
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("ana"))
	fmt.Println(isPalindrome("raro"))
}
>>>>>>> 8f08e6254e788eb25f885246452dc88613b445ff
