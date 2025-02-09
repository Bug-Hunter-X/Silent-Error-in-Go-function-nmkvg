func myFunc(a, b int) (int, error) {
  if a == 0 {
    return 0, errors.New("a cannot be zero")
  }
  return a + b, nil
}

func main() {
  result, err := myFunc(0, 5)
  if err != nil {
    fmt.Println("Error:", err)
    return //or handle the error appropriately
  } 
  fmt.Println("Result:", result)
}
