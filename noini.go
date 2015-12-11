package noini

type Uri string

type Cmd struct {
	u Uri
	k Key
	v interface{}
}

type iniStorage struct {
	version  uint
	recvBuf  chan *Cmd
	files    map[Uri]*file
	Sections map[Uri]*Section
}

func (i *ini) GetSectionsByUri(uri []string) *SectionBook {

}

func (i *ini) NextCmd() *Cmd {
	select {
	case c, ok := <-i.recvBuf:
		if !ok {
			break
		}
		return c, nil
	}
	return nil, errors.New("close cmd channel")
}

func Parse(path string) *ini {

}
