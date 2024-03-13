package main

import (
  "fmt"
)

// Интерфейс, описывающий состояние человека
type HumanState interface {
  PerformAction()
}

// Конкретное состояние - спокойное
type CalmState struct{}

func (c *CalmState) PerformAction() {
  fmt.Println("Человек спокоен.")
}

// Конкретное состояние - взволнованное
type ExcitedState struct{}

func (e *ExcitedState) PerformAction() {
  fmt.Println("Человек взволнован.")
}

// Контекст, представляющий человека и его текущее состояние
type Human struct {
  state HumanState
}

// Метод для установки нового состояния
func (h *Human) SetState(newState HumanState) {
  h.state = newState
}

// Метод для выполнения действия в зависимости от текущего состояния
func (h *Human) PerformAction() {
  h.state.PerformAction()
}

// func main() {
//   // Создаем объект человека и устанавливаем начальное состояние
//   human := &Human{}
//   human.SetState(&CalmState{})

//   // Выполняем действие в зависимости от текущего состояния
//   human.PerformAction()

//   // Изменяем состояние и выполняем другое действие
//   human.SetState(&ExcitedState{})
//   human.PerformAction()
// }