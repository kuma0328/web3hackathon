package persistance

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.IRecipeRepository = &RecipeRepository{}

// TODO なんかこの依存関係いまいちな気がする
type RecipeRepository struct {
	conn *database.Conn
}

func NewRecipeRepository(conn *database.Conn) *RecipeRepository {
	return &RecipeRepository{
		conn: conn,
	}
}

// GetRecipeByCommunityIdは指定されたcommunity idのレシピを検索します
func (repo *RecipeRepository) GetRecipeByCommunityId(ctx context.Context, communityId int) (*entity.Recipe,error){
	var dto recipeDto

	query := `
	SELECT *
	FROM recipes
	WHERE id = ?
	LIMIT 1
	`
	err := repo.conn.DB.GetContext(ctx,&dto,query,communityId)
	if err != nil {
		return nil,fmt.Errorf("RecipeRepository.GetRecipeByCommunityId GetContext Error: %w",err)
	}
	return recipeDtoToEntity(&dto),nil
}

type recipeDto struct {
	Id          int   `db:"id"`
	Name        string `db:"name"`
	ImgUrl      string `db:"img_url"`
	CommunityId int  `db:"community_id"`
}

type recipesDto []*recipeDto

// entity to dto
func recipeEntityToDto(recipe *entity.Recipe) *recipeDto {
	return &recipeDto{
		Id: recipe.Id,
		Name: recipe.Name,
		ImgUrl: recipe.ImgUrl,
	}
}

// dto to entity
func recipeDtoToEntity(dto *recipeDto) *entity.Recipe {
	return &entity.Recipe{
		Id: dto.Id,
		Name: dto.Name,
		ImgUrl: dto.ImgUrl,
	}
}
