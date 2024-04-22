package main

import (
  "fmt"
  "reset2099.com/greetings"
  "log"
)

func main(){

  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  names:=[]string{"gladis","samara","doro"}

  messages,err:=greetings.Hellos(names)

  //message,err:=greetings.Hello("josj")
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println(messages)
}
