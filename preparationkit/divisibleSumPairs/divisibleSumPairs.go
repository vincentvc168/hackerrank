package main

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here
    var sum int32 = 0
    for i1, v1 := range ar {
        for i2, v2 := range ar {
            if i2 >= i1 {
                continue
            } else if (v1 + v2) % k == 0 {
                sum++
            }
        }
    }
    return sum
}

func main() {
	divisibleSumPairs(5, 3, []int32{1,2,3,4,5})
}