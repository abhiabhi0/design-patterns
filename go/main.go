package main

import (
	"design-patterns/creational/builder"
	"design-patterns/creational/singleton"
	"fmt"
)

func runSingleton() {
	for i := 0; i < 100; i++ {
		//go singleton.GetInstance()
		go singleton.GetInstanceUsingSync()
	}

	fmt.Scanln()
}

func runBuilder() {
	dbBuilder := builder.NewDatabaseBuilder()
	db, err := dbBuilder.WithName("mysql").WithUrl("localhost", 3306).Build()
	if err != nil {
		fmt.Printf("err %v", err)
		return
	}
	fmt.Printf("%+v\n", db)
}

func main() {
	//runSingleton()
	runBuilder()
}
