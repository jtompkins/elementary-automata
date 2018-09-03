package rules

var Fifty Rule = Rule{
	false,
	false,
	true,
	true,
	false,
	false,
	true,
	false,
}

func init() {
	register(50, Fifty)
}
