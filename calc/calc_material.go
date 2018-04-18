package calc

// Тип расчета материала
type MaterialCalculationID int64

const (
	MCDoNotCalculate MaterialCalculationID = iota
	MCPaintArea
)

type MaterialCalculation struct {
	Name string
	Func func(...interface{}) []float64
}

var MaterialCalculations = [...]MaterialCalculation{
	MCDoNotCalculate: {
		Name: "Не расчитывать",
		Func: MCDoNotCalculateFunc,
	},

	MCPaintArea: {
		Name: "Покраска поверхности",
		Func: MCPaintAreaFunc,
	},
}

// MCDoNotCalculateFunc - функция-заглушка для материалов, не требующих расчета
// Входящие параметры: нет
// Результаты: пустой срез
func MCDoNotCalculateFunc(...interface{}) []float64 {
	return []float64{}
}

// MCPaintAreaFunc - расчет краски, требуемой на покраску поверхности определеной площади
// Входящие параметры:
// [0] Объем одной банки краски, в мл
// [1] Площадь поверхности, в м2
// [2] Коэффициент расхода краски, мл на 1 м2
//
// Результаты:
// [0] Количество банок краски
// [1] Объем требуемой краски, в мл
// [2] Остаток, в мл
func MCPaintAreaFunc(...interface{}) (res []float64) {
	res = []float64{0} // TODO: implement real calculations
	return
}
