package calc

// Типы входящих параметров для расчета
type ParamTypeID int

const (
	PTTotalLength ParamTypeID = iota
	PTTotalHeight
	PTBottomSpace
	PTUpSpace
	PTColumnDepth
	PTColumnStepLength
	PTColumnStepType
	PTColumnStepSpace
	PTColumnInstallMethod
	PTColumnSize
	PTBoolInstallColumns
	PTBoolColumnPaint
	PTBoolColumnCover
	PTColorColumnPaint
	PTBoolColumnBrackets
	PTProfileSheetThickness
	PTProfileSheetType
	PTBoolInstallCanvas
	PTBoolCanvasPaint
	PTColorCanvas
	PTColorCanvasPaint
	PTColorFix
	PTBoolFix
	PTHStickCount
	PTHStickBottomSpace
	PTHStickUpSpace
	PTHStickLeghth
	PTHStickSize
	PTBoolInstallHStick
	PTColorHStick
	PTBoolHStickPaint
)

const BoolParamName string = "<bool>" // Special label in Value.Name to mark the parameter as a checkbox (ON/OFF)
const ColorParamName string = "<color>"

const UndefinedParamValue = -1 // This value means the parameter is not set

type ParamValue struct {
	Value           float32                   // Parameter value
	Name            string                    // The value description, if needed
	DependentParams map[ParamTypeID][]float32 //
}

type Param struct {
	// Parameter name, can be used as label in GUI
	Name string

	// Parameter description, can be used as a hint, there can be different parameters with the same name, detailed description helps to distinguish them
	Description string

	// List of all the possible values of this parameter
	// The first value is used as default value
	// If the slice is empty, use 0.0 ase default value
	// If the slice is empty or contains only 1 value, the parameter can be modified by user to any value
	// If the slice contains more than 1 value, the parameter can only be selected from the given list of values
	Values []ParamValue

	// Region types, where the parameter make is available
	RegionTypes []RegionTypeID
}

// Шаг столбов, разбиение
const (
	ColumnStepSpecified float32 = iota // Шаг столбов: заданный
	ColumnStepEquable                  // Шаг столбов: разбить на ровные участки
)

// Шаг столбов, остаток
const (
	ColumnStepSpaceEnd   float32 = iota // Шаг столбов: заданный, остаток в конце
	ColumnStepSpaceStart                // Шаг столбов: заданный, остаток в начале
)

// Метод установки столбов
const (
	ColumnInstallMethodСoncreting float32 = iota // Бетонирование
	ColumnInstallMethodButting                   // Бутирование
	ColumnInstallMethodHILST                     // HILST
	ColumnInstallMethodFlanges                   // Фланцы
)

var Params = [...]Param{
	PTTotalLength: {
		Name: "Длина участка, в м",
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTTotalHeight: {
		Name: "Высота забора, в м",
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTBottomSpace: {
		Name: "Зазор снизу, в мм",
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTUpSpace: {
		Name: "Выступ стоблов сверху, в мм",
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnDepth: {
		Name: "Заглубление столбов, в мм",
		Values: []ParamValue{
			{Value: 1}, // TODO: Ask for default value
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnStepLength: {
		Name: "Шаг столбов, в метрах",
		Values: []ParamValue{
			{Value: 2},
			{Value: 3},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnStepType: {
		Name: "Шаг столбов, разбиение",
		Values: []ParamValue{
			{Value: ColumnStepSpecified, Name: "Заданный"},
			{Value: ColumnStepEquable, Name: "Разбить на ровные участки",
				DependentParams: map[ParamTypeID][]float32{
					PTColumnStepSpace: {},
				},
			},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnStepSpace: {
		Name: "Шаг столбов, остаток",
		Values: []ParamValue{
			{Value: ColumnStepSpaceEnd, Name: "Остаток в конце"},
			{Value: ColumnStepSpaceStart, Name: "Остаток в начале"},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnInstallMethod: {
		Name: "Метод установки столбов",
		Values: []ParamValue{
			{Value: ColumnInstallMethodСoncreting, Name: "Бетонирование"},
			{Value: ColumnInstallMethodButting, Name: "Бутирование"},
			{Value: ColumnInstallMethodHILST, Name: "HILST"},
			{Value: ColumnInstallMethodFlanges, Name: "Фланцы"},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTColumnSize: {
		Name:        "Материалы",
		Description: "Размер столбов",
		Values: []ParamValue{
			{Value: 0, Name: "60x40x2"},
			{Value: 1, Name: "60x40x3"},
			{Value: 2, Name: "60x60x2"},
			{Value: 3, Name: "60x60x3"},
			{Value: 4, Name: "80x80x3"},
			{Value: 5, Name: "100x100x5"},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
			RTBattenFenceVer,
			RTBattenFenceHor,
			RTKnxProflistFence,
			RTKnxBattenFence,
			RTKnxEcoProflistFence,
			RTKnxEcoBattenFence,
			RT2D,
			RT3D,
			RTGrandLine,
		},
	},
	PTBoolInstallColumns: { // Галочка "Монтаж" столбов
		Name:        "Монтаж",
		Description: "Монтаж столбов",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTBoolColumnPaint: {
		Name:        "Окрашивание",
		Description: "Окрашивание столбов",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float32{
					PTColorColumnPaint: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTBoolColumnCover: {
		Name:        "Заглушки",
		Description: "Заглушки для столбов",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTColorColumnPaint: {
		Name:        "Цвет",
		Description: "Цвет окрашивания столбов",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTBoolColumnBrackets: {
		Name:        "Кронштейны",
		Description: "Кронштейны на столбах",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTProfileSheetThickness: {
		Name: "Толщина профлиста, мм",
		Values: []ParamValue{
			{Value: 0.4},
			{Value: 0.45},
			{Value: 0.5},
			{Value: 0.55},
			{Value: 0.6},
			{Value: 0.65},
			{Value: 0.7},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
		},
	},
	PTProfileSheetType: {
		Name:        "Тип",
		Description: "Тип профлиста",
		Values: []ParamValue{
			{Value: 0, Name: "ССм 10"},
			{Value: 1, Name: "С10 М1"},
			{Value: 2, Name: "Мп20"},
			{Value: 3, Name: "С21"},
			{Value: 4, Name: "С44"},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
			RTProflistFenceHor,
		},
	},
	PTBoolInstallCanvas: { // Галочка "Монтаж" забора
		Name:        "Монтаж",
		Description: "Монтаж полотна забора",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTBoolCanvasPaint: {
		Name:        "Окрашивание",
		Description: "Окрашивание полотна забора",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float32{
					PTColorCanvasPaint: {},
				},
			},
			{
				Value: 1,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float32{
					PTColorCanvas: {NoColor},
				},
			},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTColorCanvas: {
		Name:        "Цвет",
		Description: "Цвет полотна забора",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTColorCanvasPaint: {
		Name:        "Цвет",
		Description: "Цвет окрашивания полотна",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTColorFix: {
		Name:        "Цвет",
		Description: "Цвет крепежа",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTBoolFix: {
		Name:        "Крепеж",
		Description: "Крепеж полотна",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float32{
					PTColorFix: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTHStickCount: {
		Name:        "Количество",
		Description: "Количество прожилин",
		Values: []ParamValue{
			{Value: 2},
			{Value: 3},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTHStickBottomSpace: {
		Name: "Нижняя прожилина из низа профнастила, мм",
		Values: []ParamValue{
			{Value: 300},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTHStickUpSpace: {
		Name: "Верхняя прожилина от верха столбов, мм",
		Values: []ParamValue{
			{Value: 250},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTHStickLeghth: { // TODO: Ask, Is in possible to enter any value??
		Name:        "Длина, м",
		Description: "Длина прожилин, м",
		Values: []ParamValue{
			{Value: 2},
			{Value: 3},
			{Value: 4},
			{Value: 6},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTHStickSize: {
		Name:        "Материалы",
		Description: "Размер прожилин",
		Values: []ParamValue{
			{Value: 0, Name: "40x20x1.5"},
			{Value: 1, Name: "40x40x2"},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTBoolInstallHStick: { // Галочка "Монтаж" прожилин
		Name:        "Монтаж",
		Description: "Монтаж прожилин",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTColorHStick: {
		Name:        "Цвет",
		Description: "Цвет прожилин",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
	PTBoolHStickPaint: {
		Name:        "Окрашивание",
		Description: "Окрашивание прожилин",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float32{
					PTColorHStick: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
		RegionTypes: []RegionTypeID{
			RTProflistFenceVer,
		},
	},
}
