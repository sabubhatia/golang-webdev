package maps

import (
	"testing"
)

// Type of name and what things they like
type person struct {
	name  string
	likes []string
}

var people []*person
var likes = map[string][]*person{}

func TestMaps(t *testing.T) {
	people = createPeople(2)
	createLikes(people)
	got := likes["Prawn"]
	t.Log(got)
	if len(got) != 1 {
		t.Errorf("Got: %d, Want: 1", len(got))
	}
}

func createLikes(p []*person) {
	// a nil slice is treated as a zero lenggth slice by range and by len().
	for _, v := range p {
		for _, l := range v.likes {
			likes[l] = append(likes[l], v) // append to the slice and create map entry !
		}
	}
}

func createPeople(num int) []*person {
	s := []*person{
		{
			name:  "Sabu",
			likes: []string{"Pheng", "Chciken", "Fish", "Noodle", "Mutton"},
		},
		{
			name:  "Pheng",
			likes: []string{"Sabu", "Chicken", "Fish", "Noodle", "Prawn"},
		},
	}

	return s
}
