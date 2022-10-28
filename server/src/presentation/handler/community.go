package handler

import (
	"net/http"

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
// GET /community/:id
func (h *CommunityHandler) GetCommunityById(ctx *gin.Context) {
	id := ctx.Param("id")

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
