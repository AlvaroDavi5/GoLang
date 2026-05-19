package objects

type Animal struct {
	species    string
	commonName string
	sound      string
}
type runningAnimal struct {
	Animal
	runningSpeed int
}
type flyingAnimal struct {
	Animal
	flyingSpeed int
}
type swimmingAnimal struct {
	Animal
	swimmingSpeed int
}

type Bird interface {
	flyingAnimal
}

type Fish interface {
	swimmingAnimal
}

type TerrestrialMammal interface {
	runningAnimal
}
type AerialMammal interface {
	flyingAnimal
}
type AquaticMammal interface {
	swimmingAnimal
}
