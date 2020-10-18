package models

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/IrisAdminApi/sysinit"
	"time"
)

type Article struct {
	gorm.Model

	Title        string    `gorm:"not null;default:'';type:varchar(256)" json:"title" validate:"required,gte=4,lte=256" comment:"标题"`
	ContentShort string    `gorm:"not null;default:'';type:varchar(512)" json:"content_short" validate:"required,gte=6,lte=512" comment:"简介"`
	Author       string    `gorm:"not null;default:'';type:varchar(30)" json:"author" comment:"作者" validate:"required,gte=4,lte=30"`
	ImageUri     string    `gorm:"type:longText" json:"image_uri" comment:"封面" validate:"required"`
	SourceUri    string    `gorm:"not null;default:'';type:varchar(512)" json:"source_uri" comment:"来源"`
	IsOriginal   bool      `gorm:"not null;default:true;type:tinyint(1)" json:"is_original" comment:"是否原创" validate:""`
	Content      string    `gorm:"type:longText" json:"content" comment:"内容" validate:"required,gte=6"`
	Status       string    `gorm:"not null;default:'';type:varchar(10)" json:"status" comment:"文章状态" validate:"required,gte=1,lte=10"`
	DisplayTime  time.Time `json:"display_time" comment:"发布时间" validate:"required"`
}

func NewArticle() *Article {
	return &Article{}
}

/**
 * 通过 id 获取 role 记录
 * @method GetArticleById
 * @param  {[type]}       role  *Article [description]
 */
func (r *Article) GetPublishedArticleById(id uint) *Article {
	IsNotFound(sysinit.Db.Where("id = ?", id).Where("status = ?", "published").First(r).Error)
	return r
}

/**
 * 通过 id 获取 role 记录
 * @method GetArticleById
 * @param  {[type]}       role  *Article [description]
 */
func (r *Article) GetArticleById(id uint) *Article {
	IsNotFound(sysinit.Db.Where("id = ?", id).First(r).Error)
	return r
}

/**
 * 通过 id 删除角色
 * @method DeleteArticleById
 */
func (r *Article) DeleteArticleById() *Article {
	if err := sysinit.Db.Delete(r).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteArticleErr:%s \n", err))
	}
	return r
}

/**
 * 获取所有的角色
 * @method GetAllArticle
 * @param  {[type]} name string [description]
 * @param  {[type]} orderBy string [description]
 * @param  {[type]} offset int    [description]
 * @param  {[type]} limit int    [description]
 */
func GetAllArticles(name, orderBy, published string, offset, limit int) ([]*Article, int64, error) {
	var articles []*Article
	var count int64

	getAll := GetAll(&Article{}, name, orderBy)
	if err := getAll.Count(&count).Error; err != nil {
		return nil, count, err
	}
	if len(published) > 0 {
		getAll = getAll.Where("status = ?", "published")
	}

	if err := getAll.Scopes(Paginate(offset, limit)).Find(&articles).Error; err != nil {
		return nil, count, err
	}

	fmt.Println(fmt.Sprintf("offset:%d limit:%d", offset, limit))

	return articles, count, nil
}

/**
 * 创建
 * @method CreateArticle
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (r *Article) CreateArticle() error {
	if err := sysinit.Db.Create(r).Error; err != nil {
		return err
	}

	return nil
}

/**
 * 更新
 * @method UpdateArticle
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func (r *Article) UpdateArticle() error {
	if err := Update(&Article{}, r); err != nil {
		return err
	}

	return nil
}
