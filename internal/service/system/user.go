package system

import (
	"errors"

	"github.com/toma-photo/internal/models/system"
	"github.com/toma-photo/internal/pkg/rnd"
	"github.com/toma-photo/internal/utils"
	"gorm.io/gorm"
)

type UserService struct{}

// 用户登录
func (userService *UserService) UserLogin(u *system.User) (*system.User, error) {
	db := Db

	flag := false
	if u.Username != "" {
		db = db.Where("username = ? OR phone = ?", u.Username, u.Username)
	}

	if u.Phone != "" {
		flag = true
		db = db.Where("phone = ?", u.Phone)
	}

	user := new(system.User)
	if err := db.Select(`id,user_uid,username,password,nick_name,
	header_img,sex,phone,city,province,country,login_at,all_volume,used_volume,
	register_type,enable_ai_photo_face,enable_ai_photo_label`).
		First(user).Error; err != nil {
		return nil, err
	}

	if !flag {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("账号或密码错误")
		}
	}

	return user, nil
}

// 用户注册
func (userService *UserService) UserRegister(u system.User) (userInter system.User, err error) {
	var user system.User
	if !errors.Is(Db.Where("username = ? ", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已经注册")
	}
	if !errors.Is(Db.Where("phone = ? ", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("手机号已经注册")
	}

	u.Password = utils.BcryptHash(u.Password)
	u.UserUID = rnd.GenerateUID('u')
	err = Db.Create(&u).Error
	return u, err
}

// 用户注销
func (userService *UserService) UserLogout(userUID string) error {
	if userUID == "" {
		return errors.New("无效用户")
	}

	return Db.Delete(new(system.User), "user_uid = ?", userUID).Error
}

// 获取用户信息
func (userService *UserService) GetUserInfo(userUID string) (*system.User, error) {
	user := new(system.User)

	if err := Db.
		Select(`id,user_uid,username,password,nick_name,
	header_img,sex,phone,email,city,province,country,login_at,all_volume,used_volume,
	register_type,enable_ai_photo_face,enable_ai_photo_label`).
		First(user, "user_uid = ?", userUID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// 更新用户信息
func (userService *UserService) UpdateUserInfo(userID uint, userUpdate map[string]interface{}) error {
	if userID == 0 {
		return errors.New("更新对象为空")
	}
	return Db.Model(new(system.User)).Where("id = ?", userID).Updates(userUpdate).Error
}

// 通过旧密码更换密码
func (userService *UserService) ChangeUserPasswordByPwd(userUID, oldPwd, newPwd string) error {
	findUser := new(system.User)
	err := Db.Model(new(system.User)).
		Select("id,user_uid,password").
		Where("user_uid = ?", userUID).
		First(findUser).Error
	if err != nil {
		return err
	}

	if !utils.BcryptCheck(oldPwd, findUser.Password) {
		return errors.New("旧密码错误")
	}

	return userService.ChangeUserPassword(userUID, newPwd)
}

// 更新密码
func (userService *UserService) ChangeUserPassword(userUID, newPwd string) error {
	if userUID == "" {
		return errors.New("未指定用户")
	}
	if newPwd == "" {
		return errors.New("新密码不能为空")
	}

	pwdHash := utils.BcryptHash(newPwd)

	return Db.Model(new(system.User)).
		Where("user_uid = ?", userUID).
		Update("password", pwdHash).
		Error
}

// 更新密码
func (userService *UserService) ChangeUserPasswordByPhone(phone, newPwd string) error {
	if phone == "" {
		return errors.New("未指定用户")
	}
	if newPwd == "" {
		return errors.New("新密码不能为空")
	}

	pwdHash := utils.BcryptHash(newPwd)

	return Db.Model(new(system.User)).
		Where("phone = ?", phone).
		Update("password", pwdHash).
		Error
}

// 更新手机号
func (userService *UserService) ChangeUserPhone(userUID, phone string) error {
	if userUID == "" {
		return errors.New("未指定用户")
	}
	if phone == "" {
		return errors.New("手机号不能为空")
	}

	return Db.Model(new(system.User)).
		Where("user_uid = ?", userUID).
		Update("phone", phone).
		Error
}
