package main

import (
    "fmt"
    "math/rand"
    "strings"
)

func fishcerYates(s string) string {
    a := strings.Split(s, " ")
    for i := range a {
        if i == 0 {
            continue
        }
        j := rand.Intn(i)
        temp := a[i]
        a[i] = a[j]
        a[j] = temp
    }
    return strings.Join(a, " ")
}

func main() {
    test := "1 2 3 4 5 6 7 8"
    fmt.Println(test)
    shuffled := fishcerYates(test)
    fmt.Println(shuffled)
}