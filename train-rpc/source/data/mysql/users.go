package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/pkg/db"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"gorm.io/gorm"
)

type UsersConfModel struct{}

func NewUsersConfModel() biz.UsersConfRepo {
	return &UsersConfModel{}
}

var (
	userDB *gorm.DB
)

func (u *UsersConfModel) CreateUser() error {
	userDB = db.GetDB()
	//userDB.Create(biz.UsersConf{})
	if userDB == nil {
		return errors.New("the users table init failed")
	}
	return nil
}

func (u *UsersConfModel) AddUser(ctx context.Context, user *biz.UsersConf) (int64, error) {
	result := userDB.WithContext(ctx).Create(user)
	if result.Error != nil {
		return 0, errors.New("add to users failed")
	}
	return user.UserId, nil
}

func (u *UsersConfModel) DelUser(ctx context.Context, userId int64) error {
	result := userDB.WithContext(ctx).Where("user_id =?", userId).Delete(biz.UsersConf{})
	if result.Error != nil {
		return errors.New("delete user record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no train record found for the given user_id")
	}
	return nil
}

func (u *UsersConfModel) GetUserDetail(ctx context.Context, getDetailParam *biz.GetUserDetailParam) (*biz.UsersConf, error) {
	user := biz.UsersConf{}
	query := userDB.WithContext(ctx)

	if getDetailParam.UserId != 0 {
		query = query.Where("user_id =?", getDetailParam.UserId)
	}
	if getDetailParam.UserName != "" {
		query = query.Where("user_name =?", getDetailParam.UserName)
	}
	if getDetailParam.Email != "" {
		query = query.Where("email =?", getDetailParam.Email)
	}
	if getDetailParam.PhoneNumber != "" {
		query = query.Where("phone_number =?", getDetailParam.PhoneNumber)
	}

	result := query.First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user record not found")
		}
		return nil, errors.New("failed to get user detail")
	}
	return &user, nil
}

func (u *UsersConfModel) UpdateUser(ctx context.Context, user *biz.UsersConf) error {
	result := userDB.WithContext(ctx).Model(biz.UsersConf{}).Where("user_id =?", user.UserId).Updates(user)
	if result.Error != nil {
		return errors.New("update user record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no user record found for the given user_id")
	}
	return nil
}
