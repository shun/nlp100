package main

import "fmt"

func main() {
    //s := "あいう"
    s := "stressed"
    rs := ([]rune(s))

    for i, j := 0, len(rs) - 1; i < j; i, j = i + 1, j - 1 {
        rs[i], rs[j] = rs[j], rs[i]
    }

    fmt.Print(string(rs))
}
