package dip

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

type Research struct {
	relationships Relationships
}

func (r *Research) Investigate() {
	// Research가 Relationships의 상위 모듈인데
	// 아래와 같이 Research의 메서드에 하위 모듈의 정보를 담으면
	// Relationships의 relaltion ([]Info) 가 달라지면
	// 아래 코드 또한 바뀌어야 한다.
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" &&
			rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

// 따라서 하위 모듈에서 메서드를 두고
func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

// 메서드를 포함하는 인터페이스를 두고
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// 인터페이스를 포함하는 상위 모듈을 만든다.
type Research2 struct {
	browser RelationshipBrowser
}

func (r *Research2) Investigate2() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
