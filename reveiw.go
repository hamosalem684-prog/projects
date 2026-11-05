package main 
import (
	"fmt"
)
func main(){
	n := 80 // this can't be true but with const it can be. i just found that out    
	first := 4.5
	if first >= n {
		fmt.Println("pass")
	}else {
		fmt.Println("fail")
	}
}

// // //################################################### THE DEFERENCE BIYWEEN CONST AND INT

func main(){
 const n = 80 // !!
	first := 4.3
	if first >= n {
		fmt.Println("pass")
	}else {
		fmt.Println("fail")
	}
}

// // //###################################################

const s = 100  // can be represented as it needed like   
int8
int16
int32
int64
float64

// // //###################################################

func main(){
	const s = 100 // so this is in fact untyped const it can be represented when it needed 
	var a int8 = s // s becomes int8 
	var i float64 = s // s becomed float64 
}

// // //####################################################

func main(){
	const s int = 100 // its typed 
	n := 40
	if n <= s {
		fmt.Println("ok not bad")
	}else {
		fmt.Println("no vary bad")
	}
}

// int8 // has 8 bits
// int64 // has 64 bits
// int32 // 32 bits
// int is something else i'm still working on it 

// #################################################


var name string = "gana"  // local functoin
var number int = 10
const h = 100 // this is untyped const not only that it can be represented as it needed
func f(){
	fmt.Println("glabal function",name)
}
func g(){
	fmt.Println("glabal function ",number)
}

func main(){
	var i int = 20 // global function
	var s string = "hanan"
	var t float64 = 22.3// that simply means float with 64bits 22.3 
	fmt.Println("the const is ",h + t)
	fmt.Println("local function ",i)
	fmt.Println("lacal function ",s)
	f() // local funtion
	g() // the same 
}

//###################################################

func main(){
	B := "A" 
	R := 'A'
	s := "A"
	for _, a := range B{
		fmt.Println("the byte ",a) // this is a byte that represents 65 which is ASCCI number 
	}
	fmt.Println("the rune ",R) // this is rune that represents 65 though which is the number of the letter 
	fmt.Println("the string",s) // the string "A"
}

//#########################################################

func slice(a,i int)[]int{
	if a >= i {
			return nil // there is no value
		}
	var answer []int
	for j := a ; j < i ; j++{
		answer = append(answer,j)
	}
	return answer
}
func main(){
	fmt.Println(slice(5,10))
	fmt.Println(slice(10,5))
	localslice([]int {1,2,3,4,5})
}

func localslice(a []int){ // local function
	fmt.Println(a)
}



// ######################################


func main(){
	f := 5
	a := make([]int,f) // [0 0 0 0 0] + 1
	for s := 0 ; s < f ; s++{
		a[s] = s + 1 // this works like this 0 + 1 = 1 so 1 + 1 = 2 all the the way to 5
	}
	fmt.Println(a)
}


//##############################################################

func main(){
	fmt.Println(slice(5,10))
}

func slice(s,i int)[]int{
	h := i - s // this to get me what bitween this which is 5 i think is that right sure ?
	box := make([]int,h) // i understand this by the way 
	for a := s ; a < i ; a++{
		// box[a] means [0 0 0 0 0] and i should put my numbers in those exeplain what a more to me 
		box[a - s] = a // why i takse the positoin number 0 and takes 5 why  
	}
	return box // a - s = 0 so = a 5 position 0 so they became [5 6 7 8 9]
}

//################################################

func main(){
	slice := make([]int,10,20)
	fmt.Println(len(slice),cap(slice))

	var p []int 
	fmt.Println("empty",p ) // [ ]

	p = append(p,5) // [5]
	fmt.Println(p)

	p = append(p,2,3,4,5) // [2 3 4 5]
	fmt.Println("at a time",p)

	box := make([]int,5)
	for a := 0 ; a < 5 ; a ++{
		box[a] = a + 1 // a on the left side is it the same a on the right side ?
	} // [1 2 3 4 5]
	fmt.Println(box)
}


//######################################
func main(){
	s := 5
/ 	slice := make([]int,s)
	for a := 0 ; a < s ; a++{
		slice[a] = 1 // [1 1 1 1 1]
	}
	fmt.Println(slice)

	s := 5 
	box := make([]int,s)
	for a := 0 ; a < 5 ; a++{ // [0 10 20 30 40]
		box[a] = a * 10
	}
	fmt.Println(box)

	s := 5 // [5 10 15 20 25]
	box := make([]int, s)
	for a := 1  ; a < s ; a++{
		box[a] = a * 5
	}
	fmt.Println(box)
}

//###########################################################

func main(){
	fmt.Println(procces(5))
    fmt.Println(slice(20,25))
	fmt.Println(slice(3,8))
}

func procces(s int)[]int{ 
	box := make([]int,s)
	for a := 0 ; a < s ; a++{
		box[a] = s - a  // [5 4 3 2 1]
	}
	return box
}

func slice(s,a int)[]int{
	result := a - s 
	box := make([]int,result)
	for i := s ; i < a ; i++{
		box[i-s] = i
	}
	return box
}
// rangeSlice(20, 25) -> [20 21 22 23 24]
//rangeSlice(3, 8)   -> [3 4 5 6 7]


//###############################################


func main(){
	var ( //  a new way of putting things it's like any way when you initlize or create int or string
		name = "ahemd"
		age = 18
	)

	name2,age2 := "ahemd",18 // many ways of creating a varaible

	f := 15 | 10 // this to compare two bits meaning give me the one that has 1 bit if both that's ok 
	b := 15 & 10 // compare 15 and 10 both must be 1 bit can you see &  
	k := 10 >> 1 // this means 10 / 2 
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(k)
	fmt.Println(name,age)
	fmt.Println(name2,age2)

}

//################################################


func main(){
	password := "Hamo321"
	username := "ahmedsalem"
	if password == "Hamo321" { // nasted if is useful when you wanna say two conditions not just one and the thing is the first condition 
		if username == "salem" { // must be correct so the next work
			    fmt.Println("welcome")
		    }else{
			    fmt.Println("try again") // I'll Print("try again") the next one isn't right 
		}
	}
}

//################################################

func str(s []string)string{
	var i string
	for h , a := range s{
		i += a // "Hello", "how", "are", "you?"
		if h < len(s)-1{ // don't put space after the last byte
			i += "\n"
		}
	}
	return i
}


func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(str(test))
}


//################################################


func main(){
	// s := []string{"omar","ali"}
	for _, r := range []string{"omar","ali"} {
		fmt.Println(string(r))
	}
}


//###################################################################


func main(){
	fmt.Println(str("hamo salme"," ahmed salem"))
	f := "m"
	s := len(f)
	fmt.Print(s)
	// fmt.Printf("%T and %v",f,f)
}

func str(t,r string)[]string{
	box := make([]string,len(t) + len(r)) // it can be len([]rune(t)) + len([]rune(r)) both gives me the same result this is an example f := "m" 
	for a , s := range t { // s := len(f) = 1 byte or len([]rune(f)) = 1 rune 
		box[a] = string(s)
	}
	keepup := len([]rune(t)) // this means continue from hamo salem to ahmed salem not from index[0]

	for h, i := range r{
		box[keepup + h] = string(i)
	} 
	return box
}


