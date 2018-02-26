package main

import (
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
	pb "github.com/sanshitsharma/golangsamples/proto_tutorial/proto"
)

var id int32
const fname = "addressBook.txt"

type Person struct {
	Name        string
	Email       string
	PhoneNumber string
	PhoneType   string
}

func init() {
	id = 1
}

func put(people []Person) (error) {
	book := &pb.AddressBook{}
	for _, person := range people {
		pbPerson := &pb.Person{}
		pbPerson.Name = person.Name
		pbPerson.Id = id
		id++
		pbPerson.Email = person.Email
		pbPerson.Phones = []*pb.Person_PhoneNumber{{person.PhoneNumber,
			pb.Person_PhoneType(pb.Person_PhoneType_value[person.PhoneType])}}

		book.People = append(book.People, pbPerson)
	}

	// Write a new address book back to disk
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
		return err
	}

	return nil
}

func get() ([]Person, error) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return nil, err
	}

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
		return nil, err
	}

	pbPeople := book.GetPeople()
	people := []Person{}
	for _, pbPerson := range pbPeople {
		person := Person{}
		person.Name = pbPerson.GetName()
		person.Email = pbPerson.GetEmail()
		phones := pbPerson.GetPhones()
		for _, phone := range phones {
			person.PhoneNumber = phone.Number
			person.PhoneType = phone.Type.String()
		}

		people = append(people, person)
	}

	return people, nil
}

func main() {
	log.Println("Welcome to Proto Messaging Tester")

	persons := []Person{{"Sanshit Sharma", "present87@gmail.com", "513-549-1487", "MOBILE"},
		{"Nitin Mathur", "mathnitin@gmail.com", "408-859-8519", "WORK"}}
	err := put(persons)
	if err != nil {
		log.Println("Failed to write to directory")
	}


	people, err := get()
	if err != nil {
		log.Println("Failed to read from directory")
	}

	for _, person := range people {
		log.Println(person)
	}
}
