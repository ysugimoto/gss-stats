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
	f.WriteLine("Data URI Size", fmt.Sprintf("%d bytes", s.DataUriSize))
	f.WriteSeparator()
	f.WriteLine("Ratio of Data URI Size", fmt.Sprintf("%3f %%", s.RatioOfDataUriSize))
	f.WriteSeparator()
	f.WriteLine("Gzipped Size", fmt.Sprintf("%d bytes", s.GzippedSize))
	f.WriteSeparator()
	f.WriteLine("Rules", fmt.Sprintf("%d rules", s.Rules))
	f.WriteSeparator()
	f.WriteLine("Selectors", fmt.Sprintf("%d selectors", s.Selectors))
	f.WriteSeparator()
	f.WriteLine("Simplicity", fmt.Sprintf("%3f %%", s.Simplicity))
	f.WriteSeparator()
	f.WriteLine("Average of Identifier", fmt.Sprintf("%3f", s.AverageOfIdentifier))
	f.WriteSeparator()
	f.WriteLine("Average of Cohesion", fmt.Sprintf("%3f", s.AverageOfCohesion))
	f.WriteSeparator()
	f.WriteLine("Most Identifier", fmt.Sprintf("%d", s.MostIdentifier))
	f.WriteSeparator()
	f.WriteLine("Most Identifier Selector", s.MostIdentifierSelector)
	f.WriteSeparator()
	f.WriteLine("Lowest Cohesion", fmt.Sprintf("%d", s.LowestCohesion))
	f.WriteSeparator()
	f.WriteLine("Lowest Cohesion Selector", s.LowestCohesionSelector)

	f.WriteSeparator()
	f.WriteLine("Total Uniqle Font Sizes", fmt.Sprintf("%d", s.TotalUniqueFontSizes))
	if s.TotalUniqueFontSizes > 0 {
		f.WriteSeparator()
	}
	for index, fontSize := range s.UniqueFontSizes {
		if index == 0 {
			f.WriteLine("Unique Font Sizes", fontSize)
		} else {
			f.WriteLine("", fontSize)
		}
	}

	f.WriteSeparator()
	f.WriteLine("Total Unique Colors", fmt.Sprintf("%d", s.TotalUniqueColors))
	if s.TotalUniqueColors > 0 {
		f.WriteSeparator()
	}
	for index, color := range s.UniqueColors {
		if index == 0 {
			f.WriteLine("Unique Colors", color)
		} else {
			f.WriteLine("", color)
		}
	}

	f.WriteSeparator()
	f.WriteLine("Total Unique Font Families", fmt.Sprintf("%d", s.TotalUniqueFontFamilies))
	if s.TotalUniqueFontFamilies > 0 {
		f.WriteSeparator()
	}
	for index, fontFamily := range s.UniqueFontFamilies {
		if index == 0 {
			f.WriteLine("Unique Font Families", fontFamily)
		} else {
			f.WriteLine("", fontFamily)
		}
	}

	f.WriteSeparator()
	f.WriteLine("ID Selectors", fmt.Sprintf("%d", s.IdSelectors))
	f.WriteSeparator()
	f.WriteLine("Universal Selectors", fmt.Sprintf("%d", s.UniversalSelectors))
	f.WriteSeparator()
	f.WriteLine("Unqualified Attribute Selectors", fmt.Sprintf("%d", s.UnqualifiedAttributeSelectors))
	f.WriteSeparator()
	f.WriteLine("JavaScript Specific Selectors", fmt.Sprintf("%d", s.JavaScriptSpecificSelectors))
	f.WriteSeparator()
	f.WriteLine("Important Keywords", fmt.Sprintf("%d", s.ImportantKeywords))
	f.WriteSeparator()
	f.WriteLine("Float Properties", fmt.Sprintf("%d", s.FloatProperties))

	if len(s.PropertiesCount) > 0 {
		f.WriteSeparator()
	}
	for index, prop := range s.PropertiesCount {
		if index == 0 {
			f.WriteLine("Properties Count", fmt.Sprintf("%s: %d", prop.Property, prop.Count))
		} else {
			f.WriteLine("", fmt.Sprintf("%s: %d", prop.Property, prop.Count))
		}
	}

	f.WriteSeparator()
	f.WriteLine("Media Queries", fmt.Sprintf("%d", s.MediaQueries))
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
