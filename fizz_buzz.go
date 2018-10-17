package main

import (
	"fmt"
	"flag"
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

var number int

func init(){
	flag.IntVar(&number, "number", 1, "number of items")
}

func main() {
	flag.Parse()

	for _, str := range fizzBuzz(number) {
		fmt.Println(str)
	}
}
