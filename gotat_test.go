package gotat

import "testing"

// func TestFizzbizz(t *testing.T) {
//   mapping := map[int]string{
//     1: "1",
//     2: "2",
//     3: "fizz",
//     4: "4",
//     5: "bizz",
//     6: "fizz",
//     7: "7",
//     8: "8",
//     9: "fizz",
//     10: "bizz",
//     15: "fizzbizz",
//     30: "fizzbizz",
//   }
//   for k, v := range mapping {
//     result := fizzbizz(k)
//     if result != v {
//       t.Error("it wrong1 result=", result, "expected=", v)
//     }
//   }
// }

func TestFizzbizz1(t *testing.T) {
  result := fizzbizz(1)
  if result != "1" {
    t.Error("it wrong1 result=", result, "expected=1")
  }
}

func TestFizzbizz2(t *testing.T) {
  result := fizzbizz(2)
  if result != "2" {
    t.Error("it wrong1 result=", result, "expected=2")
  }
}

func TestFizzbizz3(t *testing.T) {
  result := fizzbizz(3)
  if result != "fizz" {
    t.Error("it wrong1 result=", result, "expected=fizz")
  }
}

func TestFizzbizz4(t *testing.T) {
  result := fizzbizz(4)
  if result != "4" {
    t.Error("it wrong1 result=", result, "expected=4")
  }
}

func TestFizzbizz5(t *testing.T) {
  result := fizzbizz(5)
  if result != "bizz" {
    t.Error("it wrong1 result=", result, "expected=5")
  }
}

func TestFizzbizz6(t *testing.T) {
  result := fizzbizz(6)
  if result != "fizz" {
    t.Error("it wrong1 result=", result, "expected=6")
  }
}

func TestFizzbizz7(t *testing.T) {
  result := fizzbizz(7)
  if result != "7" {
    t.Error("it wrong1 result=", result, "expected=7")
  }
}

func TestFizzbizz8(t *testing.T) {
  result := fizzbizz(8)
  if result != "8" {
    t.Error("it wrong1 result=", result, "expected=8")
  }
}

func TestFizzbizz9(t *testing.T) {
  result := fizzbizz(9)
  if result != "fizz" {
    t.Error("it wrong1 result=", result, "expected=9")
  }
}

func TestFizzbizz10(t *testing.T) {
  result := fizzbizz(10)
  if result != "bizz" {
    t.Error("it wrong1 result=", result, "expected=bizz")
  }
}

func TestFizzbizz15(t *testing.T) {
  result := fizzbizz(15)
  if result != "fizzbizz" {
    t.Error("it wrong1 result=", result, "expected=fizzbizz")
  }
}

func TestFizzbizz30(t *testing.T) {
  result := fizzbizz(30)
  if result != "fizzbizz" {
    t.Error("it wrong1 result=", result, "expected=fizzbizz")
  }
}
