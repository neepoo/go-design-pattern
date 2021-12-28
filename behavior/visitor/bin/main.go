package main

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float64
	GetName() string
}

// 需要访问哪些对象需要在这个接口定义好
type Visitor interface {
	Visit(ProductInfoRetriever)
}

type Visitable interface {
	Accept(Visitor)
}

type Product struct {
	Price float64
	Name  string
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

type Rice struct {
	Product
}
type Pasta struct {
	Product
}

type Fridge struct {
	Product
}

//GetPrice overrides GetPrice method of Product type
func (f *Fridge) GetPrice() float64 {
	return f.Product.Price + 20
}

//Accept overrides "Accept" method from Product and implements the Visitable
//interface
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

type PriceVisitor struct {
	Sum float64
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

func main() {
	products := make([]Visitable, 3)
	products[0] = &Rice{Product{
		Price: 32,
		Name:  "Some rice",
	}}
	products[1] = &Pasta{Product{
		Price: 40,
		Name:  "Some pasta",
	}}
	products[2] = &Fridge{
		Product: Product{
			Price: 50,
			Name:  "A fridge",
		},
	}
	priceVisitor := &PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}
	// 32 + 40 + (50+20) = 142
	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	//Print the products list
	nameVisitor := &NamePrinter{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}
	fmt.Printf("\nProduct list:\n-------------\n%s", nameVisitor.ProductList)
}
