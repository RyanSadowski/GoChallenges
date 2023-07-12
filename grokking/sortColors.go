package main


import(
"fmt"
)
func sortColors(colors []int) []int {

	// Write your code here

	return selectionSort(colors)
}


func selectionSort(input []int) []int {
    var size int = len(input)
    for i := 0; i < size -1; i ++ {
    
        for j := i + 1; j < size; j ++ {
            if input[j] < input[i]{
                //swap
                temp := input[i]
                input[i] = input[j] 
                input[j] = temp
            }
        }
    }
    return input 
}

func bubbleSort(input []int) []int{
    var size = len(input)
    for i := 0; i < size -1; i ++ {
        for j := 0; j < size -1 - i; j ++ {
               //swap

               if(input[j] > input[j+1]) {
                   input[j], input[j+1] = input[j+1], input[j]
               }
        }
    }

    return input
    
}

func main() {
    var input = [] int{0,1,2,2,1,2,2,0,0,1,2,1,0,2,0,1}

    var output = bubbleSort(input)

    fmt.Println(output)
}
