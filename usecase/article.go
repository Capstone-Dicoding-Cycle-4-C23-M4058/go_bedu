package usecase

import (
	"go_bedu/dtos"
	"go_bedu/helpers"
	"go_bedu/models"
	"go_bedu/repositories"
)

type ArticleUsecase interface {
	GetAllArticles(page, limit int) ([]dtos.ArticleDetailResponse, int, error)
	GetArticleByID(id uint) (dtos.ArticleDetailResponse, error)
	CreateArticle(article *dtos.CreateArticlesRequest) (dtos.ArticleDetailResponse, error)
	UpdateArticle(id uint, article dtos.UpdateArticlesRequest) (dtos.ArticleDetailResponse, error)
	DeleteArticle(id uint) error
}

type articleUsecase struct {
	articleRepository repositories.ArticleRepository
}

func NewArticleUsecase(ArticleRepository repositories.ArticleRepository) ArticleUsecase {
	return &articleUsecase{ArticleRepository}
}

// GetAllArticles godoc
// @Summary      Get all articles
// @Description  Get all articles
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success      200 {object} dtos.GetAllArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [get]
// @Security BearerAuth
func (u *articleUsecase) GetAllArticles(page, limit int) ([]dtos.ArticleDetailResponse, int, error) {
	articles, count, err := u.articleRepository.GetAllArticles(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var articleResponses []dtos.ArticleDetailResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, dtos.ArticleDetailResponse{
			ArticleID:   article.ID,
			Title:       article.Title,
			Abstract:    article.Abstract,
			Image:       article.Image,
			Description: article.Description,
			Label:       article.Label,
			Slug:        article.Slug,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
		})
	}

	return articleResponses, count, nil
}

// GetArticleByID godoc
// @Summary      Get article by ID
// @Description  Get article by ID
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Success      200 {object} dtos.ArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article/{id} [get]
// @Security BearerAuth
func (u *articleUsecase) GetArticleByID(id uint) (dtos.ArticleDetailResponse, error) {
	var articleResponses dtos.ArticleDetailResponse

	article, err := u.articleRepository.GetArticleByID(id)
	if err != nil {
		return articleResponses, err
	}

	articleResponse := dtos.ArticleDetailResponse{
		ArticleID:   article.ID,
		Title:       article.Title,
		Image:       article.Image,
		Description: article.Description,
		Label:       article.Label,
		Slug:        article.Slug,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}

	return articleResponse, nil
}

// CreateArticle godoc
// @Summary      Create a new article
// @Description  Create a new article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param        request body dtos.CreateArticlesRequest true "Payload Body [RAW]"
// @Success      201 {object} dtos.ArticleCreeatedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [post]
// @Security BearerAuth
func (u *articleUsecase) CreateArticle(article *dtos.CreateArticlesRequest) (dtos.ArticleDetailResponse, error) {
	var articleResponses dtos.ArticleDetailResponse

	slug := helpers.CreateSlug(article.Title)

	CreateArticle := models.Article{
		AdministratorID: article.AdministratorID,
		Thumbnail:       article.Thumbnail,
		Title:           article.Title,
		Description:     article.Description,
		Image:           article.Image,
		Label:           article.Label,
		Slug:            slug,
		Abstract:        article.Abstract,
	}

	createdArticle, err := u.articleRepository.CreateArticle(CreateArticle)
	if err != nil {
		return articleResponses, err
	}

	articleResponse := dtos.ArticleDetailResponse{
		ArticleID:   createdArticle.ID,
		Thumbnail:   createdArticle.Thumbnail,
		Title:       createdArticle.Title,
		Abstract:    createdArticle.Abstract,
		Image:       createdArticle.Image,
		Description: createdArticle.Description,
		Label:       createdArticle.Label,
		Slug:        createdArticle.Slug,
		CreatedAt:   createdArticle.CreatedAt,
		UpdatedAt:   createdArticle.UpdatedAt,
	}

	return articleResponse, nil
}

// UpdateArticle godoc
// @Summary      Update article
// @Description  Update article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Param        request body dtos.CreateArticlesRequest true "Payload Body [RAW]"
// @Success      200 {object} dtos.ArticleStatusOKResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article [put]
// @Security BearerAuth
func (u *articleUsecase) UpdateArticle(id uint, article dtos.UpdateArticlesRequest) (dtos.ArticleDetailResponse, error) {
	var (
		articles        models.Article
		articleResponse dtos.ArticleDetailResponse
	)

	articles, err := u.articleRepository.GetArticleByID(id)
	if err != nil {
		return articleResponse, err
	}

	slug := helpers.CreateSlug(articles.Title)

	articles.Title = article.Title
	articles.Description = article.Description
	articles.Image = article.Image
	articles.Label = article.Label
	articles.Slug = slug

	articles, err = u.articleRepository.UpdateArticle(articles)
	if err != nil {
		return articleResponse, err
	}

	articleResponse.ArticleID = articles.ID
	articleResponse.Title = articles.Title
	articleResponse.Image = articles.Image
	articleResponse.Description = articles.Description
	articleResponse.Label = articles.Label
	articleResponse.Slug = articles.Slug
	articleResponse.CreatedAt = articles.CreatedAt
	articleResponse.UpdatedAt = articles.UpdatedAt

	return articleResponse, nil

}

// DeleteArticle godoc
// @Summary      Delete a article
// @Description  Delete a article
// @Tags         Admin - Article
// @Accept       json
// @Produce      json
// @Param id path integer true "ID article"
// @Success      200 {object} dtos.StatusOKDeletedResponse
// @Failure      400 {object} dtos.BadRequestResponse
// @Failure      401 {object} dtos.UnauthorizedResponse
// @Failure      403 {object} dtos.ForbiddenResponse
// @Failure      404 {object} dtos.NotFoundResponse
// @Failure      500 {object} dtos.InternalServerErrorResponse
// @Router       /admin/article/{id} [delete]
// @Security BearerAuth
func (u *articleUsecase) DeleteArticle(id uint) error {
	article, err := u.articleRepository.GetArticleByID(id)
	if err != nil {
		return err
	}

	err = u.articleRepository.DeleteArticle(article)
	return err
}
