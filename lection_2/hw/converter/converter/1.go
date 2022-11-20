package converter

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(num int) string
}

func (c *converter) IntToRoman(num int) string {
	var result string
	for i := 0; i < len(c.values); i++ {
		for num >= c.values[i] {
			result += c.symbols[i]
			num -= c.values[i]
		}
	}
	return result
}

func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
