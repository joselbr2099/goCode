package greetings

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
)

// If the funcion name starts with a capital letter can be called 
// by a funcion not in the same package
// Hello is called exported name
func Hello(name string)(string,error){

  if name==""{
    return "",errors.New("Empty name")
  }

  //message:=fmt.Sprintf(randomFormat(),name)

  message:=fmt.Sprintf(randomFormat())
  return message,nil
}

// It runs at calling this module
func init(){
  rand.Seed(time.Now().UnixNano())
}

// Not exported function
func randomFormat()string{
  formats:=[]string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
  }

  return formats[rand.Intn(len(formats))]
}

// New function with exported name Hellos
func Hellos(names []string) (map[string]string,error) {
 
  // Create a map of strings
  messages :=make(map[string]string)

  // Iterate over names slice
  for _,name:=range names {

    // Save generated greetings in message var
    message,err:= Hello(name)
    if err!=nil{
      return nil,err
    }

    // save in slice
    messages[name]=message
  }

  return messages, nil
}
