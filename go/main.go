package main

import (
	"design-patterns/creational/abstractfactory"
	"design-patterns/creational/builder"
	"design-patterns/creational/factorymethod"
	"design-patterns/creational/prototype"
	"design-patterns/creational/singleton"
	"design-patterns/structural/adapter"
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

func runPrototype() {
	// Create models
	lrModel := &prototype.MLModel{}
	lrModel.SetModelType(prototype.LR)
	lrModel.SetDescription("Linear Regression Model")
	lrModel.SetTrainingSplit(0.7)
	lrModel.SetValidationSplit(0.3)
	lrModel.SetAlpha(0.01)
	lrModel.SetBeta(0.1)

	svmModel := &prototype.MLModel{}
	svmModel.SetModelType(prototype.SVM)
	svmModel.SetDescription("Support Vector Machine Model")
	svmModel.SetTrainingSplit(0.6)
	svmModel.SetValidationSplit(0.4)
	svmModel.SetAlpha(0.02)
	svmModel.SetBeta(0.2)

	// Register models
	registry := prototype.NewModelRegistry()
	registry.RegisterModel(lrModel)
	registry.RegisterModel(svmModel)

	fmt.Printf("%+v\n", lrModel)
	fmt.Printf("%+v\n", svmModel)

	// Retrieve and clone models
	clonedLRModel := registry.GetModel(prototype.LR)
	clonedSVMModel := registry.GetModel(prototype.SVM)
	clonedLRModel.SetDescription("Linear Regression Model Cloned")
	clonedSVMModel.SetDescription("Support Vector Machine Model Cloned")

	fmt.Printf("%+v\n", clonedLRModel)
	fmt.Printf("%+v\n", clonedSVMModel)
}

func runFactoryMethod() {
	fmt.Println("Using RoundButtonFactory:")
	roundButtonFactory := &factorymethod.RoundButtonFactory{}
	factorymethod.ClientCode(roundButtonFactory)

	fmt.Println("\nUsing SquareButtonFactory:")
	squareButtonFactory := &factorymethod.SquareButtonFactory{}
	factorymethod.ClientCode(squareButtonFactory)
}

func runAbstractFactory() {
	fmt.Println("Using DarkThemeFactory:")
	darkFactory := &abstractfactory.DarkThemeFactory{}
	abstractfactory.ClientCode(darkFactory)

	fmt.Println("\nUsing LightThemeFactory:")
	lightFactory := &abstractfactory.LightThemeFactory{}
	abstractfactory.ClientCode(lightFactory)
}

func runAdapter() {
	paymentRequest := adapter.PaymentRequest{
		Name:   "John Doe",
		Phone:  "1234567890",
		Email:  "john@example.com",
		Amount: 100,
	}

	cashfreeProvider := adapter.NewCashFreePayProvider()
	adapter.ProcessPayment(cashfreeProvider, paymentRequest)
}

func main() {
	//runSingleton()
	//runBuilder()
	//runPrototype()
	//runFactoryMethod()
	//runAbstractFactory()

	runAdapter()
}
