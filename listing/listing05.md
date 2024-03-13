Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Функция test() имеет блок кода который что то делает, после этого
возвращается nil что является значением типа указателя на customError.

В функции main() объявлена переменная err типа error, далее вызывается 
функция test() и результат присваивается переменной err.
Но err не будет равна nil, т.к. err будет хранить указатель на тип customError,
даже если внутри функции test возвращается nil. Поэтому выведется error
```