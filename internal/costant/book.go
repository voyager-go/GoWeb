package costant

type BookStatus int

const (
	StatusSerialized = iota + 1
	StatusPublished
	StatusUnpublished
)

func GetBookStatusComment(status BookStatus) string {
	switch status {
	case StatusSerialized:
		return "连载中"
	case StatusPublished:
		return "已发布"
	case StatusUnpublished:
		return "已下架"
	default:
		return "未知"
	}
}
