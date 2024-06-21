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
Программа создает пользовательскую ошибку customError, реализующую интерфейс error. При проверке err != nil учитывается тип err, а не его значение. Поскольку тип err не равен nil (это тип *customError), условие err != nil будет истинным, и программа выведет "error".


```
