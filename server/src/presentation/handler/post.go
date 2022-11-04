package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/usecase"
)

type PostHandler struct {
	uc usecase.IPostUsecase
}

func NewPostHandler(uc usecase.IPostUsecase) *PostHandler {
	return &PostHandler{
		uc: uc,
	}
}

func (h *PostHandler) GetPostById(ctx *gin.Context) {
	idString := ctx.Param("community_id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	post, err := h.uc.GetPostById(ctx, id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

}
func (h *PostHandler) UpdatePostOfId(ctx *gin.Context) {

}
func (h *PostHandler) DeletePostOfId(ctx *gin.Context) {

}
func (h *PostHandler) CreateNewPost(ctx *gin.Context) {

}

type postJson struct {
	Id          int    `json:"id"`
	CommunityId int    `json:"community_id"`
	Img         []byte `json:"img"`
	Content     string `json:"content"`
	UserId      int    `json:"user_id"`
	User        usersJson
}

func postEntityToJson(c *entity.Post) postJson {
	return postJson{
		Id:          c.Id,
		CommunityId: c.CommunityId,
		Content:     c.Content,
		UserId:      c.UserId,
	}
}

func postJsonToEntity(j *postJson) *entity.Post {
	return &entity.Post{
		Id:          j.Id,
		CommunityId: j.CommunityId,
		Content:     j.Content,
		UserId:      j.UserId,
	}
}
