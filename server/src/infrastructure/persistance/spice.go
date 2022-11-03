package persistance

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.ISpiceRepository = &SpiceRepository{}

type SpiceRepository struct {
	conn *database.Conn
}

func NewSpiceRepository(conn *database.Conn) repository.ISpiceRepository {
	return &SpiceRepository{
		conn: conn,
	}
}

func (repo *SpiceRepository) GetSpicesByRecipeId(ctx context.Context, recipeId int) (entity.Spices,error){
	var dtos spicesDto
	
	query := `
	SELECT *
	FROM spices
	WHERE recipe_id = ?
	`
	err := repo.conn.DB.SelectContext(ctx,&dtos,query,recipeId)
	if err != nil {
		return nil,fmt.Errorf("SpiceRepository.GetSpicesByRecipeId SelectContext Error: %w",err)
	}
	return spicesDtoToEntity(dtos),nil
}

func (repo *SpiceRepository) PostNewSpice(ctx context.Context,spice *entity.Spice ,recipeId int) (*entity.Spice,error) {
	dto := spiceEntityToDto(spice)
	dto.RecipeId = recipeId

	query := `
	INSERT INTO spices (name, recipe_id)
	VALUES (:name, :recipe_id)
	`
	res,err := repo.conn.DB.NamedExecContext(ctx,query,dto)
	if err!= nil{
		return nil, fmt.Errorf("SpiceRepository.PostNewSpice NamedExecContext Error : %w", err)
	}

	id,err := res.LastInsertId()
	if err != nil{
		return nil, fmt.Errorf("SpiceRepository.PostNewSpice LastInsertId Error : %w", err)
	}
	dto.Id = (int)(id)

	return spiceDtoToEntity(dto),nil
}

type spiceDto struct {
	Id   int `db:"id"`
	Name string `db:"name"`
	RecipeId int  `db:"recipe_id"`
}

type spicesDto []*spiceDto

// entity to dto
func spicesEntityToDto(spices entity.Spices) spicesDto {
	var s spicesDto
	for _,spice := range spices {
		s = append(s, spiceEntityToDto(spice))
	}
	return s
}

func spiceEntityToDto(spice *entity.Spice) *spiceDto {
	return &spiceDto{
		Id: spice.Id,
		Name: spice.Name,
	}
}


// dto to entity
func spicesDtoToEntity(dtos spicesDto) entity.Spices {
	var s entity.Spices
	for _,dto:= range dtos {
		s = append(s, spiceDtoToEntity(dto))
	}
	return s
}

func spiceDtoToEntity(dto *spiceDto)  *entity.Spice {
	return &entity.Spice{
		Id: dto.Id,
		Name: dto.Name,
	}
}