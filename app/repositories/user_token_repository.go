package repositories

import (
	"blog/app/models"
	"blog/database"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

func NewUserTokenRepository() *UserTokenRepository {
	return &UserTokenRepository{database.DB()}
}

type UserTokenRepository struct {
	db *gorm.DB
}

func (this *UserTokenRepository) GetByToken(token string) *models.UserToken {
	if len(token) == 0 {
		return nil
	}
	return this.Take(this.db, "token = ?", token)
}

func (this *UserTokenRepository) Get(id int64) *models.UserToken {
	ret := &models.UserToken{}
	if err := this.db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *UserTokenRepository) Take(where ...interface{}) *models.UserToken {
	ret := &models.UserToken{}
	if err := this.db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *UserTokenRepository) Find(cnd *simple.SqlCnd) (list []models.UserToken) {
	cnd.Find(this.db, &list)
	return
}

func (this *UserTokenRepository) FindOne(cnd *simple.SqlCnd) (ret *models.UserToken) {
	cnd.FindOne(this.db, &ret)
	return
}

func (this *UserTokenRepository) FindPageByCnd(cnd *simple.SqlCnd) (list []models.UserToken, paging *simple.Paging) {
	cnd.Find(this.db, &list)
	count := cnd.Count(this.db, &models.UserToken{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (this *UserTokenRepository) Create(t *models.UserToken) (err error) {
	return this.db.Create(t).Error
}

func (this *UserTokenRepository) Update(t *models.UserToken) (err error) {
	err = this.db.Save(t).Error
	return
}

func (this *UserTokenRepository) Updates(id int64, columns map[string]interface{}) (err error) {
	err = this.db.Model(&models.UserToken{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *UserTokenRepository) UpdateColumn(id uint, name string, value interface{}) (err error) {
	err = this.db.Model(&models.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *UserTokenRepository) Delete(id int64) {
	this.db.Delete(&models.UserToken{}, "id = ?", id)
}
