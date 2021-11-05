package main

import "fmt"
import "reflect"
import "runtime"

func main()  {
  runtime.GOMAXPROCS(4)

  go bacatipe1()
  bacatipe2()
  var input string
  fmt.Scanln(&input)
}

func bacatipe1()  {
  var a = 10
  var b = "Budi"
  var reflectValue1 = reflect.ValueOf(a)
  var reflectValue2 = reflect.ValueOf(b)
  fmt.Println("tipe  variabel 1A:", reflectValue1.Type())
  fmt.Println("tipe  variabel 2A:", reflectValue2.Type())
}


func bacatipe2()  {
  var a = 20
  var b = "Amir"
  var reflectValue1 = reflect.ValueOf(a)
  var reflectValue2 = reflect.ValueOf(b)
  fmt.Println("tipe  variabel 1B:", reflectValue1.Type())
  fmt.Println("tipe  variabel 2B:", reflectValue2.Type())
}
