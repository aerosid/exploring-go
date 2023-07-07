package db

import (
	"fmt"
	"testing"
)

func TestDbPing(t *testing.T) {
	db1, err1 := GetInstance()
	if err1 == nil {
		p1 := fmt.Sprintf("%p", db1)
		db2, err2 := GetInstance()
		if err2 == nil {
			p2 := fmt.Sprintf("%p", db2)
			if p1 == p2 {
				t.Log("success: p1 | p2 =", p1, "|", p2)
			} else {
				t.Error("success: p1 | p2 =", p1, "|", p2)
			}
		} else {
			t.Error("failure:", err2.Error())
		}
	} else {
		t.Error("failure:", err1.Error())
	}
}
