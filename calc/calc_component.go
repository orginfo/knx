package calc

type ComponentTypeID int64

const (
	CMRegion ComponentTypeID = iota
	CMColumns
	CMHStick
	CMFilling
)

type Component struct {
	Name   string
	Parts  []Part
	Params []ParamTypeID
}

var Components = [...]Component{

	// Common parameters
	CMRegion: {
		Name:   "Параметры участка",
		Params: []ParamTypeID{PTTotalLength, PTTotalHeight, PTBottomSpace, PTUpSpace},
	},

	// Columns
	CMColumns: {
		Name: "Столбы",
		Parts: []Part{
			{Name: "Материал"},
			{Name: "Кронштейны"},
			{Name: "Заглушки"},
			{Name: "Монтаж"},
			{Name: "Покраска(услуга)"},
			{Name: "Покраска(материал)"},
			{Name: "Цемент"},
			{Name: "Щебень"},
			{Name: "Песок"},
			{Name: "HILST"},
		},
		Params: []ParamTypeID{
			PTColumnDepth,
			PTColumnStepLength,
			PTColumnStepType,
			PTColumnStepSpace,
			PTColumnInstallMethod,
			PTColumnSize,
			PTBoolInstallColumns,
			PTBoolColumnPaint,
			PTBoolColumnCover,
			PTColorColumnPaint,
			PTBoolColumnBrackets,
		},
	},

	// Horizintal Stick
	CMHStick: {
		Name: "Прожилины",
		Parts: []Part{
			{Name: "Материал"},
			{Name: "Монтаж"},
			{Name: "Окрашивание(услуга)"},
			{Name: "Окрашивание(материал)"},
		},
		Params: []ParamTypeID{
			PTHStickCount,
			PTHStickBottomSpace,
			PTHStickUpSpace,
			PTHStickLeghth,
			PTHStickSize,
			PTBoolInstallHStick,
			PTColorHStick,
			PTBoolHStickPaint,
		},
	},

	// Filling
	CMFilling: {
		Name: "Наполнение",
		Parts: []Part{
			{Name: "Материал"},
			{Name: "Крепеж"},
			{Name: "Монтаж"},
			{Name: "Окрашивание(услуга)"},
			{Name: "Окрашивание(материал)"},
		},
		Params: []ParamTypeID{
			PTProfileSheetThickness,
			PTProfileSheetType,
			PTBoolInstallCanvas,
			PTBoolCanvasPaint,
			PTColorCanvas,
			PTColorCanvasPaint,
			PTColorFix,
			PTBoolFix,
		},
	},
}
