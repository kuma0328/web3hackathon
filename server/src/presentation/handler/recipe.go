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

// GET community/:community_id/recipe
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

// POST community/:community_id/recipe
func (h *RecipeHandler) CreateNewRecipeOfCommunity(ctx *gin.Context){
	communityIdString := ctx.Param("community_id")
	communityId,err:=strconv.Atoi(communityIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	var j recipeJson
	if err := ctx.BindJSON(&j); err !=nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	recipe,err := h.uc.CreateNewRecipeOfCommunity(ctx.Request.Context(),recipeJsonToEntity(&j),communityId)
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



// entity to json
func recipeEntityToJson(recipe *entity.Recipe) *recipeJson {
	return &recipeJson{
		Id: recipe.Id,
		Name: recipe.Name,
		ImgUrl: recipe.ImgUrl,
		RecipeSteps: recipeStepsEntityToJson(recipe.RecipeSteps),
		Spices: spicesEntityToJson(recipe.Spices),
	}
}

func spicesEntityToJson(spices entity.Spices) spicesJson {
	var s spicesJson
	for _,spice := range spices {
		s = append(s, spiceEntityToJson(spice))
	}
	return s
}




// json to entity
func recipeJsonToEntity(r *recipeJson) *entity.Recipe {
	return &entity.Recipe{
		Id: r.Id,
		Name: r.Name,
		ImgUrl: r.ImgUrl,
	}
}

