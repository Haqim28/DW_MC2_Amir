package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profile {

	var prof Profile
	if name == "Sasuke" {
		prof.Name = name
		prof.Health = 200
		prof.Power = 100
		prof.Exp = 0
	} else if name == "Goku" {
		prof.Name = name
		prof.Health = 400
		prof.Power = 300
		prof.Exp = 100
	} else if name == "Naruto" {
		prof.Name = name
		prof.Health = 150
		prof.Power = 200
		prof.Exp = 50
	}

	return prof
}

func PowerUp(profile Profile, angka int) Profile {

	profile.Health = profile.Health + (profile.Health * angka)
	profile.Power = profile.Power + (profile.Power * angka)
	profile.Exp = profile.Exp + (profile.Exp * angka)
	return profile
}
func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("========Heros Power Up========")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}
