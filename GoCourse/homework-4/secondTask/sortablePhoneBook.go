package main

import (
	"fmt"
	"sort"
)

type ContactInfo struct {
	Name         string `json:"name"`
	PhoneNumbers []int  `json:"phoneNumbers"`
	Country      string `json:"country"`
	City         string `json:"city"`
}

type PhoneBook []ContactInfo

func (ph PhoneBook) Len() int {
	return len(ph)
}

func (ph PhoneBook) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph PhoneBook) Less(i, j int) bool {
	return ph[i].Name < ph[j].Name
}

func main() {
	phoneBook := PhoneBook{}

	firstPerson := ContactInfo{
		Name:         "Sergey",
		PhoneNumbers: []int{89999999999, 89991231231},
		Country:      "Russia",
		City:         "Ramenskoye",
	}

	secondPerson := ContactInfo{
		Name:         "Katya",
		PhoneNumbers: []int{89999999992, 89992351231},
		Country:      "Russia",
		City:         "Luberci",
	}

	thirdPerson := ContactInfo{
		Name:         "Anatolie",
		PhoneNumbers: []int{89999999932, 89992551231},
		Country:      "Russia",
		City:         "Kazani",
	}

	phoneBook = append(phoneBook, firstPerson)
	phoneBook = append(phoneBook, secondPerson)
	phoneBook = append(phoneBook, thirdPerson)

	sort.Sort(phoneBook)

	fmt.Println(phoneBook)

}
