package syserror

type ErrorNoLogin struct {
	UnKnowError
}
func(this ErrorNoLogin)Code()int{
	return 1005
}
func(this ErrorNoLogin)Error()string{

	return "未登入"
}