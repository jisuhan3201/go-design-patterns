package srp

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

var entryCount = 0

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d : %s", entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func RemoveEntry(index int) {
	// ...
}

// 여기까지가 Journal struct 의 entries 값에 대한 편집  (Add, Remove 등) 의 구현이고

// 아래와 같이 Load or Save 할 때에 메서드로 하면 단일 책임 원칙에서 벗어난다.
// 또한 다른 구현체 (Journal이 아닌 Book 등)이 생긴다면 Load, Save 등 같은 메서드를 구현해야한다.
func (j *Journal) Load(filename string) {

}
func (j *Journal) LoadFromWeb(url *url.URL) {

}

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

// 따라서 아래와 같이 단일 책입 원칙에 따라 독립된 함수로 만들거나,
// Persistence와 같은 구현체로 분리한다.
var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}
