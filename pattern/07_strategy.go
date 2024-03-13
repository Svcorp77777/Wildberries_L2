package main

import "fmt"

// Интерфейс стратегии
type BehaviorStrategy interface {
  Act()
}

// Реализация конкретной стратегии 1
type FriendlyBehavior struct{}

func (f *FriendlyBehavior) Act() {
  fmt.Println("Приветливое приветствие!")
}

// Реализация конкретной стратегии 2
type AggressiveBehavior struct{}

func (a *AggressiveBehavior) Act() {
  fmt.Println("Агрессивное приветствие. Будь осторожен!")
}

// Контекст
type Human struct {
  behavior BehaviorStrategy
}

func (h *Human) SetBehavior(behavior BehaviorStrategy) {
  h.behavior = behavior
}

func (h *Human) Interact() {
  h.behavior.Act()
}

// func main() {
//   // Создаем объект человека
//   human := &Human{}

//   // Устанавливаем стратегию приветствия
//   friendlyBehavior := &FriendlyBehavior{}
//   human.SetBehavior(friendlyBehavior)

//   // Человек приветствует дружелюбно
//   human.Interact()

//   // Меняем стратегию на агрессивную
//   aggressiveBehavior := &AggressiveBehavior{}
//   human.SetBehavior(aggressiveBehavior)

//   // Теперь человек приветствует агрессивно
//   human.Interact()
// }