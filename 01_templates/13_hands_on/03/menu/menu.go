package menu

import "log"

const (
	Breakfast = "Breakfast"
	Lunch     = "Lunch"
	Dinner    = "Dinner"
)

type item struct {
	Name        string
	Description string
	Price       float64
}
type meal struct {
	Name  string
	Items []item
}

type menu struct {
	Restaurant string
	Breakfast  meal
	Lunch      meal
	Dinner     meal
}

type menus []menu

func Menus(restaurants ...string) menus {
	m := menus{
		menu{
			Restaurant: "The Joint",
			Breakfast:  breakfast(),
			Lunch:      lunch(),
			Dinner:     dinner(),
		},
		menu{
			Restaurant: "Breakfast Club",
			Breakfast:  breakfast(),
		},
		menu{
			Restaurant: "The SpeakEasy",
			Dinner:     dinner(),
		},
		menu{
			Restaurant: "WorkPlace",
			Breakfast:  breakfast(),
			Lunch:      lunch(),
		},
	}

	return m
}

func Menu(restaurant string) menu {
	if len(restaurant) < 1 {
		log.Panicf(("Empty strings are not allwoed for reataurant names\n"))
	}

	m := menu{
		Restaurant: restaurant,
		Breakfast:  breakfast(),
		Lunch:      lunch(),
		Dinner:     dinner(),
	}

	return m
}

func breakfast() meal {
	m := meal{
		Name: Breakfast,
		Items: []item{
			{"Waffles", "Fruit waffles of seasonal fruits", 8.99},
			{"Eggs & Toast", "Eggs your way with choice of toast", 7.99},
			{"American breakfast", "Eggs, toast, bacon, ham, tomato", 11.28},
		},
	}

	return m
}

func lunch() meal {
	m := meal{
		Name: Lunch,
		Items: []item{
			{"Soup & roll", "Soup of the day with a hard sourdough roll", 5.99},
			{"Pasta Special", "Pasta of the day", 9.99},
			{"Chicken chow mein", "Egg noodles with chicken", 7.99},
		},
	}

	return m
}

func dinner() meal {
	m := meal{
		Name: Dinner,
		Items: []item{
			{"Steak AU Poivré", "Steak with black pepper cooked your way", 23.99},
			{"Egg plant parmesan", "Cooked to perfection with fresh farm grown italian egg plants", 18.99},
			{"Fish & Chips", "Sole fillet in a beer batetr cooked golden brown", 27.99},
		},
	}

	return m
}
