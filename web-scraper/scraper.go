package main

import (
	"github.com/gocolly/colly"
)

// defining a data structure to store the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {

	c := colly.NewCollector()
	// initializing the slice of structs that will contain the scraped data
	var pokemonProducts []PokemonProduct

	c.Visit("https://scrapeme.live/shop/")

	// iterating over the list of HTML product elements
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		pokemonProduct := PokemonProduct{}

		// scraping the data of interest
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		// adding the product instance with scraped data to the list of products
		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

}
