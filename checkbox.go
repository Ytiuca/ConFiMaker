package main

type Checkbox struct {
	widgetName string
	label      string
	arg        string
	IntVar     string
	isChecked  bool
}

func newCheckbox(label string, isChecked bool) *Checkbox {
	return &Checkbox{
		"CTkCheckBox",
		"_" + label,
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
		DOUBLE_INDENT + cb.label + ".pack(pady=(0,10))" +
		NEWLINE
}

func (cb *Checkbox) ToGetter() string {
	return DOUBLE_INDENT + cb.label + "=" + cb.IntVar + ".get()" +
		NEWLINE
}

func (cb *Checkbox) ToArg() string {
	return ",f" + QUOTATION + "{" + QUOTATION + "--" + cb.arg + QUOTATION + "if " + cb.label + "==1 else " + QUOTATION + "--no-" + cb.arg + QUOTATION + "}" + QUOTATION
}
