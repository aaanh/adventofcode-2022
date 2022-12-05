package problems

import (
  "fmt"
  "strings"
  // "io/ioutil"
)

func P5() {
  // Read input file
  // data, err := ioutil.ReadFile("inputs/p5.txt")
  // if err != nil {
  //   panic(err)
  // }

  // fileContents := string(data)

  arrOfInputs := []string{"FDBZTJRN",
    "RSNJH",
    "CRNJGZFQ",
    "FVNGRTQ",
    "LTQF",
    "QCWZBRGN",
    "FCLSNHM",
    "DNQMTJ",
    "PGS"}

  var table [][]string

  for i := 0; i < len(arrOfInputs); i++ {
    table = append(table, tmp)
    for j := 0; j < len(arrOfInputs[i]); j++ {
      fmt.Printf(strings.Split(arrOfInputs[i], "")[j])
      table[i] = append(table[i], strings.Split(arrOfInputs[i], "")[j])
    }
    fmt.Println()
  }

  fmt.Println("Problem 5a")
  // fmt.Println(strings.Split(arrOfInputs[len(arrOfInputs) -1], "")[0])
  // fmt.Println(len(arrOfInputs))
  // fmt.Println(len(arrOfInputs[len(arrOfInputs)-1]))
  fmt.Println(table)

}