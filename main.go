package main

import (
	"MessageGO/repository"
	"MessageGO/transport"
	"fmt"
)

func main() {
	fmt.Println("start server MessageGO")
	//Создание таблицы бд
	repository.CreateDataBaseTable()
	//запуск роутера
	transport.RouterHandler(":3500")
}
