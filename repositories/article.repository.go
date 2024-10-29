package repositories

import (
	"belajar-go-fiber/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *models.Article) error
	FindAll(page, limit, offset int) ([]*models.Article, error)
	Update(article *models.Article) error
	FindByID(id string) (*models.Article, error)
	Count() (int, error)
	Delete(id string) error
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{DB: db}
}

func (ar *articleRepository) CreateArticle(article *models.Article) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(article).Error; err != nil {
			return err
		}
		return nil
	})
}

func (ar *articleRepository) FindAll(page, limit, offset int) ([]*models.Article, error) {
	var articles []*models.Article
	err := ar.DB.Offset(offset).Limit(limit).Find(&articles).Error
	return articles, err
}

func (ar *articleRepository) Update(article *models.Article) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(article).Error; err != nil {
			return err
		}
		return nil
	})
}

func (ar *articleRepository) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := ar.DB.First(&article, id).Error
	return &article, err
}

func (ar *articleRepository) Count() (int, error) {
	var count int64
	err := ar.DB.Model(&models.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ar *articleRepository) Delete(id string) error {
	return ar.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
			return err
		}
		return nil
	})
}
