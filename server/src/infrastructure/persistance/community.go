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

func (repo *CommunityRepository) GetCommunityAll() (entity.Communities, error) {
	var dtos communityDtos

	query := `
	SELECT *
	FROM communities
	`
	err := repo.conn.DB.Select(&dtos, query)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.GetCommunityAll Get Error : %w", err)
	}
	return communityDtosToEntity(dtos), nil
}

func (repo *CommunityRepository) GetCommunityById(id int) (*entity.Community, error) {
	var dto communityDto

	query := `
	SELECT * 
	FROM communities
	WHERE id = ?
	LIMIT 1
	`
	err := repo.conn.DB.Get(&dto, query, id)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.GetCommunityById Get Error : %w", err)
	}
	return communityDtoToEntity(&dto), nil
}

func (repo *CommunityRepository) UpdateCommunityOfId(community *entity.Community) (*entity.Community, error) {
	dto := communityEntityToDto(community)

	query := `
	UPDATE communities
	SET name    = :name,
		img_url = :img_url,
		content = :content
	WHERE id = :id
	`
	_, err := repo.conn.DB.NamedExec(query, &dto)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.UpdateCommunityOfId NamedExec Error : %w", err)
	}
	return communityDtoToEntity(&dto), nil

}

func (repo *CommunityRepository) DeleteCommunityOfId(id int) error {
	query := `
	DELETE FROM communities
	WHERE id = :id
	`
	// TODO 存在しないidを削除しても成功とされてしまう
	_, err := repo.conn.DB.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("CommunityRepository.DeleteCommunityOfId NamedExec Error : %w", err)
	}
	return nil
}

func (repo *CommunityRepository) CreateNewCommunity(community *entity.Community) (*entity.Community, error) {
	dto := communityEntityToDto(community)

	query := `
	INSERT INTO communities (name, img_url, content)
	VALUES (:name,:img_url,:content)
	`
	res, err := repo.conn.DB.NamedExec(query, &dto)

	// TODO もっといい感じにバインドしたい
	id, _ := res.LastInsertId()
	dto.Id = (int)(id)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.CreateNewCommunity NamedExec Error : %w", err)
	}
	return communityDtoToEntity(&dto), nil
}

type communityDto struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	ImgUrl  string `db:"img_url"`
	Content string `db:"content"`
}

type communityDtos []*communityDto

func communityDtoToEntity(dto *communityDto) *entity.Community {
	return &entity.Community{
		Id:      dto.Id,
		Name:    dto.Name,
		ImgUrl:  dto.ImgUrl,
		Content: dto.Content,
	}
}

func communityDtosToEntity(dtos communityDtos) entity.Communities {
	var communities entity.Communities
	for _, d := range dtos {
		communities = append(communities, communityDtoToEntity(d))
	}
	return communities
}

func communityEntityToDto(c *entity.Community) communityDto {
	return communityDto{
		Id:      c.Id,
		Name:    c.Name,
		ImgUrl:  c.ImgUrl,
		Content: c.Content,
	}
}
