package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/usecase"
)

type RecipeStepHandler struct {
	uc usecase.IRecipeStepUsecase
}

func NewRecipeStepHandler(uc usecase.IRecipeStepUsecase) *RecipeStepHandler {
	return &RecipeStepHandler{
		uc:uc,
	}
}

// POST /community/:community_id/recipe/step
func (h *RecipeStepHandler) PostNewRecipeStep(ctx *gin.Context) {
	communityIdString := ctx.Param("community_id")
	communityId,err:=strconv.Atoi(communityIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	var j recipeStepJson
	if err := ctx.BindJSON(&j); err !=nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	step,err := h.uc.PostNewRecipeStep(ctx.Request.Context(),recipeStepJsonToEntity(&j),communityId)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	stepJson := recipeStepEntityToJson(step)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data":stepJson},
	)
}

type recipeStepJson struct {
	Id      int `json:"id"`
	Number  int `json:"number"`
	Content string `json:"content"`
}

type recipeStepsJson []*recipeStepJson

// entity to json
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
		Number: recipeStep.Number,
		Content: recipeStep.Content,
	}
}


// json to entity
func recipeStepJsonToEntity(r *recipeStepJson) *entity.RecipeStep {
	return &entity.RecipeStep{
		Id: r.Id,
		Number: r.Number,
		Content: r.Content,
	}
}

