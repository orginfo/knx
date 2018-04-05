package calc

type Component struct {
	Name  string
	Parts []Part
}

const (
	CMColumns int = iota
	CMHStick
)

var Components = [...]Component{
	CMColumns: {
		Name: "Столбы",
		Parts: []Part{
			{Name: "Столб", MC: MCDoNotCalculate},
		},
	},

	CMHStick: {
		Name: "Прожилины",
		Parts: []Part{
			{Name: "Прожилина", MC: MCDoNotCalculate},
		},
	},
}
