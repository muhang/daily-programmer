package main

import "fmt"

func getGarland(s string) int {
    ret := 0
    for i := range s {
        if s[:i] == s[len(s)-i:] {
            ret = i
        }
    }
    return ret
}

func printGarland(s string, r int) string {
    garlandIdx := getGarland(s)
    ret := ""
    for i := 0; i < r-1; i++ {
        ret += s[:garlandIdx+1]
    }
    ret += s
    if (len(ret) > 0) {
        return ret
    } else {
        return "I can't make a garland out of that word :("
    }
}

func main() {
   str := printGarland("onion", 10)
   fmt.Println(str)
}