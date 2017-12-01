package tag

import (
	"github.com/rogeecn/gostriptags"
)

type Rule struct {
	EscapeInValid bool                   // escape invalid tag
	TrimSpace     bool                   // trim html space
	ValidTags     map[string]interface{} // default validation tag, see also this.Init method
	ValidAttrs    map[string]bool        // default validation attribute
	DisableAttrs  map[string]bool        // default disabled attribute
}

func Strip(rawHtml string, rule Rule) (string, error) {
	stripTags := striptags.NewStripTags()
	stripTags.EscapeInValid = rule.EscapeInValid
	stripTags.TrimSpace = rule.TrimSpace
	stripTags.ValidAttrs = rule.ValidAttrs
	stripTags.DisableAttrs = rule.DisableAttrs

	return stripTags.Fetch(rawHtml)
}
