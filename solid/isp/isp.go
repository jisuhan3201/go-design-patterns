package isp

type Document struct {
}

// 이렇게 하나의 인터페이스에 많은 기능을 넣어놓으면
// 모든 구조체들이 이러한 기능들을 모두 가지고 있어야 한다.
// 따라서 인터페이스를 작게 분리한다.
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// ok if you need a multifunction device
type MultiFunctionPrinter struct {
	// ...
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
	// ...
}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated 된 기능인데 구현해야한다.
// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// 인터페이스를 작게 분리한다.
// better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct {
	// ...
}

func (m MyPrinter) Print(d Document) {
	// ...
}

// combine interfaces
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {
	//
}

func (p Photocopier) Print(d Document) {
	//
}

// 여러 기능을 담은 인터페이스를 만든다.
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// 아니면 여러 기능을 담은 구조체를 만든다.
// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
