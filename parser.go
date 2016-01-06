package noini

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	DEFAULT_VALUE_TYPE   = "string"
	DEFAULT_SECTION_NAME = "/defalut"
)

// Parser store the status and meta info through parsing the whole path
type Parser struct {
	modeValueType bool // if true parse as typed ini
	fs            *FilesSet
	lines         []string
	comment       []string
	uriPrefix     string   // uriPrefix to trim
	current       *Section // current section
	fileUri       Uri
	reader        *bufio.Reader // input stream
}

func (p *Parser) Uri(path string) (Uri, error) {
	if pos := strings.IndexAny(path, p.uriPrefix); pos == 0 {
		return Uri(path[len(p.uriPrefix):]), nil
	}
	return "", fmt.Errorf("invalid uri path:%s or uriPrefix %s ", path, p.uriPrefix)
}

type Uri string

type ValueType string

type KeyName string

// section key
type Key struct {
	name      KeyName
	lineNum   int
	comments  []string
	valuetype ValueType
	value     string
	belongTo  *Section
}

// section in file
type Section struct {
	uri      Uri
	name     string
	keys     map[KeyName]*Key
	sum      string
	file     *File
	comments []string
	lines    []string
}

func (s *Section) Uri() Uri {
	return s.uri
}

func (p *Parser) setReader(reader io.Reader) {
	p.reader = bufio.NewReader(reader)
	p.current = &Section{}
	p.current.uri = p.fileUri + DEFAULT_SECTION_NAME
	return
}

// parse file's section
// @todo validate value of the same key in the same file name have the same valuetype
// @todo validate valuetype
func (p *Parser) nextSection() (*Section, error) {
	var comments []string
	var result *Section
	for {
		line, err := p.reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimSpace(line)
		length := len(line)
		switch {
		case line[0] == '[' && line[length-1] == ']': // New section.
			result := p.current
			goto next
			result.comments = comments
			p.current = &Section{uri: Uri(p.uriPrefix + "/" + line[1:length-1])}
			p.current.lines = append(p.current.lines, line)
			return result, nil
		case line[0] == '#':
			comments = append(comments, line)
			goto next
		default:
			pos := strings.IndexAny(line, "=")
			if pos <= 0 {
				return nil, fmt.Errorf("ParseSection error: line num: %d ", strconv.Itoa(pos))
			}
			// parse the key
			left := strings.TrimSpace(line[0:pos])
			fields := strings.Fields(left)
			key := Key{name: KeyName(left), comments: comments, belongTo: result}
			// parse the value
			var vtype ValueType = DEFAULT_VALUE_TYPE
			if len(fields) == 2 {
				vtype = ValueType(strings.TrimSpace(fields[1]))
			}
			key.valuetype = vtype
			key.value = strings.TrimSpace(line[pos+1:])
			p.current.keys[key.name] = &key
			key.comments = comments
			comments = comments[0:0]
			goto next
		}
	next:
		p.current.lines = append(p.current.lines, line)
		if err == io.EOF {
			return p.current, nil
		}
	}
}

// @todo fmt lines of the file
func (p *Parser) fmt() error {
	return nil
}

// 一组相关的file集合（比如文件名相同的），递归扫描root目录下所有同名文件,并且组织成同一个集合,用于类型检测
// 类型检测递归扫描路径，发现冲突记录第一个冲突位置和文件
type FilesSet struct {
	fileName      FileName
	valueTypeBook map[KeyName]ValueType // valueType of all key in deferent file
	firstConflict map[Uri]KeyName       // file Uri and conflicted KeyName
}
