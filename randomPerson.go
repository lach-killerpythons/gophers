package main

import(
  "fmt"
  "math/rand"
  "time"
  "os"
  "log"
  "bufio"
)
//name1 int, name2 int
func namePick(name1 int, name2 int) (firstName string, secondName string) {

    f, err := os.Open("names.txt")
    counter :=1
    first := true
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        if counter == name1 || counter == name2{
          if first{
            first = false
            firstName = scanner.Text()
          } else{
            secondName = scanner.Text()
          }
        }
        //fmt.Println(counter)
        counter++
        //fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return firstName, secondName
}

func swap(x int, y int) (out1 int, out2 int){

  if x < y{
    out1 = x
    out2 = y
  } else if y < x{
    out1 = y
    out2 = x
  } else if x == y{
    out1 = x
    out2 = x+1
  } else{
    out1 = 1
    out2 = 2
  }
  return out1, out2
}

func main() {
  rand.Seed(time.Now().UnixNano())
  name1 := rand.Intn(18238)
  name2 := rand.Intn(18238)
  //namePick()
  fmt.Println(namePick(name1,name2))


}
