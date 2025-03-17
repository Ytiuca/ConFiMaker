package main

type Checkbox struct {
	widgetName string
	label      string
	IntVar     string
	isChecked  bool
}

func newCheckbox(label string, isChecked bool) *Checkbox {
	return &Checkbox{
		"CTkCheckBox",
		label,
		"self." + label + "_var",
		isChecked,
	}
}

func (cb *Checkbox) Name() string {
	return cb.widgetName
}

func (cb *Checkbox) ToPython() string {
	defaultValue := "0"
	if cb.isChecked {
		defaultValue = "1"
	}
	return DOUBLE_INDENT + cb.IntVar + "=IntVar(value=" + defaultValue + ")" +
		NEWLINE +
		DOUBLE_INDENT + cb.label + "=" + cb.widgetName + "(self,text=\"" + cb.label + "\",variable=" + cb.IntVar + ")" +
		NEWLINE +
		DOUBLE_INDENT + cb.label + ".pack()" +
		NEWLINE
}

func (cb *Checkbox) ToGetter() string {
	return DOUBLE_INDENT + cb.label + "=" + cb.IntVar + ".get()" +
		NEWLINE
}

func (cb *Checkbox) ToArg() string {
	return ",f" + QUOTATION + "{" + QUOTATION + "--" + cb.label + QUOTATION + "if " + cb.label + "==1 else " + QUOTATION + "--no-" + cb.label + QUOTATION + "}" + QUOTATION
}
