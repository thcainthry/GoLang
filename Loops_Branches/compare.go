// comparision operators 

package main

import(
	"fmt"
	)
func main(){
	fmt.Println("No Minors allowed here.")

	var age = 41
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v \n", age,minor)
}
