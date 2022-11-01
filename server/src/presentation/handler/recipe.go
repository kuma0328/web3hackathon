package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/usecase"
)

type RecipeHandler struct {
	uc usecase.IRecipeUsecase
}

func NewRecipeHandler(uc usecase.IRecipeUsecase) *RecipeHandler {
	return &RecipeHandler{
		uc:uc,
	}
}

// GET /recipe/:community_id
func (h *RecipeHandler) GetRecipeByCommunityId(ctx *gin.Context) {
	communityIdString := ctx.Param("community_id")
	communityId,err:=strconv.Atoi(communityIdString)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	recipe,err := h.uc.GetRecipeByCommunityId(ctx.Request.Context(),communityId)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	recipeJson := recipeEntityToJson(recipe)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data":recipeJson},
	)
}





type recipeJson struct {
	Id          int   `json:"id"`
	Name        string `json:"name"`
	ImgUrl      string `json:"img_url"`
	RecipeSteps recipeStepsJson `json:"recipe_steps"`
	Spices      spicesJson `json:"spices"`
}

type recipesJson []*recipeJson

type recipeStepJson struct {
	Id      int `json:"id"`
	Number  int `json:"number"`
	Content string `json:"content"`
}

type recipeStepsJson []*recipeStepJson

type spiceJson struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type spicesJson []*spiceJson

func recipeEntityToJson(recipe *entity.Recipe) *recipeJson {
	return &recipeJson{
		Id: recipe.Id,
		Name: recipe.Name,
		ImgUrl: recipe.ImgUrl,
		RecipeSteps: recipeStepsEntityToJson(recipe.RecipeSteps),
		Spices: spicesEntityToJson(recipe.Spices),
	}
}

func recipeStepsEntityToJson(recipeSteps entity.RecipeSteps) recipeStepsJson {
	var r recipeStepsJson
	for _,step := range recipeSteps {
		r = append(r, recipeStepEntityToJson(step))
	}
	return r
}

func recipeStepEntityToJson(recipeStep *entity.RecipeStep) *recipeStepJson {
	return &recipeStepJson{
		Id: recipeStep.Id,
		Number: recipeStep.Id,
		Content: recipeStep.Content,
	}
}

func spicesEntityToJson(spices entity.Spices) spicesJson {
	var s spicesJson
	for _,spice := range spices {
		s = append(s, spiceEntityToJson(spice))
	}
	return s
}


func spiceEntityToJson(spice *entity.Spice) *spiceJson {
	return &spiceJson{
		Id: spice.Id,
		Name: spice.Name,
	}
}