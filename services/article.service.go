package services

import (
	"belajar-go-fiber/config"
	"belajar-go-fiber/models"
	"belajar-go-fiber/repositories"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type ArticleService interface {
	CreateArticle(article *models.Article) error
	FindAll(page, limit int) ([]*models.Article, error)
	UpdateArticle(article *models.Article) error
	FindByID(id string) (*models.Article, error)
	Count() (int, error)
	DeleteArticle(id string) error
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (a *articleService) CreateArticle(article *models.Article) error {
	if article == nil {
		return errors.New("article is required")
	}
	return a.repo.CreateArticle(article)
}

func (a *articleService) FindAll(page, limit int) ([]*models.Article, error) {
	cacheKey := "articles:page:" + strconv.Itoa(page) + ":limit:" + strconv.Itoa(limit)

	val, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == redis.Nil {
		// Data tidak ada di cache, ambil dari PostgreSQL
		offset := (page - 1) * limit
		articles, err := a.repo.FindAll(page, limit, offset) // Pass pagination parameters
		if err != nil {
			return nil, err
		}
		// Simpan hasil ke cache
		jsonData, marshalErr := json.Marshal(articles)
		if marshalErr != nil {
			return nil, marshalErr
		}
		config.RedisClient.Set(config.Ctx, cacheKey, jsonData, 0) // 0 berarti tidak ada expiry

		return articles, nil
	}

	// Data ditemukan di cache, unmarshal JSON
	var articles []*models.Article
	unmarshalErr := json.Unmarshal([]byte(val), &articles)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return articles, nil

}

func (a *articleService) UpdateArticle(article *models.Article) error {
	return a.repo.Update(article)
}

func (a *articleService) FindByID(id string) (*models.Article, error) {
	article, err := a.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *articleService) Count() (int, error) {
	return a.repo.Count()
}

func (a *articleService) DeleteArticle(id string) error {
	return a.repo.Delete(id)
}
