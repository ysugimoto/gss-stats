package main

import (
	_ "encoding/json"
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/gssp"
	"github.com/ysugimoto/gssp-stats"
	"os"
)

func usage() {
	text := `==============================================
GSSP-STATS: Go Style-Sheet Statistics Analyzer
==============================================
Usage:
    gssp-stats [source_file ...] [options]

Options:
    -h, --help : Show this help
    -j, --json : Output as JSON string
`
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	args := cliarg.NewArguments()
	args.Alias("h", "help", false)
	args.Alias("j", "json", false)

	args.Parse()

	if help, _ := args.GetOptionAsBool("help"); help {
		usage()
	}

	if args.GetCommandSize() == 0 {
		fmt.Println("[ERROR] Source CSS file must be suppied.")
		usage()
	}

	result := gssp.NewParseResult([]*gssp.CSSDefinition{})

	files := args.GetCommands()
	for _, file := range files {
		parser := gssp.NewParser()
		if parsed, err := parser.ParseFile(file); err != nil {
			fmt.Println(err.Error)
			os.Exit(1)
		} else {
			result.Merge(parsed)
		}
	}

	stat := stats.NewStats(result)
	stat.Analyze()
	if outputJson, _ := args.GetOptionAsBool("json"); outputJson {
		//out, _ := json.Marshal(stat.PropertiesCount)
		fmt.Println(stat.ToPrettyJsonString())
	} else {
		stat.FormatOutput()
	}
}
