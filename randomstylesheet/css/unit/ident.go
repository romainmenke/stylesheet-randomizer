package unit

type Ident string

func (x Ident) Css() string {
	return string(x)
}

func (x Ident) UnitType() string {
	return "ident"
}

func (x Ident) Unit() string {
	return "ident"
}

func (x Ident) Value() interface{} {
	return string(x)
}

func Transparant() Value {
	return Ident("transparant")
}

func Inherit() Value {
	return Ident("inherit")
}

func Initial() Value {
	return Ident("initial")
}

func Unset() Value {
	return Ident("unset")
}

func Normal() Value {
	return Ident("normal")
}
