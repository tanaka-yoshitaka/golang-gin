package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func SetLaguageMessage(l []language.Tag) gin.HandlerFunc {
	return func(c *gin.Context) {
		t, _, _ := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
		matcher := language.NewMatcher(l)
		tag, _, _ := matcher.Match(t...)
		p := message.NewPrinter(tag)
		c.Set("language_message", p)
		c.Next()
	}
}
