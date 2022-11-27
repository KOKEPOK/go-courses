package converter

/*
	The structure of the private fields of the converter

for converting from one number system to another.
*/
type converter struct {
	symbols []string // Integer value. Example: 1, 5, 3.
	values  []int    // Analogue of the Latin meaning. Example: I, V, III.
}

// Converter interface, public methods.
type Converter interface {
	IntToRoman(num int) string
}

// Method for converting from decimal to Roman.
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

// Constructor.
func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
