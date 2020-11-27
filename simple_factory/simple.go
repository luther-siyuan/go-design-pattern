package simple_factory

// 定义一个借口
type API interface {
	Say(name string) string
}

// 根据类别返回一个实现
func NewAPI(t int) API {
	if t == 1 {
		return &hiApi{}
	}
	return &helloApi{}
}

// API 的一个实现
type hiApi struct {
}

func (*hiApi) Say(name string) string {
	return "hi"
}

// API 的一个实现
type helloApi struct {
}

func (*helloApi) Say(name string) string {
	return "hello"
}
