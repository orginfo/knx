package calc

// Типы участков
// TODO: добавить недостающие типы участков: Кирпичный забор, Сварной забор, Калитки, Ворота и т.д.
type Region struct {
	Name       string
	Components []ComponentTypeID // Типы компонентов по умолчанию, для начального заполнения бд, могут быть изменены
}

type RegionTypeID int64

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
		Components: []ComponentTypeID{CMRegion, CMColumns, CMHStick, CMFilling},
	},

	RTProflistFenceHor: {
		Name:       "Классический забор из профнастила горизонтальный",
		Components: []ComponentTypeID{CMRegion, CMColumns, CMFilling},
	},

	RTBattenFenceVer: {
		Name:       "Классический забор из штакетника вертикальный",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RTBattenFenceHor: {
		Name:       "Классический забор из штакетника горизонтальный",
		Components: []ComponentTypeID{CMColumns, CMFilling},
	},
	RTKnxProflistFence: {
		Name:       "Забор KNX из профнастила",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RTKnxBattenFence: {
		Name:       "Забор KNX из штакетника",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RTKnxEcoProflistFence: {
		Name:       "Забор KNX ЭКО из профнастила",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RTKnxEcoBattenFence: {
		Name:       "Забор KNX ЭКО из штакетника",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RT2D: {
		Name:       "2D забор",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RT3D: {
		Name:       "3D забор",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
	RTGrandLine: {
		Name:       "Модульный забор Grand Line",
		Components: []ComponentTypeID{CMColumns, CMHStick, CMFilling},
	},
}
