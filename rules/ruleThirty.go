package rules

var Thirty Rule = Rule{
	false,
	false,
	false,
	true,
	true,
	true,
	true,
	false,
}

func init() {
	register(30, Thirty)
}
