package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//findian()
	//go_slice()
	//makejson()
	//read()
}

/*Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.*/

func findian() {
	var text string
	fmt.Print("Write a string")
	fmt.Scan(&text)
	aux := strings.ToLower(text)
	if strings.Contains(aux, "a") && strings.HasPrefix(aux, "i") && strings.HasSuffix(aux, "n") {
		fmt.Print("Found\n")
	} else {
		fmt.Print("Not foud\n")
	}
}

/*Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.*/
func go_slice() {
	my_slice := make([]int, 3)
	var aux string
	var i int
	var sum int
	for {
		fmt.Print("Write an integer:\n")
		fmt.Scan(&aux)
		if aux == "x" {
			break
		} else {
			j, err := strconv.Atoi(aux)
			if err != nil {
				log.Fatal(err)
			}
			my_slice[i] = j
			sum += my_slice[i]
		}
		i++
	}
	fmt.Print("My slice: ", my_slice, "\n")
	sort.Sort(sort.IntSlice(my_slice))
	fmt.Print("My sort slice: ", my_slice, "\n")
	fmt.Print("The sumatory", sum, "\n")
}

/*Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.*/
func makejson() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Write your name: \n")
	name, _ := scanner.ReadString('\n')
	fmt.Print("Write your adress: \n")
	adress, _ := scanner.ReadString('\n')
	person := make(map[string]string)
	person["name"] = name
	person["adress"] = adress
	person_json, err := json.Marshal(person)
	if err != nil {
		fmt.Println("ERROR ", err)
	} else {
		fmt.Printf("%s", person_json)
	}
}

/*Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).
*/
func read() {
	type Names struct {
		f_name string
		l_name string
	}
	var file_name string
	var i = 0
	var persons []Names
	fmt.Print("Write the directory")
	fmt.Scanln(&file_name)
	content, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	aux := strings.Fields(string(content))
	for i < len(aux) {
		persons = append(persons, Names{f_name: aux[i], l_name: aux[i+1]})
		i += 2
	}
	for j := 0; j < len(persons); j++ {
		fmt.Print(persons[j])
	}
}

// Calcular el factorial de un numero dado
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
