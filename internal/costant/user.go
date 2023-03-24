package costant

type UserStatus int

const (
	UserEnabled  UserStatus = 1 // 启用
	UserDisabled UserStatus = 2 // 禁用
)

// GetUserStatusComment 获取状态枚举的注释说明
func GetUserStatusComment(status UserStatus) string {
	switch status {
	case UserEnabled:
		return "启用"
	case UserDisabled:
		return "禁用"
	default:
		return "未知"
	}
}
