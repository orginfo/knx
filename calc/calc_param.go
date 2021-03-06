package calc

// Типы входящих параметров для расчета
type ParamTypeID int64

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
	Value           float64                   // Parameter value
	Name            string                    // The value description, if needed
	DependentParams map[ParamTypeID][]float64 //
}

type Param struct {
	// Parameter priority, determines can it depend on another parameter or not
	// To be dependent the parameter should have higher priority than the main parameter
	Prio int

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
}

// Шаг столбов, разбиение
const (
	ColumnStepSpecified float64 = iota // Шаг столбов: заданный
	ColumnStepEquable                  // Шаг столбов: разбить на ровные участки
)

// Шаг столбов, остаток
const (
	ColumnStepSpaceEnd   float64 = iota // Шаг столбов: заданный, остаток в конце
	ColumnStepSpaceStart                // Шаг столбов: заданный, остаток в начале
)

// Метод установки столбов
const (
	ColumnInstallMethodСoncreting float64 = iota // Бетонирование
	ColumnInstallMethodButting                   // Бутирование
	ColumnInstallMethodHILST                     // HILST
	ColumnInstallMethodFlanges                   // Фланцы
)

var Params = [...]Param{
	PTTotalLength: {
		Name: "Длина участка, в м",
	},
	PTTotalHeight: {
		Name: "Высота забора, в м",
	},
	PTBottomSpace: {
		Name: "Зазор снизу, в мм",
	},
	PTUpSpace: {
		Name: "Выступ стоблов сверху, в мм",
	},
	PTColumnDepth: {
		Name: "Заглубление столбов, в мм",
		Values: []ParamValue{
			{Value: 1}, // TODO: Ask for default value
		},
	},
	PTColumnStepLength: {
		Name: "Шаг столбов, в метрах",
		Values: []ParamValue{
			{Value: 2},
			{Value: 3},
		},
	},
	PTColumnStepType: {
		Name: "Шаг столбов, разбиение",
		Values: []ParamValue{
			{Value: ColumnStepSpecified, Name: "Заданный"},
			{Value: ColumnStepEquable, Name: "Разбить на ровные участки",
				DependentParams: map[ParamTypeID][]float64{
					PTColumnStepSpace: {},
				},
			},
		},
	},
	PTColumnStepSpace: {
		Prio: 1,
		Name: "Шаг столбов, остаток",
		Values: []ParamValue{
			{Value: ColumnStepSpaceEnd, Name: "Остаток в конце"},
			{Value: ColumnStepSpaceStart, Name: "Остаток в начале"},
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
				DependentParams: map[ParamTypeID][]float64{
					PTColorColumnPaint: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTBoolColumnCover: {
		Name:        "Заглушки",
		Description: "Заглушки для столбов",
		Values: []ParamValue{
			{Value: 0, Name: BoolParamName},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTColorColumnPaint: {
		Prio:        1,
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
				DependentParams: map[ParamTypeID][]float64{
					PTColorCanvasPaint: {},
				},
			},
			{
				Value: 1,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float64{
					PTColorCanvas: {NoColor},
				},
			},
		},
	},
	PTColorCanvas: {
		Prio:        1,
		Name:        "Цвет",
		Description: "Цвет полотна забора",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTColorCanvasPaint: {
		Prio:        1,
		Name:        "Цвет",
		Description: "Цвет окрашивания полотна",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTColorFix: {
		Prio:        1,
		Name:        "Цвет",
		Description: "Цвет крепежа",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTBoolFix: {
		Name:        "Крепеж",
		Description: "Крепеж полотна",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float64{
					PTColorFix: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
	},
	PTHStickCount: {
		Name:        "Количество",
		Description: "Количество прожилин",
		Values: []ParamValue{
			{Value: 2},
			{Value: 3},
		},
	},
	PTHStickBottomSpace: {
		Name: "Нижняя прожилина из низа профнастила, мм",
		Values: []ParamValue{
			{Value: 300},
		},
	},
	PTHStickUpSpace: {
		Name: "Верхняя прожилина от верха столбов, мм",
		Values: []ParamValue{
			{Value: 250},
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
	},
	PTHStickSize: {
		Name:        "Материалы",
		Description: "Размер прожилин",
		Values: []ParamValue{
			{Value: 0, Name: "40x20x1.5"},
			{Value: 1, Name: "40x40x2"},
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
		Prio:        1,
		Name:        "Цвет",
		Description: "Цвет прожилин",
		Values: []ParamValue{
			{Value: 0, Name: ColorParamName},
		},
	},
	PTBoolHStickPaint: {
		Name:        "Окрашивание",
		Description: "Окрашивание прожилин",
		Values: []ParamValue{
			{
				Value: 0,
				Name:  BoolParamName,
				DependentParams: map[ParamTypeID][]float64{
					PTColorHStick: {},
				},
			},
			{Value: 1, Name: BoolParamName},
		},
	},
}
