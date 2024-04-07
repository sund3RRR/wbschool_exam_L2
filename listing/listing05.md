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
```

Так как интерфейс содержит два поля dynamicType и dynamicValue, dynamicType в данном случае != nil, что приводит к тому,
что ошибка не проходит проверку на nil.