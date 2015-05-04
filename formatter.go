package stats

import (
	"fmt"
)

type Formatter struct {
}

func (f Formatter) Format(s *Stats) {
	f.WriteHeader()
	f.WriteLine("StyleSheets", fmt.Sprintf("%d files", s.StyleSheets))
	f.WriteSeparator()
	f.WriteLine("Size", fmt.Sprintf("%d bytes", s.Size))
	f.WriteSeparator()
	f.WriteLine("DataUriSize", fmt.Sprintf("%d bytes", s.DataUriSize))
	f.WriteSeparator()
	f.WriteLine("GzippedSize", fmt.Sprintf("%d bytes", s.GzippedSize))
	f.WriteSeparator()
	f.WriteLine("Rules", fmt.Sprintf("%d rules", s.Rules))
	f.WriteSeparator()
	f.WriteLine("Selectors", fmt.Sprintf("%d selectors", s.Selectors))
	f.WriteSeparator()
	f.WriteLine("Simplicity", fmt.Sprintf("%3f", s.Simplicity))
	f.WriteSeparator()
	f.WriteLine("AverageOfIdentifier", fmt.Sprintf("%3f", s.AverageOfIdentifier))
	f.WriteSeparator()
	f.WriteLine("AverageOfCohesion", fmt.Sprintf("%3f", s.AverageOfCohesion))
	f.WriteSeparator()
	f.WriteLine("MostIdentifier", fmt.Sprintf("%d", s.MostIdentifier))
	f.WriteSeparator()
	f.WriteLine("MostIdentifierSelector", s.MostIdentifierSelector)
	f.WriteSeparator()
	f.WriteLine("LowestCohesion", fmt.Sprintf("%d", s.LowestCohesion))
	f.WriteSeparator()
	f.WriteLine("LowersCohesionSelector", s.LowersCohesionSelector)

	f.WriteSeparator()
	f.WriteLine("TotalUniqueFontSizes", fmt.Sprintf("%d", s.TotalUniqueFontSizes))
	if s.TotalUniqueFontSizes > 0 {
		f.WriteSeparator()
	}
	for index, fontSize := range s.UniqueFontSizes {
		if index == 0 {
			f.WriteLine("UniqueFontSizes", fontSize)
		} else {
			f.WriteLine("", fontSize)
		}
	}

	f.WriteSeparator()
	f.WriteLine("TotalUniqueColors", fmt.Sprintf("%d", s.TotalUniqueColors))
	if s.TotalUniqueColors > 0 {
		f.WriteSeparator()
	}
	for index, color := range s.UniqueColors {
		if index == 0 {
			f.WriteLine("UniqueColors", color)
		} else {
			f.WriteLine("", color)
		}
	}

	f.WriteSeparator()
	f.WriteLine("TotalUniqueFontFamilies", fmt.Sprintf("%d", s.TotalUniqueFontFamilies))
	if s.TotalUniqueFontFamilies > 0 {
		f.WriteSeparator()
	}
	for index, fontFamily := range s.UniqueFontFamilies {
		if index == 0 {
			f.WriteLine("UniqueFontFamilies", fontFamily)
		} else {
			f.WriteLine("", fontFamily)
		}
	}

	f.WriteSeparator()
	f.WriteLine("IdSelectors", fmt.Sprintf("%d", s.IdSelectors))
	f.WriteSeparator()
	f.WriteLine("UniversalSelectors", fmt.Sprintf("%d", s.UniversalSelectors))
	f.WriteSeparator()
	f.WriteLine("UnqualifiedAttributeSelectors", fmt.Sprintf("%d", s.UnqualifiedAttributeSelectors))
	f.WriteSeparator()
	f.WriteLine("JavaScriptSpecificSelectors", fmt.Sprintf("%d", s.JavaScriptSpecificSelectors))
	f.WriteSeparator()
	f.WriteLine("ImportantKeywords", fmt.Sprintf("%d", s.ImportantKeywords))
	f.WriteSeparator()
	f.WriteLine("FloatProperties", fmt.Sprintf("%d", s.FloatProperties))

	if len(s.PropertiesCount) > 0 {
		f.WriteSeparator()
	}
	for index, prop := range s.PropertiesCount {
		if index == 0 {
			f.WriteLine("PropertiesCount", fmt.Sprintf("%s: %d", prop.property, prop.count))
		} else {
			f.WriteLine("", fmt.Sprintf("%s: %d", prop.property, prop.count))
		}
	}
	f.WriteFooter()

}

func (f Formatter) WriteHeader() {
	fmt.Printf("===============================================================================\n")
	fmt.Printf("| %-75s |\n", "Gssp-Stats Result")
	fmt.Printf("===============================================================================\n")
}

func (f Formatter) WriteLine(prop, value string) {
	fmt.Printf("| %-36s | %-36s |\n", prop, value)
}

func (f Formatter) WriteSeparator() {
	fmt.Printf("|-----------------------------------------------------------------------------|\n")
}

func (f Formatter) WriteFooter() {
	fmt.Printf("===============================================================================\n")
}
