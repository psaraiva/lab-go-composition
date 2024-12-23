package main

import "fmt"

type People struct {
	Name  string
	Birth string
}

type Civilian struct {
	People
	Document string
}

type Student struct {
	Civilian
	School string
}

type Teacher struct {
	Civilian
	School     string
	Discipline string
}

type Military struct {
	Civilian
	Rank string
}

func main() {
	people := People{
		Name:  "John",
		Birth: "01/01/2000",
	}

	fmt.Println(
		people,
		people.Name,
		people.Birth,
	)

	civilian := Civilian{
		People: People{
			Name:  "Wick",
			Birth: "05/10/1990",
		},
		Document: "123456789",
	}

	fmt.Println(
		civilian,
		civilian.Name,
		civilian.Birth,
		civilian.Document,
	)

	student := Student{
		Civilian: Civilian{
			People: People{
				Name:  "Anny",
				Birth: "07/09/2013",
			},
			Document: "123456789",
		},
		School: "School of Arts",
	}

	fmt.Println(
		student,
		student.Name,
		student.Birth,
		student.Document,
		student.School,
	)

	teacher := Teacher{
		Civilian: Civilian{
			People: People{
				Name:  "Rose",
				Birth: "09/12/1980",
			},
			Document: "123456789",
		},
		School:     "School of Arts",
		Discipline: "Math",
	}

	fmt.Println(
		teacher,
		teacher.Name,
		teacher.Birth,
		teacher.Document,
		teacher.School,
		teacher.Discipline,
	)

	military := Military{
		Civilian: Civilian{
			People: People{
				Name:  "Jimmy",
				Birth: "12/12/1955",
			},
			Document: "123456788",
		},
		Rank: "General",
	}

	fmt.Println(
		military,
		military.Name,
		military.Birth,
		military.Document,
		military.Rank,
	)
}
