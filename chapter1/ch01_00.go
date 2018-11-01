package main

import "fmt"

func reverse(s string) string {
    rs := ([]rune(s))

    for i, j := 0, len(rs) - 1; i < j; i, j = i + 1, j - 1 {
        rs[i], rs[j] = rs[j], rs[i]
    }

    return string(rs)

}

func main() {
    //s := "あいう"
    s := "stressed"

    fmt.Print(reverse(s))
}
