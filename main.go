package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)


type GenParams struct {
    Characters int
    Charset []rune
}


func main() {
   letters := []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123467890!@#$%^&*()-=[]{}\|"':;?/>.<,~`)
    var pass string
    lenght, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("could not")
    }
    for i := 0; i < lenght; i++ {
        currentChar := rand.Intn(len(letters) - 1)
        pass = fmt.Sprintf("%v%v", pass, string(letters[currentChar]))
    }
    fmt.Println(pass)
}
