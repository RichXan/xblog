package designpattern

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	persons := []*Person{
		NewPerson("John", "Male", 25, 3000),
		NewPerson("Alice", "Female", 30, 4000),
		NewPerson("Bob", "Male", 35, 5000),
		NewPerson("Emma", "Female", 28, 3500),
		NewPerson("Mike", "Male", 22, 2500),
	}

	male := &MaleCriteria{}
	female := &FemaleCriteria{}
	ageRange := NewAgeCriteria(25, 35)
	highSalary := NewSalaryCriteria(4000)

	// 查找所有男性
	malePersons := male.MeetCriteria(persons)
	fmt.Println("Males:")
	for _, person := range malePersons {
		fmt.Printf("Name: %s, Gender: %s\n", person.name, person.gender)
	}

	// 查找年龄在25-35之间且薪资高于4000的人
	ageAndSalary := NewAndCriteria(ageRange, highSalary)
	result := ageAndSalary.MeetCriteria(persons)
	fmt.Println("\nAge between 25-35 and high salary:")
	for _, person := range result {
		fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n",
			person.name, person.age, person.salary)
	}

	// 查找女性或高薪资的人
	femaleOrHighSalary := NewOrCriteria(female, highSalary)
	result = femaleOrHighSalary.MeetCriteria(persons)
	fmt.Println("\nFemales or high salary:")
	for _, person := range result {
		fmt.Printf("Name: %s, Gender: %s, Salary: %.2f\n",
			person.name, person.gender, person.salary)
	}
}
