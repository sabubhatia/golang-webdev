package course

type course struct {
	Number string 
	Name string
	Units uint8
}

type semester struct {
	Term string
	Year uint16
	Courses []course
}

type year struct {
	Fall semester 
	Spring semester 
	Summer semester
}


func GetCourses(yr uint16) *year {
	y := year {
		Fall: semester {
			Term: "Fall",
			Year: yr,
			Courses: []course{
				{"CS-01-1", "Introduction to Go Programming", 2},
				{"CS-02-1", "Introduction to Operating Systems", 4},
				{"CS-03-1", "Introduction to Compilers", 4},
				{"CS-04-1", "Introduction to Computer Architecture", 4},
				{"CS-05-1", "Theory of Computer Science Part 1", 4},
			},
		},
		Spring: semester{
			Term: "Spring",
			Year: yr,
			Courses: []course{
				{"CS-01-2", "Advanced Go Programming", 2},
				{"CS-02-2", "Advanced Operating Systems", 4},
				{"CS-03-2", "Advanced Compilers", 4},
				{"CS-04-2", "Advanced Computer Architecture", 4},
				{"CS-05-2", "Theory of Computer Science Part 2", 4},	
			},
		},
	}

	return &y
}