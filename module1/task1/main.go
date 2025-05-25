package main

import (
	"KorolEgor/module1/task1/modals"
	"fmt"
)

type Describer interface {
	Describe() string
}

func main() {
	meUser := modals.NewMe("Вовчик", 19, "Клуби", "Розробник")
	friendUser := modals.NewFriend("im_coderrr", 20, "gin", "daryana, maybe baybe")
	relativeUser := modals.NewRelative("Olexandr", 67, "beer", "Тітанік")

	people := []Describer{meUser, friendUser, relativeUser}
	for _, p := range people {
		fmt.Println(p.Describe())
	}

	meUser.Greet()
	friendUser.Greet()
	relativeUser.Greet()

	meUser.SetHobby("Клуби")
	friendUser.SetHobby("gin")
	relativeUser.SetHobby("vaib")

	meUser.Login()
	fmt.Printf("Logined (після Login): %v\n", meUser.Logined)

	copyMe := *meUser
	copyMe.Logout()
	fmt.Printf("Logined (після Logout копії): %v\n", meUser.Logined)
}
