package main

import (
	"fmt"
	"strings"
	
)

//create struct
type Animal struct {
	Food string
	Locomotion string
	Noise string
}

//funcs 
func (a *Animal) Eat() { fmt.Println(a.Food)  }
func (a *Animal) Move() { fmt.Println(a.Locomotion) }
func (a *Animal) Speaks() { fmt.Println(a.Noise) }

func selOption(a Animal,opt string){
	switch opt{
		case "eat": a.Eat()
		case "move": a.Move()
		case "speak": a.Speaks()
		default : fmt.Println("	Unknow: ",opt, "option")
	}
}

//main func
func main(){
	cow:=Animal{Food: "grass",Locomotion: "walk",Noise: "moo"}
	bird:=Animal{Food: "worms",Locomotion: "fly",Noise: "peep"}
	snake:=Animal{Food: "mice",Locomotion: "slither",Noise: "hsss"}

	//user prompt
	var animal, opt string

	fmt.Println("Insert: <animal> <option>")
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s",&animal,&opt)
		animal=strings.TrimSuffix(animal,"\n")
		opt=strings.TrimSuffix(opt,"\n")
		switch animal{
			case "cow":	selOption(cow,opt)
			case "bird": selOption(bird,opt)
			case "snake": selOption(snake,opt)	
			default: fmt.Println("	Unknow: ",animal, "animal") 
		}
		animal,opt="","" 	
	}
		
}