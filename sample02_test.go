package sample01

import (
  "testing"
)

func TestHello(t *testing.T) {
  actual := Hello("hoge")
  expected := "hello, hoge"
  if actual != expected {
    t.Errorf("actual %v\nwant %v", actual, expected)
  }
}