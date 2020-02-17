package main

import (
	"encoding/json"
	"io/ioutil"
)

type PersonInfo struct {
	Name         string `json:"name"`
	PhoneNumbers []int  `json:"phoneNumbers"`
	Country      string `json:"country"`
	City         string `json:"city"`
}

type PhoneBook struct {
	PersonsInfo []PersonInfo `json:"personsInfo"`
}

func (phoneBook *PhoneBook) Save() error {
	jsonPhoneBook, err := json.Marshal(phoneBook)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile("phoneBook.json", jsonPhoneBook, 0600); err != nil {
		return err
	}
	return nil
}

func main() {
	phoneBook := PhoneBook{
		PersonsInfo: []PersonInfo{},
	}

	firstPerson := PersonInfo{
		Name:         "Sergey",
		PhoneNumbers: []int{89999999999, 89991231231},
		Country:      "Russia",
		City:         "Ramenskoye",
	}

	secondPerson := PersonInfo{
		Name:         "Katya",
		PhoneNumbers: []int{89999999992, 89992351231},
		Country:      "Russia",
		City:         "Luberci",
	}

	phoneBook.PersonsInfo = append(phoneBook.PersonsInfo, firstPerson)
	phoneBook.PersonsInfo = append(phoneBook.PersonsInfo, secondPerson)

	phoneBook.Save()
}
