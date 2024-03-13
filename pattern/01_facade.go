package pattern

import "fmt"

// Subsystem 1
type Brain struct{}

func (b *Brain) Think() {
	fmt.Println("Thinking...")
}

// Subsystem 2
type Heart struct{}

func (h *Heart) Pump() {
	fmt.Println("Heart pumping...")
}

// Subsystem 3
type Lungs struct{}

func (l *Lungs) Breathe() {
	fmt.Println("Breathing...")
}

// Facade
type HumanFacade struct {
	brain *Brain
	heart *Heart
	lungs *Lungs
}

func NewHumanFacade() *HumanFacade {
	return &HumanFacade{
		brain: &Brain{},
		heart: &Heart{},
		lungs: &Lungs{},
	}
}

// Facade method
func (hf *HumanFacade) Live() {
	hf.brain.Think()
	hf.heart.Pump()
	hf.lungs.Breathe()
}

// func main() {
// 	// Client code using the facade
// 	human := NewHumanFacade()
// 	human.Live()
// }
