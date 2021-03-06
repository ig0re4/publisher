package blackfriday

import (
	"bufio"
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
)

// Whitelisted element tags, attributes on particular tags, attributes that are
// interpreted as protocols (again on particular tags), and allowed protocols.
var (
	whitelistTags      map[string]bool
	whitelistAttrs     map[string]map[string]bool
	protocolAttrs      map[string]map[string]bool
	whitelistProtocols [][]byte
)

func init() {
	whitelistTags = toSet([]string{
		"a", "b", "blockquote", "br", "caption", "cite", "code", "col",
		"colgroup", "dd", "div", "dl", "dt", "em",
		"h1", "h2", "h3", "h4", "h5", "h6",
		"i", "img", "li", "ol", "p", "pre", "q", "small", "strike", "strong",
		"sub", "sup", "table", "tbody", "td", "tfoot", "th", "thead", "tr", "u",
		"ul"})
	whitelistAttrs = map[string]map[string]bool{
		"a":   toSet([]string{"href", "title", "rel"}),
		"img": toSet([]string{"src", "alt", "title"}),
	}
	protocolAttrs = map[string]map[string]bool{
		"a":   toSet([]string{"href"}),
		"img": toSet([]string{"src"}),
	}
	whitelistProtocols = [][]byte{
		[]byte("http://"),
		[]byte("https://"),
		[]byte("ftp://"),
		[]byte("mailto:"),
	}
}

func toSet(keys []string) map[string]bool {
	m := make(map[string]bool, len(keys))
	for _, k := range keys {
		m[k] = true
	}
	return m
}

// Sanitizes the given input by parsing it as HTML5, then whitelisting known to
// be safe elements and attributes. All other HTML is escaped, unsafe attributes
// are stripped.
func sanitizeHtmlSafe(input []byte) []byte {
	r := bytes.NewReader(input)
	var w bytes.Buffer
	tokenizer := html.NewTokenizer(r)
	wr := bufio.NewWriter(&w)

	// Iterate through all tokens in the input stream and sanitize them.
	for t := tokenizer.Next(); t != html.ErrorToken; t = tokenizer.Next() {
		switch t {
		case html.TextToken:
			// Text is written escaped.
			wr.WriteString(tokenizer.Token().String())
		case html.SelfClosingTagToken, html.StartTagToken:
			// HTML tags are escaped unless whitelisted.
			tag, hasAttributes := tokenizer.TagName()
			tagName := string(tag)
			if whitelistTags[tagName] {
				wr.WriteString("<")
				wr.Write(tag)
				for hasAttributes {
					var key, val []byte
					key, val, hasAttributes = tokenizer.TagAttr()
					attrName := string(key)
					// Only include whitelisted attributes for the given tagName.
					tagWhitelistedAttrs, ok := whitelistAttrs[tagName]
					if ok && tagWhitelistedAttrs[attrName] {
						// For whitelisted attributes, if it's an attribute that requires
						// protocol checking, do so and strip it if it's not known to be safe.
						tagProtocolAttrs, ok := protocolAttrs[tagName]
						if ok && tagProtocolAttrs[attrName] {
							if !protocolAllowed(val) {
								continue
							}
						}
						wr.WriteByte(' ')
						wr.Write(key)
						wr.WriteString(`="`)
						wr.WriteString(html.EscapeString(string(val)))
						wr.WriteByte('"')
					}
				}
				wr.WriteString(">")
			} else {
				wr.WriteString(html.EscapeString(string(tokenizer.Raw())))
			}
		case html.EndTagToken:
			// Whitelisted tokens can be written in raw.
			tag, _ := tokenizer.TagName()
			if whitelistTags[string(tag)] {
				wr.Write(tokenizer.Raw())
			} else {
				wr.WriteString(html.EscapeString(string(tokenizer.Raw())))
			}
		case html.CommentToken:
			// Comments are not really expected, but harmless.
			wr.Write(tokenizer.Raw())
		case html.DoctypeToken:
			// Escape DOCTYPES, entities etc can be dangerous
			wr.WriteString(html.EscapeString(string(tokenizer.Raw())))
		default:
			tokenizer.Token()
			panic(fmt.Errorf("Unexpected token type %v", t))
		}
	}
	err := tokenizer.Err()
	if err != nil && err != io.EOF {
		panic(tokenizer.Err())
	}
	wr.Flush()
	return w.Bytes()
}

func protocolAllowed(attr []byte) bool {
	for _, prefix := range whitelistProtocols {
		if bytes.HasPrefix(attr, prefix) {
			return true
		}
	}
	return false
}
