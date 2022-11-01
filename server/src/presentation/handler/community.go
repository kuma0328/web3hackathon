package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/usecase"
)

type CommunityHandler struct {
	uc usecase.ICommunityUsecase
}

func NewCommunityHandler(uc usecase.ICommunityUsecase) *CommunityHandler {
	return &CommunityHandler{
		uc: uc,
	}
}

// GetCommunityByIdはIdを指定してcommuntyを取得するハンドラーです
// GET /community/:community_id
func (h *CommunityHandler) GetCommunityById(ctx *gin.Context) {
	idString := ctx.Param("community_id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	community, err := h.uc.GetCommunityById(id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	communityJson := communityEntityToJson(community)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": communityJson},
	)
}

// UpdateCommunityOfIdは指定したidのcommunity情報を更新するハンドラーです
// PUT /community/:community_id
func (h *CommunityHandler) UpdateCommunityOfId(ctx *gin.Context){
	idString := ctx.Param("community_id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	var json communityJson
	if err := ctx.BindJSON(&json);err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	community,err:= h.uc.UpdateCommunityOfId(id, communityJsonToEntity(&json))
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}
	communityJson := communityEntityToJson(community)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data":communityJson},
	)
}

// DeleteCommunityOfIdは指定したidのcommunityを削除するハンドラーです
// DELETE /community/:community_id
func (h *CommunityHandler) DeleteCommunityOfId(ctx *gin.Context){
	idString := ctx.Param("community_id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}

	err = h.uc.DeleteCommunityOfId(id)
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"ok":fmt.Sprintf("success delete community ( id : %d )",id)},
	)
}


// CreateNewCommunityは新しいCommunityのデータを受け取って生成するハンドラーです
// POST /community
func (h *CommunityHandler) CreateNewCommunity(ctx *gin.Context){
	var json communityJson
	if err:=ctx.BindJSON(&json);err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}
	community,err := h.uc.CreateNewCommunity(communityJsonToEntity(&json))
	if err!=nil{
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error":err.Error()},
		)
		return
	}
	communityJson := communityEntityToJson(community)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data":communityJson},
	)
}

type communityJson struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	ImgUrl  string `json:"img_url"`
	Content string `json:"content"`
	Recipe  recipesJson
	User    usersJson
}

func communityEntityToJson(c *entity.Community) communityJson {
	return communityJson{
		Id:      c.Id,
		Name:    c.Name,
		ImgUrl:  c.ImgUrl,
		Content: c.Content,
		// TODO toJson関数作ってから直す
		Recipe: nil,
		User:   nil,
	}
}

func communityJsonToEntity(j *communityJson) *entity.Community {
	return &entity.Community{
		Id:      j.Id,
		Name:    j.Name,
		ImgUrl:  j.ImgUrl,
		Content: j.Content,
	}
}
