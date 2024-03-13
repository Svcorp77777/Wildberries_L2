package main

import "fmt"

// HumanHandler интерфейс обработчика для цепочки вызовов
type HumanHandler interface {
  HandleRequest(request string) bool
  SetNextHandler(handler HumanHandler)
}

// BaseHumanHandler базовая реализация обработчика
type BaseHumanHandler struct {
  nextHandler HumanHandler
}

func (b *BaseHumanHandler) SetNextHandler(handler HumanHandler) {
  b.nextHandler = handler
}

// GenderHandler обработчик для определения пола человека
type GenderHandler struct {
  BaseHumanHandler
}

func (gh *GenderHandler) HandleRequest(request string) bool {
  if request == "Gender" {
    fmt.Println("Определен пол человека")
    return true
  } else if gh.nextHandler != nil {
    return gh.nextHandler.HandleRequest(request)
  }
  return false
}

// AgeHandler обработчик для определения возраста человека
type AgeHandler struct {
  BaseHumanHandler
}

func (ah *AgeHandler) HandleRequest(request string) bool {
  if request == "Age" {
    fmt.Println("Определен возраст человека")
    return true
  } else if ah.nextHandler != nil {
    return ah.nextHandler.HandleRequest(request)
  }
  return false
}

// Пример использования паттерна "Цепочка вызовов"
// func main() {
//   // Создаем обработчики
//   genderHandler := &GenderHandler{}
//   ageHandler := &AgeHandler{}

//   // Устанавливаем цепочку обработчиков
//   genderHandler.SetNextHandler(ageHandler)

//   // Запускаем цепочку вызовов
//   genderHandler.HandleRequest("Gender")
//   genderHandler.HandleRequest("Age")
//   genderHandler.HandleRequest("Name")
// }