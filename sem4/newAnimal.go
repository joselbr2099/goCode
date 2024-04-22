package main

import(
	"fmt"
)

type Animal interface{
	Eat()
	Move()
	Speaks()
}

//create struct
type Cow struct {
	Food string
	Locomotion string
	Noise string
}
//funcs 
func (c Cow) Eat() { fmt.Println(c.Food)  }
func (c Cow) Move() { fmt.Println(c.Locomotion) }
func (c Cow) Speaks() { fmt.Println(c.Noise) }


//create struct
type Snake struct {
	Food string
	Locomotion string
	Noise string
}
//funcs 
func (s Snake) Eat() { fmt.Println(s.Food)  }
func (s Snake) Move() { fmt.Println(s.Locomotion) }
func (s Snake) Speaks() { fmt.Println(s.Noise) }


//create struct
type Bird struct {
	Food string
	Locomotion string
	Noise string
}
//funcs 
func (b Bird) Eat() { fmt.Println(b.Food)  }
func (b Bird) Move() { fmt.Println(b.Locomotion) }
func (b Bird) Speaks() { fmt.Println(b.Noise) }

func main(){
	//create interface
	var a Animal

	//define vars
	var command,name,option string 
	var animalMap map[string]struct{}
	
	//create types
	var cow = Cow{Food: "grass",Locomotion: "walk",Noise: "moo"}
	bird:=Bird{Food: "worms",Locomotion: "fly",Noise: "peep"}
	snake:=Snake{Food: "mice",Locomotion: "slither",Noise: "hsss"}

	getHelp()
	for{
		fmt.Print("> ")
		fmt.Scan("%s %s %s",&command,&name,&option)
		if ( len(command)!=0 && len(name)!=0 && len(option)!=0 ){
			switch command{
				case "newanimal":
					if getAnimal(strings.ToLower(animal),animalMap){
						animalMap[name]=strings.ToLower(option)
						fmt.Println(" Created it!")
					}else{
						fmt.Prinln(" error animal ",option," not exist")
					}														
				case "query":
					if animal,ok:=animalMap[name]; ok{
						getOption(a,animal,option)
					}
				
				default : fmt.Println(" Command: %s",command, "not found")			
			}	
		}else{
			getHelp()
		}	
	}
}

func getHelp(){
	//user input
	fmt.Println("type: newanimal <name> <cow|snake|bird> for create a new animal")
	fmt.Println("type: query <name> <eat|move|speaks>	fer query  information	")
}

func getAnimal(animal string)bool{
		switch animal{
			case "cow": return true
			case "bird": return true
			case "snake": return true
			default : return false
		}
}

func getOption(a Animal, animal type, option string){
	a=animal
	switch option{
		case "ceat":a.Eat() 
		case "move":a.Move()
		case "speaks":a.Speaks()
		default : fmt.Prinln(" rare animal")
	}	
}