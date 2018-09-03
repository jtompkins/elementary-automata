package rules

var Ninety Rule = Rule{
	false,
	true,
	false,
	true,
	true,
	false,
	true,
	false,
}

func init() {
	register(90, Ninety)
}
