package middleware

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func MyCustomLogger(l []language.Tag) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("-------MyCustomLogger Start------")

		t, _, _ := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
		matcher := language.NewMatcher(l)
		tag, _, _ := matcher.Match(t...)
		p := message.NewPrinter(tag)
		fmt.Println(reflect.TypeOf(p))
		c.Set("lang", p)
		// fmt.Println(p.Sprintf("bread"))

		c.Next()
		fmt.Println("-------MyCustomLogger End------")
	}
}
