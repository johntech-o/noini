// not only ini parser
package noini

type ini struct {
	s *Storage // storage for all the scections of the input path
	r *Router  // explain the subscribe uri to session (a group of sections in different files)
}

func New(path, uriPrefix string) (*ini, error) {
	storage := NewStorage(path, uriPrefix)
	err := storage.Parse()
	if err != nil {
		return nil, err
	}
	router := NewRouter()
	return &ini{s: storage, r: router}, nil
}

// register your own uri router otherwise use default route parser
func (i *ini) RegisterRouter(parser ParseFunc) {
	i.r.parser = parser
}

// uri slice subscribed by user order by priority
// according router to parse uri slice ,return session to user
func (i *ini) SubByUri(uris []Uri) (*Session, error) {
	return i.r.Parse(uris)
}
