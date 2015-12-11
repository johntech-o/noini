package noini

type ValueType string

type Parser struct {
	modeValueType      bool
	modeValueTypeCheck int
	checkMap           map[FileName]map[KeyName]ValueType
}

type KeyName string

type Key struct {
	name    KeyName
	lineNum int
	comment []string
	section *Section
}

type Value struct {
	typ     ValueType
	section *Section
}

type SectionBook struct {
	mergeSection  *Section
	sectionMap    map[Uri]*Section
	sortedSection []*Section
	root          *ini
}

func (s *SectionBook) WriteCmd(cmd) {

}

type Section struct {
	uri        string
	parent     *Section
	name       string
	key        Key
	value      Value
	sum        string
	priority   int
	lineAmount int
	file       *File
}

type FileName string

type File struct {
	uri      Uri
	name     FileName
	ext      string
	sections map[Uri]*Section
	lines    []byte
}
