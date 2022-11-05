package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/usecase"
)

type SpiceHandler struct {
	uc usecase.ISpiceUsecase
}

func NewSpiceHandler(uc usecase.ISpiceUsecase) *SpiceHandler {
	return &SpiceHandler{
		uc: uc,
	}
}

// POST /community/:community_id/recipe/spice
func (h *SpiceHandler) PostNewSpice(ctx *gin.Context) {
	communityIdString := ctx.Param("community_id")
	communityId,err:=strconv.Atoi(communityIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	var j spiceJson
	if err := ctx.BindJSON(&j); err !=nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	spice, err := h.uc.PostNewSpice(ctx.Request.Context(),spiceJsonToEntity(&j),communityId)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	spiceJson := spiceEntityToJson(spice)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data":spiceJson},
	)
}


type spiceJson struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type spicesJson []*spiceJson

// entity to json
func spiceEntityToJson(spice *entity.Spice) *spiceJson {
	return &spiceJson{
		Id: spice.Id,
		Name: spice.Name,
	}
}


// json to entity
func spiceJsonToEntity(s *spiceJson) *entity.Spice {
	return &entity.Spice{
		Id: s.Id,
		Name: s.Name,
	}
}