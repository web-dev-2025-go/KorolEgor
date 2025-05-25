package modals

import "fmt"

type Person struct {
	Name  string
	Age   uint8
	Hobby string
}

func NewPerson(name string, age uint8, hobby string) Person {
	return Person{Name: name, Age: age, Hobby: hobby}
}

type Friend struct {
	Person Person
	Phone  string
}

func NewFriend(name string, age uint8, hobby, phone string) *Friend {
	return &Friend{Person: NewPerson(name, age, hobby), Phone: phone}
}

func (f *Friend) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, phone: %s",
		f.Person.Name, f.Person.Age, f.Person.Hobby, f.Phone)
}

func (f *Friend) Greet() {
	fmt.Println("Привіт від друга:", f.Person.Name)
}

func (f *Friend) SetHobby(hobby string) {
	f.Person.Hobby = hobby
}

type Me struct {
	Person  Person
	Job     string
	Logined bool
}

func NewMe(name string, age uint8, hobby, job string) *Me {
	return &Me{Person: NewPerson(name, age, hobby), Job: job}
}

func (m *Me) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, job: %s",
		m.Person.Name, m.Person.Age, m.Person.Hobby, m.Job)
}

func (m *Me) Greet() {
	fmt.Println("Привіт від мене:", m.Person.Name)
}

func (m *Me) SetHobby(hobby string) {
	m.Person.Hobby = hobby
}

func (m *Me) Login() {
	m.Logined = true
}

func (m *Me) Logout() {
	m.Logined = false
}

type Relative struct {
	Person Person
	Note   string
}

func NewRelative(name string, age uint8, hobby, note string) *Relative {
	return &Relative{Person: NewPerson(name, age, hobby), Note: note}
}

func (r *Relative) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, note: %s",
		r.Person.Name, r.Person.Age, r.Person.Hobby, r.Note)
}

func (r *Relative) Greet() {
	fmt.Println("Привіт від родича:", r.Person.Name)
}

func (r *Relative) SetHobby(hobby string) {
	r.Person.Hobby = hobby
}
