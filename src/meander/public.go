package meander

type Facade interface {
	Public() interface{}
}

func Pulibc(o interface{}) interface {} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return 0
}