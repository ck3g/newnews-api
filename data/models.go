package data

type Models struct {
	Items Item
}

func New() Models {
	return Models{
		Items: Item{},
	}
}
