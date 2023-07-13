package main

import(
    "fmt"
)

func sortColors(colors []int) []int {

    return threePointer(colors)
}



func threePointer(input []int) []int {
    var red int = 0
    var white int = 0
    var blue = len(input) -1

    for white <= blue {

        if input[white] == 0 {
            if input[red] != 0 {
                input[red], input[white] = input[white], input[red]
            }

            white ++
            red ++
        } else if input[white] == 1 {
            white ++
        } else {
            if input[blue] != 2 {
                input[white], input[blue] = input[blue], input[white]
            }

            blue--
        }
    }

    return input
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
