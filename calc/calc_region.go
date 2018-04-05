package calc

// Типы участков
// TODO: добавить недостающие типы участков: Кирпичный забор, Сварной забор, Калитки, Ворота и т.д.
type Region struct {
	Name       string
	Components []int // Типы компонентов по умолчанию, для начального заполнения бд, могут быть изменены
}

type RegionTypeID int

const (
	RTProject RegionTypeID = iota
	RTProflistFenceVer
	RTProflistFenceHor
	RTBattenFenceVer
	RTBattenFenceHor
	RTKnxProflistFence
	RTKnxBattenFence
	RTKnxEcoProflistFence
	RTKnxEcoBattenFence
	RT2D
	RT3D
	RTGrandLine
)

var Regions = [...]Region{
	RTProject: {
		Name: "Проект",
	},
	RTProflistFenceVer: {
		Name:       "Классический забор из профнастила вертикальный",
		Components: []int{CMColumns, CMHStick},
	},

	RTProflistFenceHor: {
		Name:       "Классический забор из профнастила горизонтальный",
		Components: []int{CMColumns},
	},

	RTBattenFenceVer: {
		Name:       "Классический забор из штакетника вертикальный",
		Components: []int{CMColumns, CMHStick},
	},
	RTBattenFenceHor: {
		Name:       "Классический забор из штакетника горизонтальный",
		Components: []int{CMColumns},
	},
	RTKnxProflistFence: {
		Name:       "Забор KNX из профнастила",
		Components: []int{CMColumns, CMHStick},
	},
	RTKnxBattenFence: {
		Name:       "Забор KNX из штакетника",
		Components: []int{CMColumns, CMHStick},
	},
	RTKnxEcoProflistFence: {
		Name:       "Забор KNX ЭКО из профнастила",
		Components: []int{CMColumns, CMHStick},
	},
	RTKnxEcoBattenFence: {
		Name:       "Забор KNX ЭКО из штакетника",
		Components: []int{CMColumns, CMHStick},
	},
	RT2D: {
		Name:       "2D забор",
		Components: []int{CMColumns, CMHStick},
	},
	RT3D: {
		Name:       "3D забор",
		Components: []int{CMColumns, CMHStick},
	},
	RTGrandLine: {
		Name:       "Модульный забор Grand Line",
		Components: []int{CMColumns, CMHStick},
	},
}
