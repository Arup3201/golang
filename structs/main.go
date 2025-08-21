package main

import (
	"encoding/json"
	"fmt"
	"golang/structs/movie"
	"log"
)

func main() {
	var movies = []movie.Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	{
		// marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshalling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	}

	{
		// marshal indented
		data, err := json.MarshalIndent(movies, "", "	")
		if err != nil {
			log.Fatalf("JSON marshalling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

		// unmarshal
		var titles []struct {
			Title string
			Year  int `json:"released"`
		}
		if err = json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshalling failed: %s", err)
		}
		fmt.Println(titles)
	}
}

// func main() {
// 	p := person.NewPerson("Arup", 23, "Kolkata")

// 	fmt.Printf("Name: %s, Age: %d, Location: %s\n", p.Name, p.Age, p.Location)

// 	fmt.Println(p.SayHello("Subha"))
// 	fmt.Println(p.CelebrateBirthday())

// 	fmt.Printf("Name: %s, Age: %d, Location: %s\n", p.Name, p.Age, p.Location)
// }
