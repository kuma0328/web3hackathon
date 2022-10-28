package entitiy

// Recipeはレシピに関する構造体です
type Recipe struct {
	Id        int
	Name      string
	ImgUrl    string
	RecipeSteps RecipeSteps
	Spices Spices
}

// RecipeStepはレシピの手順に関する構造体です
type RecipeStep struct {
	Id      int
	Number  int
	Content string
}

type RecipeSteps []RecipeStep

// Spiceはレシピにのっている調味料に関する構造体です
type Spice struct {
	Id     int
	Name   string
}

type Spices []Spice