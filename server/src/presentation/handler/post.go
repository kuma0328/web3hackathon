package handler

import (
	"fmt"
	"io"
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

func (h *PostHandler) GetDataByPostById(ctx *gin.Context) {
	idString := ctx.Param("id")
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

	postJson := postEntityToJson(post)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": postJson},
	)

}

func (h *PostHandler) GetImgByPostById(ctx *gin.Context) {
	idString := ctx.Param("id")
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
	ctx.Data(http.StatusOK, "image/jpeg", post.Img)

}

func (h *PostHandler) UpdatePostOfId(ctx *gin.Context) {
	img, _, err := ctx.Request.FormFile("img")
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	source, _ := io.ReadAll(img)

	newContent := ctx.Request.FormValue("content")
	newCommunityIdString := ctx.Request.FormValue("community_id")
	newCommunityId, err := strconv.Atoi(newCommunityIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"community_id error": err.Error()},
		)
		return
	}

	newUserIdString := ctx.Request.FormValue("user_id")
	newUserId, err := strconv.Atoi(newUserIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"user_id error": err.Error()},
		)
		return
	}

	newIdString := ctx.Request.FormValue("id")
	newId, err := strconv.Atoi(newIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	post := &entity.Post{}
	post.Id = newId
	post.Content = newContent
	post.CommunityId = newCommunityId
	post.UserId = newUserId
	post.Img = source

	newpost, err := h.uc.UpdatePostOfId(ctx, post)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", newpost.Img)

}
func (h *PostHandler) DeletePostOfId(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err = h.uc.DeletePostOfId(ctx, id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": fmt.Sprintf("success delete posts ( id : %d )", id)},
	)

}
func (h *PostHandler) CreateNewPost(ctx *gin.Context) {
	img, _, err := ctx.Request.FormFile("img")
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	source, _ := io.ReadAll(img)

	newContent := ctx.Request.FormValue("content")
	newCommunityIdString := ctx.Request.FormValue("community_id")
	newCommunityId, err := strconv.Atoi(newCommunityIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	newUserIdString := ctx.Request.FormValue("user_id")
	newUserId, err := strconv.Atoi(newUserIdString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	post := &entity.Post{}
	post.Content = newContent
	post.CommunityId = newCommunityId
	post.UserId = newUserId
	post.Img = source

	newpost, err := h.uc.CreateNewPost(ctx, post)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.Data(http.StatusOK, "image/jpeg", newpost.Img)
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
