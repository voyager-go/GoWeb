package costant

type Status int

const (
	Enabled  Status = 1 // 启用
	Disabled Status = 2 // 禁用
)

// GetStatusComment 获取状态枚举的注释说明
func GetStatusComment(status Status) string {
	switch status {
	case Enabled:
		return "启用"
	case Disabled:
		return "禁用"
	default:
		return "未知"
	}
}
