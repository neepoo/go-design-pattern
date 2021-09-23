package decorator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}
	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	require.Containsf(t, pizzaResult, expectedText, "When calling the add ingredient of the pizza decorator it "+
		"must return the text %s the expected text, not '%s'", pizzaResult,
		expectedText)
}

func TestPOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	require.Errorf(t, err, "When calling AddIngredient on the onion decorator without "+
		"an IngredientAdder on its Ingredient field must return an error, not a "+
		"string with '%s'", onionResult)

	onion.Ingredient = &PizzaDecorator{}
	onionResult, err = onion.AddIngredient()
	require.NoError(t, err)
	require.Containsf(t, onionResult, "onion", "When calling the add ingredient of the onion decorator it "+
		"must return a text with the word 'onion', not '%s'", onionResult)
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()
	require.Errorf(t, err, "When calling AddIngredient on the meat decorator without "+
		"an IngredientAdder in its Ingredient field must return an error, "+
		"not a string with '%s'", meatResult)

	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	require.NoError(t, err)
	require.Containsf(t, meatResult, "meat", "When calling the add ingredient of the meat decorator it "+
		"must return a text with the word 'meat', not '%s'", meatResult)
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients: meat, onion"
	require.Containsf(t, pizzaResult, expectedText, "When asking for a pizza with onion and meat the returned "+
		"string must contain the text '%s' but '%s' didn't have it",
		expectedText, pizzaResult)

	t.Log(pizzaResult)
}
