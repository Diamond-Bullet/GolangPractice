package design_pattern

// todo Visitor Pattern

type Host interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(visitFunc VisitFunc)
}

type InfoHost struct {
	F1 string
	F2 string
	F3 string
}

func (i *InfoHost) Accept(visitor Visitor) {
	visitor.Visit(func(info *InfoHost) error {
		println(info)
		return nil
	})
}

type F1Visitor struct{}

func (f *F1Visitor) Visit(visitFunc VisitFunc) {
}

type VisitFunc func(info *InfoHost) error
