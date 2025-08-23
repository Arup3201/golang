package person

type Person struct {
	Name     string
	Age      int
	Location string
}

func NewPerson(name string, age int, location string) Person {
	return Person{
		Name:     name,
		Age:      age,
		Location: location,
	}
}

func (p Person) SayHello(to string) string {
	return "Hello, " + to
}

func (p *Person) CelebrateBirthday() string {
	p.Age += 1
	return "Happy birthday, " + p.Name
}
