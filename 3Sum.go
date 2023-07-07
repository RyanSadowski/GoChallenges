package main
 
import (
    "sort"
    "fmt"
)


func findSumOfThree(nums []int, target int) bool {
  sort.Ints(nums)
	for i := range nums {
        left := i + 1
        right := len(nums) - 1
        
        for left < right {
          
          result := nums[i] + nums[left] + nums[right]
          fmt.Println("i:", i, "left:", left, "right:", right, "result:", result)
        
          if (result == target){
            return true
          }

          if(result < target){
            left = left + 1
          }

          if(result >= target){
            right = right - 1 
          }
      
        }
    }                
  return false
}
