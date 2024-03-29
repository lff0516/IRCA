package v1

import (
	"qa/model"
	"qa/serializer"
)

// ChangePassword 修改用户密码
type ChangePassword struct {
	Password        string `form:"password" json:"password" binding:"required,min=6,max=18"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=18"`
}

// Valid 验证表单
func (service *ChangePassword) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return serializer.ErrorResponse(serializer.CodePasswordConfirmError)
	}
	return nil
}

// Change 修改密码
func (service *ChangePassword) Change(user *model.User) *serializer.Response {

	// 表单验证
	if err := service.Valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.ErrorResponse(serializer.CodeUnknownError)
	}

	// 更新数据库
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}

	return serializer.OkResponse(nil)
}
