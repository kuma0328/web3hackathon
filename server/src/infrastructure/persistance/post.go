package persistance

import (
	"context"
	"fmt"
	"log"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.IPostRepository = &PostRepository{}

type PostRepository struct {
	conn *database.Conn
}

func NewPostRepository(conn *database.Conn) repository.IPostRepository {
	return &PostRepository{
		conn: conn,
	}
}

func (repo *PostRepository) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	var dto postDto

	query := `
	SELECT * 
	FROM posts
	WHERE id = ?
	LIMIT 1
	`
	err := repo.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.GetPostById Get Error : %w", err)
	}
	return postDtoToEntity(&dto), nil
}

func (repo *PostRepository) UpdatePostOfId(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	dto := postEntityToDto(post)

	query := `
	UPDATE posts
	SET img		= :img,
		content = :content
	WHERE id 	= :id
	`
	_, err := repo.conn.DB.NamedExecContext(ctx, query, &dto)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.UpdatePostOfId NamedExec Error : %w", err)
	}
	return postDtoToEntity(&dto), nil
}
func (repo *PostRepository) DeletePostOfId(ctx context.Context, id int) error {
	query := `
	DELETE FROM posts
	WHERE id = :id
	`

	_, err := repo.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("PostRepository.DeletePostOfId NamedExec Error : %w", err)
	}
	return nil
}
func (repo *PostRepository) CreateNewPost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	log.Println("1")
	dto := postEntityToDto(post)
	log.Println("2")

	query := `
	INSERT INTO posts (community_id, content, img, user_id)
	VALUES (:community_id,:content,:img,:user_id)
	`
	log.Println("3")

	res, err := repo.conn.DB.NamedExecContext(ctx, query, &dto)
	log.Println("4")
	if err != nil {
		return nil, fmt.Errorf("PostRepository.CreateNewPost NamedExec Error : %w", err)
	}

	id, err := res.LastInsertId()
	dto.Id = (int)(id)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.CreateNewPost NamedExec Error : %w", err)
	}

	return postDtoToEntity(&dto), nil
}

type postDto struct {
	Id          int    `db:"id"`
	CommunityId int    `db:"community_id"`
	Content     string `db:"content"`
	Img         []byte `db:"img"`
	Like        int    `db:"like"`
	UserId      int    `db:"user_id"`
}

type postDtos []*postDto

func postDtoToEntity(dto *postDto) *entity.Post {
	return &entity.Post{
		Id:          dto.Id,
		Content:     dto.Content,
		CommunityId: dto.CommunityId,
		Img:         dto.Img,
		UserId:      dto.UserId,
	}
}

func postDtosToEntity(dtos postDtos) entity.Posts {
	var posts entity.Posts
	for _, d := range dtos {
		posts = append(posts, postDtoToEntity(d))
	}
	return posts
}

func postEntityToDto(p *entity.Post) postDto {
	return postDto{
		Id:          p.Id,
		Content:     p.Content,
		CommunityId: p.CommunityId,
		Img:         p.Img,
		UserId:      p.UserId,
	}
}
