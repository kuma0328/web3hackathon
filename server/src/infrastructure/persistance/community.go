package persistance

import (
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.ICommunityRepository = &CommunityRepository{}

type CommunityRepository struct {
	conn *database.Conn
}

func NewCommunityRepository(conn *database.Conn) repository.ICommunityRepository {
	return &CommunityRepository{
		conn: conn,
	}
}

func (repo *CommunityRepository) GetCommunityById(id string) (*entity.Community, error) {
	var dto communityDto

	query := `
	SELECT * 
	FROM communities
	WHERE id=?
	LIMIT 1
	`
	err := repo.conn.DB.Get(&dto, query, id)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.GetCommunityById Select Error : %w", err)
	}
	return communityDtotoToEntity(&dto), nil
}

type communityDto struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	ImgUrl  string `db:"img_url"`
	Content string `db:"content"`
}

type communityDtos []*communityDto

func communityDtotoToEntity(dto *communityDto) *entity.Community {
	return &entity.Community{
		Id:      dto.Id,
		Name:    dto.Name,
		ImgUrl:  dto.ImgUrl,
		Content: dto.Content,
	}
}

func communityDtostoToEntity(dtos communityDtos) entity.Communities {
	var communities entity.Communities
	for _, d := range dtos {
		communities = append(communities, communityDtotoToEntity(d))
	}
	return communities
}
