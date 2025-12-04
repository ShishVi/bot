package product

type Product struct {
	Title    string
	Calories int
}

var allProducts = []Product{
	{Title: "картошка", Calories: 356},
	{Title: "пицца", Calories: 956},
	{Title: "кока-кола", Calories: 160},
	{Title: "сыр", Calories: 550},
	{Title: "масло", Calories: 460},
}
