package main
import (
	"errors"
	"fmt"
)

type student struct {
	ID int
	The_name string
	fullname string
	age int
	Username string
	grades []int
	comment string
}
func main(){
	first := student {
		ID: 222,
		The_name: "ahmed",
		fullname: "ahmed salem",
		age: 17,
		Username: "ahmed@123",
		grades: []int {70,80,50,50},
		comment: "good not bad",
	}
	massage , problem := check(first.age)
	if problem != nil {
		fmt.Printf("the age [%v]\n",problem)
	}

    fmt.Println(check2(first.ID))
    fmt.Printf("the firstname [%v]\n",first.The_name)
    fmt.Printf("the fullname [%v]\n",first.fullname)
	fmt.Printf("the Username [%v]\n",first.Username)
    fmt.Printf("the total of the subjects %d\n",first.grades)
    fmt.Printf("the qailty [%v]\n",first.comment)
	
    fmt.Print(massage)
	
	password := "Hamo@123"
	gmail := "Hamosalemm123@gmail.com"
	test := true 
	test2 := true 
	for _, pass := range password {
		if !(pass >= 'a' && pass <= 'z')&&
		   !(pass >= 'A' && pass <= 'Z')&&
		   !(pass >= '0' && pass <= '9')&&
		   pass != '@'{
			test = false 
			break
		}
	}
	fmt.Printf("here's your password [%v]\n",password)
	if test {
		fmt.Println("welcome")
	}else {
		fmt.Println("try again")
	}
	// the gmail 
	for _, gmail1 := range gmail{
		c := gmail1
		if !(c >= 'a' && c <= 'z')&&
		   !(c >= 'A' && c <= 'Z')&&
		   !(c >= '0' && c <= '9')&&
		   (c != '@' && c != '.'){
			test2 = false 
			break
		} 
	}
	fmt.Printf("and also the gmail [%v]\n",gmail)
	if test2{
		fmt.Println("valid")
	}else {
		fmt.Println("passbley missing comma or something")
	}
}

func check(s int)(string,error){
	if s < 18 {
		return "",errors.New("we accept at least 18")
	}
	return "valid",nil
}

func check2(s int)string{
	if s > 300 {
		return "the ID is invalid"
	}
	return "the ID is valid"
}
