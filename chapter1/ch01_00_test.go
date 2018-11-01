package main

import (
    "fmt"
    "testing"
)

func TestReverse(t *testing.T) {
    target := "stressed"
    result := "desserts"


    ret := reverse(target)
    if result != ret {
        t.Fatal(fmt.Sprintf("faild test, %s, %s", result, ret))
    }

}
