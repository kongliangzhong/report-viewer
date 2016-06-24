package main

import (
    "testing"
)

func Test_NewBaseBarOption(t *testing.T) {
    bbo := NewBaseBarOption()
    t.Log(bbo)
    t.Log("pass")
}
