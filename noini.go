// 入口文件
package noini

type ini struct {
	s *Storage
	r *Router
}

func New(path string) (*ini, error) {
	storage := NewStorage()
	err := storage.ParsePath(path)
	if err != nil {
		return nil, err
	}
	router = NewRouter(storage)
	return &ini{s: storage, r: router}
	return i, nil
}

// register your own uri router otherwise use default route parser
func (i *ini) RegisterRouter(parser ParseFunc) {
	r.parser = parser
}

// 订阅指定uri，按优先级传入，会返回订阅的 Session
// 用户订阅的uri是一个Session，会根据默认的Router来解析订阅规则
func (i *ini) SubByUri(uri []string) (*Session, error) {
	return i.r.Parse(subs)
}
