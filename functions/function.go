package functions

import (
   "fmt"
   "time"
   "reflect" //untuk mengetahui tipe data dari suatu var
   "strconv"

   Model "simple-go/model"
)

func Scanning(){
   var name string
   var nameLength int
   
   //Scan == Scanner (Java)
      //Scanning and reading the given input texts
   fmt.Scan(&name)
   fmt.Scan(&nameLength)
   
   // Print
   fmt.Printf("This name %s contains %d letters", name, nameLength)
}

func Pointer(){
   var initName = "test"
   var name *string = &initName //name merujuk pada alamat memory initName
   fmt.Println(name) //Output: 0xc000040230
   fmt.Println(*name) //Output: test

   //let's change the value of name
   *name = "update test"
   fmt.Println(name) //Output: 0xc000040230
   fmt.Println(*name) //Output: update test

   fmt.Println(&initName) //Output: 0xc000040230
   fmt.Println(initName) //Output: update test
}

func Lists(){
   var numbers = []int{1,2,3,4,5}
   fmt.Println(numbers) //Output: [1 2 3 4 5]
   fmt.Printf("number index 0: %d", numbers[0])
   fmt.Println()

   //Convert int to string
   str := strconv.Itoa(numbers[0])
   fmt.Println("number index 0 [Converted]: " + str)
}

func Loops(){
   for i := 0; i < 3; i++{
      fmt.Printf("Iteration-%d", i)
      fmt.Println()
   }
}

func LoopsAndRange(){
   var fruits = [4]string{"apple", "grape", "banana", "melon"}

   for i, fruit := range fruits {
      fmt.Printf("elemen %d %s\n", i, fruit)
   }
}

func Maps(){
   var chicken = map[string]int{"januari": 50, "februari": 40}
   var value, isExist = chicken["mei"] //50

   fmt.Println("Value: " + strconv.Itoa(value))
   fmt.Println("isExist: " + strconv.FormatBool(isExist))
   fmt.Println(reflect.TypeOf(strconv.FormatBool(isExist))) //tipe data kembalian FormatBool()

   fmt.Println(reflect.TypeOf(chicken["februari"])) //int
}

//https://dasarpemrogramangolang.novalagung.com/A-fungsi-sebagai-parameter.html
func FunctionAsParams(datas []string, callback func(string) bool) []string{
   var dataResults []string
   for i, data := range datas{
      fmt.Println("element-" + strconv.Itoa(i) + " with data: " + data)

      if exist := callback(data); exist{
         dataResults = append(dataResults, data)
      }
   }

   return dataResults
}


func PropertyWithinOrOutStruct(){
   var student = struct{
      pupil Model.Student
      age int
   }{} //{}==initialize with null value

   var student2 = struct{
      pupil Model.Student
      age int
   }{
      pupil: Model.Student{
         Nim: "M17929192",
         Email: "merry@gmail.com"},
      age: 17,
   }

   fmt.Println(student.pupil.ToString) //print memory address
   fmt.Println(student2.pupil.ToString()) //print value
}

//https://dasarpemrogramangolang.novalagung.com/A-interface-kosong.html
//Interface is a data type that can accommodate any type of data
func Interface() {
   //init, response bertipe list dengan key string dan value any type of data
   var response = []map[string]interface{}{
      {
         "status": 1,
         "message": "message",
         "timestamp": time.Now(),
      },
      {"status": 2, "message": "message", "timestamp": time.Now()},
   }

   for _, each := range response {
      fmt.Println(each["status"], each["message"], each["timestamp"])
   }
}