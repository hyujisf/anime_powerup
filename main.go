package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {

	var profile = []Profile{
		{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		},
		{
			Name:   "Goku",
			Health: 400,
			Power:  300,
			Exp:    100,
		},
		{
			Name:   "Naruto",
			Health: 150,
			Power:  200,
			Exp:    50,
		},
	}

	// var CharProfile profile
	for _, data := range profile {
		if data.Name == name {
			return data
		}
	}
	return Profile{}
}

func main() {
	profile := MakeProfile("Naruto")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 4)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)

}

func PowerUp(profile Profile, level int) Profile {
	profile.Health += level * profile.Health
	profile.Power += level * profile.Power
	profile.Exp += level * profile.Exp
	return profile
}
