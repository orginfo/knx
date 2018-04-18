package calc

// Results of the calculation
type ResultTypeID int64

const (
	RSNomenclature ResultTypeID = iota
)

type Result struct {
	Name        string
	Description string
}

var Results = [...]Result{
	RSNomenclature: {
		Name:        "Позиция",
		Description: "Номенклатура, расход материала или услуги",
	},
}
