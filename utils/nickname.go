package utils

import "strconv"

var Nicknames1 = []string{
	"Alice", "Bella", "Charlie", "Diana", "Eva", "Fiona", "George", "Hannah", "Ivy", "Jack",
	"Karen", "Liam", "Mia", "Nathan", "Olivia", "Peter", "Quincy", "Ruby", "Sam", "Tina",
	"Ursula", "Violet", "Wendy", "Xander", "Yara", "Zane",
}

var Nicknames2 = []string{
	"Ava", "Betty", "Carol", "Dora", "Ella", "Freya", "Gabe", "Hailey", "Iris", "Jasmine",
	"Kara", "Logan", "Mason", "Nina", "Omar", "Phoebe", "Queenie", "Richy", "Sophia", "Tyler",
	"Uriel", "Vince", "Will", "Ximena", "Yvette", "Zack",
}

var DzName = "洞主"

var ExtraNamePrefix = "Anonymous"

func GenNickname(id int) string {
	switch {
	case id == 0:
		return DzName
	case id <= 26:
		return Nicknames1[id-1]
	case id <= 26*27:
		return Nicknames1[(id-1)/26-1] + " " + Nicknames2[(id-1)%26]
	default:
		return ExtraNamePrefix + strconv.Itoa(id-26*27)
	}
}
