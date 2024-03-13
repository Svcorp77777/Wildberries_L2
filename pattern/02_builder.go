package pattern

// Human - структура, которую мы строим
type Human struct {
	Name    string
	Age     int
	Gender  string
	Height  float64
	Weight  float64
	Address string
}

// HumanBuilder - интерфейс для пошагового построения объекта Human
type HumanBuilder interface {
	SetName(name string) HumanBuilder
	SetAge(age int) HumanBuilder
	SetGender(gender string) HumanBuilder
	SetHeight(height float64) HumanBuilder
	SetWeight(weight float64) HumanBuilder
	SetAddress(address string) HumanBuilder
	Build() Human
}

// ConcreteHumanBuilder - конкретная реализация интерфейса HumanBuilder
type ConcreteHumanBuilder struct {
	human Human
}

func NewConcreteHumanBuilder() *ConcreteHumanBuilder {
	return &ConcreteHumanBuilder{}
}

func (b *ConcreteHumanBuilder) SetName(name string) HumanBuilder {
	b.human.Name = name
	return b
}

func (b *ConcreteHumanBuilder) SetAge(age int) HumanBuilder {
	b.human.Age = age
	return b
}

func (b *ConcreteHumanBuilder) SetGender(gender string) HumanBuilder {
	b.human.Gender = gender
	return b
}

func (b *ConcreteHumanBuilder) SetHeight(height float64) HumanBuilder {
	b.human.Height = height
	return b
}

func (b *ConcreteHumanBuilder) SetWeight(weight float64) HumanBuilder {
	b.human.Weight = weight
	return b
}

func (b *ConcreteHumanBuilder) SetAddress(address string) HumanBuilder {
	b.human.Address = address
	return b
}

func (b *ConcreteHumanBuilder) Build() Human {
	return b.human
}

// func main() {
// 	// Использование шаблона строителя
// 	builder := NewConcreteHumanBuilder()

// 	human := builder.SetName("John").
// 		SetAge(25).
// 		SetGender("Male").
// 		SetHeight(180.0).
// 		SetWeight(75.0).
// 		SetAddress("123 Main St").
// 		Build()

// 	fmt.Printf("Created Human: %+v\n", human)
// }
