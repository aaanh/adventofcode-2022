package problems

import (
  "fmt"
  "strings"
  "io/ioutil"
  "strconv"
)

func parseOperation(s string) []int {
  arrSplitted := strings.Split(s, " ")
  target, _ := strconv.Atoi(arrSplitted[1])
  origin, _ := strconv.Atoi(arrSplitted[3])
  destin, _ := strconv.Atoi(arrSplitted[5])

  return []int{target, origin, destin}
}

func P5() {
  // Read input file
  data, err := ioutil.ReadFile("inputs/p5.txt")
  if err != nil {
    panic(err)
  }

  fileContents := string(data)

  arrOfInputs := []string{
    "FDBZTJRN",
    "RSNJH",
    "CRNJGZFQ",
    "FVNGRTQ",
    "LTQF",
    "QCWZBRGN",
    "FCLSNHM",
    "DNQMTJ",
    "PGS"}
  

  fmt.Println("Problem 5a")

}
