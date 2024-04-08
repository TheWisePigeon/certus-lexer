package tokens

type TokenKind string

type Token struct {
	Kind    TokenKind
	Literal string
}

const (
	LEFT_BRACE    = "{"
	RIGHT_BRACE   = "}"
	HASHTAG       = "#"
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"
	GET           = "GET"
	POST          = "POST"
	PUT           = "PUT"
	DELETE        = "DELETE"
	HTTP          = "HTTP"
	HEADERS       = "HEADERS"
	KEY           = "KEY"
	VALUE         = "VALUE"
	COLON         = ":"
	QUOTE         = "\""
	ILLEGAL       = "ILLEGAL"
	EOF           = "EOF"
)
