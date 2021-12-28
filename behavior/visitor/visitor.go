package visitor

import (
	"fmt"
	"io"
	"os"
)

/*
MessageA and MessageB both have a Msg field to store the text they will print.
*/
type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

// Visitor 实现
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

// Visitable interface has a method called Accept(Visitor) that will
// execute the decoupled algorithm.
// 实现他的就是单纯的数据类, java的record
type Visitable interface {
	Accept(Visitor)
}

type MessageVisitor struct {
}

func (mv *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

func (mv *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

/*
The very important thing here is that we can add more
functionality to both Messages now without altering their types. We can just create a new
visitor type that does any other thing on the Visitable, for example, we can create a
visitor to “add” a method that prints the contents of the Msg field.
*/
type MsgFieldVisitorPrinter struct{}

func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
