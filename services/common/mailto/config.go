package mailto

var (
	defaultConfig = Config{To: []string{"mail@example.com"}, InnerText: "Click Me!"}
)

// Config contains parameters to create a mailto tag
type Config struct {
	To        []string
	Cc        []string
	Bcc       []string
	Subject   string
	Body      string
	InnerText string
}

// ReplyConfig returns a config, prepending "Re: " onto the ReplyingToSubject (the subject of the mail you are replying to)
// This causes gmail to group conversations, preserving threading.
// Still need to test with other clients, more of a workaround for In-Reply-To or References headers (RFC 6068 & 5322)
func ReplyConfig(To, Cc, Bcc []string, ReplyingToSubject, Body, InnerText string) Config {

	return Config{
		To:        To,
		Cc:        Cc,
		Bcc:       Bcc,
		Subject:   "Re: " + ReplyingToSubject,
		Body:      Body,
		InnerText: InnerText,
	}
}
