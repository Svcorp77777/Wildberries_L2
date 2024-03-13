package main

import "fmt"

// Интерфейс посетителя
type Visitor interface {
    VisitHuman(human *Human)
}

// Интерфейс элемента, поддерживающего посещение
type Element interface {
    Accept(visitor Visitor)
}

// Конкретный элемент - человек
type Human struct {
    Name   string
    Age    int
    Height float64
}

func (h *Human) Accept(visitor Visitor) {
    visitor.VisitHuman(h)
}

// Конкретный посетитель - операция "Информация о человеке"
type InfoVisitor struct{}

func (v *InfoVisitor) VisitHuman(human *Human) {
    fmt.Printf("Имя: %s, Возраст: %d, Рост: %.2f м\n", human.Name, human.Age, human.Height)
}

// Конкретный посетитель - операция "Пожелание здоровья"
type HealthVisitor struct{}

func (v *HealthVisitor) VisitHuman(human *Human) {
    fmt.Printf("Пожелание здоровья для %s!\n", human.Name)
}

// func main() {
//     // Создаем человека
//     person := &Human{
//         Name:   "Иван",
//         Age:    30,
//         Height: 1.75,
//     }

//     // Создаем посетителей
//     infoVisitor := &InfoVisitor{}
//     healthVisitor := &HealthVisitor{}

//     // Применяем операции посетителей к человеку
//     person.Accept(infoVisitor)
//     person.Accept(healthVisitor)
// }
