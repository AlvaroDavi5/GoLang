package objects

import (
	"fmt"
)

func NewRunningAnimal(species string, commonName string, sound string, runningSpeed int) *runningAnimal {
	return &runningAnimal{
		Animal: Animal{
			species:    species,
			commonName: commonName,
			sound:      sound,
		},
		runningSpeed: runningSpeed,
	}
}

func NewFlyingAnimal(species string, commonName string, sound string, flyingSpeed int) *flyingAnimal {
	return &flyingAnimal{
		Animal: Animal{
			species:    species,
			commonName: commonName,
			sound:      sound,
		},
		flyingSpeed: flyingSpeed,
	}
}

func NewSwimmingAnimal(species string, commonName string, sound string, swimmingSpeed int) *swimmingAnimal {
	return &swimmingAnimal{
		Animal: Animal{
			species:    species,
			commonName: commonName,
			sound:      sound,
		},
		swimmingSpeed: swimmingSpeed,
	}
}

func (animal *Animal) Speak() {
	fmt.Printf("The %s (%s) makes a sound: %q\n", animal.commonName, animal.species, animal.sound)
}

func (ra *runningAnimal) Run() {
	fmt.Printf("The %s runs at %d km/h\n", ra.commonName, ra.runningSpeed)
}

func (fa *flyingAnimal) Fly() {
	fmt.Printf("The %s flies at %d km/h\n", fa.commonName, fa.flyingSpeed)
}

func (sa *swimmingAnimal) Swim() {
	fmt.Printf("The %s swims at %d km/h\n", sa.commonName, sa.swimmingSpeed)
}
