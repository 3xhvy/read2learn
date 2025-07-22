package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeInvalidParam = 20003
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeInvalidParam: "Param is invalid",
}
