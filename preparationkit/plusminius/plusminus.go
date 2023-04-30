package main

import (
	"fmt"
)


func main() {
	var arr [8]int32 = [8]int32{1,2,3,-1,-2,-3,0,0}
    fmt.Println(arr)
    var result_positive, result_negative, result_zero float64
    for i:= 0; i < len(arr); i++ {
        switch {
        case arr[i] > 0:
			//fmt.Println("Positive")
            result_positive = result_positive + 1.0/float64(len(arr))
        case arr[i] < 0:
            result_negative += 1.0/float64(len(arr))
        default:
            result_zero += 1.0/float64(len(arr))
        }
    }
    fmt.Printf("%.6f\n",result_positive)
    fmt.Printf("%.6f\n",result_negative)
    fmt.Printf("%.6f\n",result_zero)
}
