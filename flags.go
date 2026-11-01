package main
import "fmt"

func main(){
	password := ""
	check := len(password) > 0 // so where is true shouldn't i put it somewhere ? 
	for _, a := range password {
		if !(a >= 'a' && a <= 'z')&&
		   !(a >= 'A' && a <= 'Z')&&
		   !(a >= '0' && a <= '9')&&
		   !(a == ' ' ){
			check = false
			break 
		   }
	}
	if check {
		fmt.Println("welcome")
	}else {
		fmt.Println("try again")
	}
}


// #######################################################


func main(){
	idenifer := "Hamo 123"
	check := len(idenifer) > 5

	for _, a := range idenifer {
		if !(a >= 'a' && a <= 'z')&&
		   !(a >= 'A' && a <= 'Z')&&
		   !(a >= '0' && a <= '9')&&
		   !(a == '_'){
			check = false 
			break 
		   }
	}

	if check {
		fmt.Println("ok valid one")
	}else {
		fmt.Println("no invalid one")
	}
}

// ##########################################
  
// found or not

func main(){
	var list []int = []int {100,200,900,700,400}
	found := true  // that means i haven't seen anything or started
	for s := 0 ; s < len(list) ; s++{
		if list[s] == 100 {
			found = false   // this will print not exist although it's exist and i wanna know how this loops and see things just how it works 
			break
		}
	}

	if found {
		fmt.Println("exist")
	}else{
		fmt.Println("not exist")
	}

	for a := 0 ; a < len(list) ; a++{
		for j := a + 1 ; j < len(list) ; j++{
			if list[a] > list[j] {
				list[a] , list[j] = list[j] , list[a]
			}
		}
	}
	list = append(list,500)
	list[len(list)-1] = 800 // i wanna say last one 
	fmt.Print(list)
}

// ################################################
// check the password


func main(){
	password := "Hamo 123"
	check := true // this means this string "Hamo 123" true now but after that i checks so that it can see the final result  

	for _, a := range password{
		if (a < 'a' || a > 'z')&& // true 
		   (a < 'A' || a > 'Z')&& // true  
		   (a < '0' || a > '9')&& // true 
		   a != ' ' { // true 
			check = false // and than the qeistion is now false = true result false which means not valid  
			break // but if all of  them false that would not make check false work and i will print valid 
		}
	}
	if check {
		fmt.Println("valid") // and i gives me this 
	}else {
		fmt.Println("not valid")
	}
}

func main(){
	number(3241)
}
func number(s int){
	sorted := true // so that means something depinding on what i will do right ? // it means let's suppose that you are sorted 
	var i []int 
	for s > 0 { // is that mean 4321 > 0 and that's true so it moves on ?
		i = append(i,s%10)
		s /= 10
		sorted = false // if something is about to be sorted stop the code afterwords give me not sorted   
	}
	if sorted {
		fmt.Println("sorted",i)
	}else{
		fmt.Println("not sorted",i)
	}
}

// some usful stuff 

found
valid
sorted
empty
loggedIn
gameOver
hasDuplicate


//###########

found := false   // not found yet
valid := true    // assume valid
sorted := true   // assume sorted
empty := true    // assume empty

