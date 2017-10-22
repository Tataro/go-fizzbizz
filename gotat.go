package gotat

import "strconv"

func fizzbizz(x int) string {
  result := ""
  if (x % 3 == 0) {
    result += "fizz"
  }
  if (x % 5 == 0) {
    result += "bizz"
  }
  if (result == "") {
    result = strconv.Itoa(x)
  }
  return result
}
