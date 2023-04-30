package main

import "fmt"

func miniMaxSum(arr []int32) {
    // Write your code here
    var sum, min, max int64 = 0,0,0
    min = int64(arr[0])
    max = int64(arr[0])
    for i:= 0; i < len(arr); i++ {
        sum += int64(arr[i])
        if (int64(arr[i]) < min) {
            min = int64(arr[i])
        }
        if (int64(arr[i]) > max) {
            max = int64(arr[i])
        }
    }
    // Expected Output 2063136757 2744467344
    //fmt.Printf("%d \n", sum)
    fmt.Printf("%d ", sum-max)
    fmt.Printf("%d", sum-min)
}

func main() {
    var arr []int32 = []int32{256741038, 623958417, 467905213, 714532089, 938071625}
    // slice := []int32{1,2,3,4,5}
    // arr = append(arr, slice...)
    miniMaxSum(arr)
}