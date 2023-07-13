package main

func isPalindrome(inputString string) bool{
	left := 0
    right := len(inputString) - 1
    for left <= right {
        if(inputString[left] != inputString[right]){
			return false
		}
		left = left + 1
        right = right -1 
		
    }
    return true
}
