package link

import (
	"olymp/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	res := repo.Database.DB.Create(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link

	res := repo.Database.DB.First(&link, "hash = ?", hash)

	if res.Error != nil {
		return nil, res.Error
	}

	return &link, nil
}

func (repo *LinkRepository) isInBase(id uint) error {
	var link Link
	res := repo.Database.DB.First(&link, id)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	res := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if err := res.Error; err != nil {
		return nil, err
	}
	return link, nil
}
func (repo *LinkRepository) Delete(id uint) error {
	res := repo.Database.DB.Delete(&Link{}, id)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
