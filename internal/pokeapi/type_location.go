package pokeapi

type PokedexLocation struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
	Count int
}
