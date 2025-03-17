package main

type Input struct {
	widgetName  string
	label       string
	StringVar   string
	defaultText string
}

func newInput(label string, defaultText string) *Input {
	return &Input{
		"CTkEntry",
		label,
		"self." + label + "_var",
		defaultText,
	}
}

func (in *Input) Name() string {
	return in.widgetName
}

func (in *Input) ToPython() string {
	return DOUBLE_INDENT + in.StringVar + "=StringVar(value=" + QUOTATION + in.defaultText + QUOTATION + ")" +
		NEWLINE +
		DOUBLE_INDENT + in.label + "=" + in.widgetName + "(self,placeholder_text=" + QUOTATION + in.label + QUOTATION + ",textvariable=" + in.StringVar + ")" +
		NEWLINE +
		DOUBLE_INDENT + in.label + ".pack()" +
		NEWLINE
}

func (in *Input) ToGetter() string {
	return DOUBLE_INDENT + in.label + "=" + in.StringVar + ".get()" +
		NEWLINE
}

func (in *Input) ToArg() string {
	return ",f" + QUOTATION + "--" + in.label + "={" + in.label + "}" + QUOTATION
}
