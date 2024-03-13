package main

import "fmt"

// Command interface определяет общий метод для выполнения команды.
type Command interface {
  Execute()
}

// HumanReceiver представляет получателя команды, который выполняет конкретные действия.
type HumanReceiver struct {
  Name string
}

// ConcreteCommand представляет конкретную команду, связанную с HumanReceiver.
type JumpCommand struct {
  Receiver *HumanReceiver
}

// Execute реализует метод Execute для команды JumpCommand.
func (c *JumpCommand) Execute() {
  fmt.Printf("%s прыгает!\n", c.Receiver.Name)
}

// Invoker представляет объект, который инициирует выполнение команды.
type Invoker struct {
  Command Command
}

// ExecuteCommand вызывает метод Execute для текущей команды.
func (invoker *Invoker) ExecuteCommand() {
  invoker.Command.Execute()
}

// func main() {
//   // Создаем получателя команды (человека)
//   human := &HumanReceiver{Name: "Иван"}

//   // Создаем конкретную команду и передаем ей получателя
//   jumpCommand := &JumpCommand{Receiver: human}

//   // Создаем исполнителя команды (Invoker) и устанавливаем ему команду
//   invoker := &Invoker{Command: jumpCommand}

//   // Выполняем команду
//   invoker.ExecuteCommand()
// }
