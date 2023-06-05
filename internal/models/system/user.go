package system

import (
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/pkg/format"
	"github.com/toma-photo/internal/pkg/rnd"
	"gorm.io/gorm"
)

const (
	DefaultUserVolume int64 = 500 * 1024 // 用户默认容量, 单位 B
)

type User struct {
	global.Model
	UserUID            string `json:"userUID" gorm:"column:user_uid; type:varchar(64); unique; index; comment:'用户UID'"`
	Username           string `json:"username" gorm:"column:username; type:varchar(128); unique; index; comment:'用户账号'"`
	Password           string `json:"password" gorm:"column:password; type:varchar(128); comment:'密码'"`
	NickName           string `json:"nickName" gorm:"column:nick_name; type:varchar(128); comment:'昵称'"`
	HeaderImg          string `json:"headerImg" gorm:"column:header_img; type:varchar(256); comment:'头像'"`
	Sex                int8   `json:"sex" gorm:"column:sex; type:tinyint(4); default:1; comment:'性别, 1表示男 2表示女, 默认1'"`
	Phone              string `json:"phone" gorm:"column:phone; type:varchar(64); comment:'手机号'"`
	Email              string `json:"email" gorm:"column:email; type:varchar(128); comment:'邮箱'"`
	City               string `json:"city" gorm:"column:city; type:varchar(128); comment:'城市'"`
	Province           string `json:"province" gorm:"column:province; type:varchar(128); comment:'省'"`
	Country            string `json:"country" gorm:"column:country; type:varchar(128); comment:'国家'"`
	LoginAt            int64  `json:"loginAt" gorm:"column:login_at; comment:'登录时间, 时间戳'"`
	AllVolume          int64  `json:"allVolume" gorm:"column:all_volume; comment:'总限额, 单位B'"`
	UsedVolume         int64  `json:"usedVolume" gorm:"column:used_volume; comment:'用户已经使用的限额, 单位B'"`
	RegisterType       string `json:"registerType" gorm:"column:register_type; default:web; type:VARCHAR(32); comment:'用户注册类型,小程序,app,pc端,web端'"`
	EnableAiPhotoFace  bool   `json:"enableAiPhotoFace" gorm:"column:enable_ai_photo_face; default:false; comment:'启用人脸分类, 1启用, 0禁用'"`
	EnableAiPhotoLabel bool   `json:"enableAiPhotoLabel" gorm:"column:enable_ai_photo_label; default:false; comment:'启用标签分类, 1启用, 0禁用'"`

	DisplayAllVolume  string  `json:"displayAllVolume" gorm:"-"`
	DisplayUsedVolume string  `json:"displayUsedVolume" gorm:"-"`
	UsedVolumePercent float64 `json:"usedVolumePercent" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.UserUID == "" {
		user.UserUID = rnd.GenerateUID('u')
	}
	if user.AllVolume == 0 {
		user.AllVolume = DefaultUserVolume
	}
	return nil
}

func (u *User) InitUserInfo() {
	u.CalculateVolume()
}
func (u *User) CalculateVolume() {
	// 使用百分比计算
	u.UsedVolumePercent = float64(u.UsedVolume / u.AllVolume)
	u.DisplayAllVolume = format.ByteCountIEC(uint64(u.AllVolume))
	u.DisplayUsedVolume = format.ByteCountIEC(uint64(u.UsedVolume))
}
