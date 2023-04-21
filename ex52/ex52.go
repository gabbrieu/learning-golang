//- Partindo do código abaixo, ordene os []user por idade e sobrenome.

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// your code goes here
type sortByAge []user

func (user sortByAge) Len() int {
	return len(user)
}

func (user sortByAge) Less(i, j int) bool {
	return user[i].Age < user[j].Age
}

func (user sortByAge) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

type sortByLastName []user

func (user sortByLastName) Len() int {
	return len(user)
}

func (user sortByLastName) Less(i, j int) bool {
	return user[i].Last < user[j].Last
}

func (user sortByLastName) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(sortByAge(users))
	fmt.Printf("\n\n%v", users)

	sort.Sort(sortByLastName(users))
	fmt.Printf("\n\n%v", users)
}
