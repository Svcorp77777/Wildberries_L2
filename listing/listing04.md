Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Программа выведет цифры от 0 до 9 и произайдёт deadlock.

Будет создан канал ch, далее запустица анонимная горутина, которая пишет
в канал числа от 0 до 9, после чего основная горутина читает из канала
и выводит значения.
И так как нет явного закрытия канала ch произойдет deadlock на строке 
for n := range ch потому что будет ожидать дополнительных данных.
Для решения проблемы нужно использовать close(ch) чтобы закрыть канал
после записи данных в него.
```