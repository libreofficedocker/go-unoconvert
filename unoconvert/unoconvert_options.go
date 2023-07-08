package unoconvert

var (
	// The file type/extension of the output file (ex pdf). Required when using stdout
	ConvertTo = "--convert-to"
	// The export filter to use when converting. It is selected automatically if not specified.
	Filter = "--filter"
	// Options for the export filter, in name=value format. Use true/false for boolean values.
	FilterOptions = "--filter-options"
	// Update the indexes before conversion. Can be time consuming.
	UpdateIndex = "--update-index"
	// Skip updating the indexes.
	DontUpdateIndex = "--dont-update-index"
)

type UnoconvertOption struct {
	Key   string
	Value string
}
