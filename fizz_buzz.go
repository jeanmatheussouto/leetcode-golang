package main

import (
	"fmt"
)

func fizzBuzz(n int) []string {
    var r []string
    var str string

    for i := 1; i <= n; i ++ {
        if (i % 3 == 0 && i % 5 == 0) {
            str = "FizzBuzz"
        } else if(i % 3 == 0) {
            str = "Fizz"
        } else if(i % 5 == 0) {
            str = "Buzz"
        } else {
            str = fmt.Sprintf("%d", i)
        }
        r = append(r, str)
    }
    return r
}

func main() {
	result := fizzBuzz(15)
	fmt.Println(result)
}
