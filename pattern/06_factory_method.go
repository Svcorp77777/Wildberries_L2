package main

import "fmt"

// Интерфейс для создания объекта Human
type HumanCreator interface {
  CreateHuman() Human
}

// Интерфейс, представляющий человека
type Human interface {
  Speak()
}

// Реализация структуры для представления мужчины
type Man struct{}

func (m Man) Speak() {
  fmt.Println("I am a man.")
}

// Реализация структуры для представления женщины
type Woman struct{}

func (w Woman) Speak() {
  fmt.Println("I am a woman.")
}

// Реализация фабричного метода для создания мужчин
type ManCreator struct{}

func (mc ManCreator) CreateHuman() Human {
  return Man{}
}

// Реализация фабричного метода для создания женщин
type WomanCreator struct{}

func (wc WomanCreator) CreateHuman() Human {
  return Woman{}
}

// func main() {
//   // Использование фабричных методов для создания объектов Human
//   manFactory := ManCreator{}
//   man := manFactory.CreateHuman()
//   man.Speak()

//   womanFactory := WomanCreator{}
//   woman := womanFactory.CreateHuman()
//   woman.Speak()
// }