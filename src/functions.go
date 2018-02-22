package spec

func ReturnMultiple() (string, string) {
  return "foo", "bar"
}

func CreateIncrementerClosure() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}