Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
вывод 2, 1
если отложенная функция является литералом функции , а окружающая функция имеет именованные параметры результата , которые находятся в области действия литерала, отложенная функция может получить доступ к результирующим параметрам и изменить их до того, как они будут возвращены. Поэтому в первом случае мы получаем значение - 2.

```