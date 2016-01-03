package noini

import (
	"bufio"
	"io"
	"strconv"
)

type Parser struct {
	modeValueType bool // if true parse as typed ini
	fs            *FilesSet
	lines         []string
	comment       []int
}

type Uri string

type ValueType string

type KeyName string

// section key
type Key struct {
	name      KeyName
	lineNum   int
	comment   []string
	valuetype ValueType
	value     string
	belongTo  *Section
}

// section in file
type Section struct {
	uri        Uri
	name       string
	keys       map[KeyName]*Key
	sum        string
	lineAmount int
	file       *File
}

// parse file
func (p *Parser) ParseSection(reader io.Reader, section *Section) error {
	stream := bufio.NewReader(reader)
	lineNum = 0
	var comment []string
	for {
		line, err := stream.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimSpace(line)
		p.lines = append(p.lines, line)
		length := len(line)
		switch {
		case l[0] == '#': // parse comment
			comment = append(comment, lineNum-1)
			goto next
		case l[0] == '[' && l[length-1] == ']': // New section.
			goto next
		default:
			pos := strings.IndexAny(line, "=")
			if pos <= 0 {
				return nil, fmt.Errorf("ParseSection error: line num: %d ", strconv.Itoa(pos))
			}
			// parse the key
			left := strings.TrimSpace(line[0:pos])
			fields := strings.Fields(left)
			key := Key{name: left, lineNum: lineNum, comment: comment, belongTo: &section}
			// parse the value
			var vt ValueType = "string"
			if len(fields) == 2 {
				vt = strings.TrimSpace(fields[1])
			}
			key.valuetype = vt
			key.value = strings.TrimSpace(text[pos+1:])
			section[key.name] = &key
			key.comment = comment
			comment = comment[0:0]
			goto next
		}
	next:
		lineNum++
		if err == io.EOF {
			section.lineAmount = lineNum + 1
			return nil, section
		}
	}
}

// 一组相关的file集合（比如文件名相同的），递归扫描root目录下所有同名文件,并且组织成同一个集合,用于类型检测
// 类型检测递归扫描路径，发现冲突记录第一个冲突位置和文件
type FilesSet struct {
	fileName      FileName
	valueTypeBook map[KeyName]ValueType // valueType of all key in deferent file
	firstConflict map[Uri]KeyName       // file Uri and conflicted KeyName
}
