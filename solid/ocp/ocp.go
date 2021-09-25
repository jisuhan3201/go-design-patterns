package ocp

type Product struct {
	name  string
	color string
	size  string
}

type Filter struct {
}

// Product가 가지는 color, name등으로 필터를 하려면
// 아래와 같이 코드의 중복이 일어난다. 따라서 개방적이지 않다(Open 에 제한적이다)
func (f *Filter) filterByColor(
	products []Product, color string) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}
func (f *Filter) filterBySize(
	products []Product, size string) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == size {
			result = append(result, &products[i])
		}
	}
	return result
}
func (f *Filter) filterByColorAndSize(
	products []Product, color, size string) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color string
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size string
}

func (spec *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// 여러 조건은?
type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type BetterFilter struct {
}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
