package persistance

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.IRecipeStepRepository = &RecipeStepRepository{}

type RecipeStepRepository struct {
	conn *database.Conn
}

func NewRecipeStepRepository(conn *database.Conn) repository.IRecipeStepRepository {
	return &RecipeStepRepository{
		conn: conn,
	}
}

func (repo *RecipeStepRepository) GetRecipeStepsByRecipeId(ctx context.Context, recipeId int) (entity.RecipeSteps, error){
	var dtos recipeStepsDto

	query := `
	SELECT *
	FROM recipe_steps
	WHERE recipe_id = ?
	`
	err := repo.conn.DB.SelectContext(ctx,&dtos,query,recipeId)
	if err != nil {
		return nil,fmt.Errorf("RecipeStepRepository.GetRecipeStepsByRecipeId SelectContext Error: %w",err)
	}
	return recipeStepsDtoToEntity(dtos),nil
}








type recipeStepDto struct {
	Id      int `db:"id"`
	Number  int `db:"number"`
	Content string `db:"content"`
	RecipeId int  `db:"recipe_id"`
}

type recipeStepsDto []*recipeStepDto

// entity to dto
func recipeStepsEntityToDto(recipeSteps entity.RecipeSteps) recipeStepsDto {
	var r recipeStepsDto
	for _,step := range recipeSteps {
		r = append(r, recipeStepEntityToDto(step))
	}
	return r
}

func recipeStepEntityToDto(recipeStep *entity.RecipeStep) *recipeStepDto {
	return &recipeStepDto{
		Id: recipeStep.Id,
		Number: recipeStep.Id,
		Content: recipeStep.Content,
	}
}



// dto to entity
func recipeStepsDtoToEntity(dtos recipeStepsDto) entity.RecipeSteps {
	var r entity.RecipeSteps
	for _,dto := range dtos {
		r = append(r, recipeStepDtoToEntity(dto))
	}
	return r
}

func recipeStepDtoToEntity(dto *recipeStepDto) *entity.RecipeStep {
	return &entity.RecipeStep{
		Id: dto.Id,
		Number:dto.Id,
		Content: dto.Content,
	}
}