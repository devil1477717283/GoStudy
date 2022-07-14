package main

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	S:=NewSet()
	S.Put("Hello")
	S.Put("World")
	got1:=S.Keys()
	want1:=[]string{"Hello", "World"}
	if !reflect.DeepEqual(got1,want1) {
		t.Errorf("got1 %#v, want1 %#v", got1,want1)
	}
	got2:=S.Contains("Hello")
	want2:=true
	if got2!=want2 {
		t.Errorf("got2 %t, want2 %t", got2,want2)
	}
	got3:=S.Contains("你好")
	want3:=false
	if got3!=want3 {
		t.Errorf("got3 %t, want3 %t", got3,want3)
	}
	S.Remove("Hello")
	got4:=S.ReturnMap()
	want4:=map[string]bool{
		"World":true,
	}
	if !reflect.DeepEqual(got4,want4) {
		t.Errorf("got4 %#v, want4 %#v", got4,want4)
	}
	got5,got6:=S.PutIfAbsent("Hello")
	want5:=""
	want6:=true
	if got5!=want5 || got6!=want6{
		t.Errorf("got5 %#v want5 %#v, got6 %#v want6 %#v", got5,want5,got6,want6)
	}
	got7,got8:=S.PutIfAbsent("Hello")
	want7:="Hello"
	want8:=false
	if got7!=want7 || got8!=want8{
		t.Errorf("got7 %#v want7 %#v, got8 %#v want8 %#v", got7,want7,got8,want8)
	}
}