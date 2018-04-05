package calc

type Color struct {
	Value int
	Name  string
}

type ColorScheme struct {
	Name   string
	Colors []Color
}

const NoColor = -1

var ColorSchemes = []ColorScheme{
	{
		Name: "RAL", Colors: []Color{
			{Value: NoColor, Name: "Цинк"},
			{Value: 0xdcc6a1, Name: "RAL1014 (сл.-бежевый)"},
			{Value: 0xe5d5bb, Name: "RAL1015 (бежевый)"},
			{Value: 0x8d764c, Name: "RAL1036 (бронза)"},
			{Value: 0xdf5e34, Name: "RAL2004 (оранжевый)"},
			{Value: 0x8a383e, Name: "RAL3003 (рубин)"},
			{Value: 0x663d43, Name: "RAL3005 (кр.вино)"},
			{Value: 0x7d3d3d, Name: "RAL3011 (красно-коричневый)"},
			{Value: 0xbb3939, Name: "RAL3020 (ярко-красный)"},
			{Value: 0x3f4b87, Name: "RAL5002 (ультрамарин)"},
			{Value: 0x2b5d8e, Name: "RAL5005 (синий)"},
			{Value: 0x097d7d, Name: "RAL5021 (мор. волна)"},
			{Value: 0x456540, Name: "RAL6002 (зеленый)"},
			{Value: 0x365148, Name: "RAL6005 (зел. мох)"},
			{Value: 0x07695a, Name: "RAL6026 (изумруд)"},
			{Value: 0x00784b, Name: "RAL6029 (зел. трава)"},
			{Value: 0xa0a0a0, Name: "RAL7004 (серый)"},
			{Value: 0x534542, Name: "RAL8017 (коричневый)"},
			{Value: 0x7b3e25, Name: "RAL8029 (медь перламутр)"},
			{Value: 0xd8d8d0, Name: "RAL9002 (белая ночь)"},
			{Value: 0xeeeeec, Name: "RAL9003 (белый)"},
			{Value: 0x6c4e3c, Name: "Printech \"Дерево\" 4201"}, // Цвет не соответствует
		},
	},
}
