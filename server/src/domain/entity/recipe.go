package entitiy

// Recipeはレシピに関する構造体です
type Recipe struct {
	Id        int
	Name      string
	ImgUrl    string
	Community Community
}

// RecipeStepはレシピの手順に関する構造体です
type RecipeStep struct {
	Id      int
	Number  int
	Content string
	Recipe  Recipe
}

// Spicesはレシピにのっている調味料に関する構造体です
type Spices struct {
	Id     int
	Name   string
	Recipe Recipe
}
