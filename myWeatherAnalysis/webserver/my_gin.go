package webserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"log"
	"myWeatherAnalysis/util"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func AddRouterForLogseqJournal(r *gin.Engine) {
	r.GET("/logseq", func(c *gin.Context) {
		content, err := util.GenJournalContentOfLastThreeDays(util.LogseqBaseDirPath)
		if err != nil {
			log.Panicln(err)
		}
		//c.String(http.StatusOK, "%v", content)
		//c.Data(http.StatusOK, "text/markdown; charset=utf-8", []byte(content))
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		p := parser.NewWithExtensions(extensions)

		htmlFlags := html.CommonFlags | html.TOC | html.LazyLoadImages
		r := html.NewRenderer(html.RendererOptions{Flags: htmlFlags})

		htmlContent := markdown.ToHTML([]byte(content), p, r)
		preStyle := "font-family: inherit; font-size: inherit; line-height: inherit; white-space: pre-wrap; word-wrap: break-word;"
		htmlContent = []byte(fmt.Sprintf("<pre style=\"%s\">%s</pre>", preStyle, string(htmlContent)))
		//htmlContent = []byte(strings.ReplaceAll(string(htmlContent), " ", "&nbsp;"))

		fmt.Println(string(htmlContent))

		c.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
	})
}

func AddRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/r1", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"键":   "值",
			"<M>": "?!&$",
		})
	})
	r.GET("/param/:param", func(c *gin.Context) {
		value := c.Param("param")
		//fmt.Printf("Param value is: %v\n", value)
		c.String(http.StatusOK, "Param value is: %v\n", value)
		c.String(http.StatusOK, "%v", c.FullPath())
	})
	r.GET("/query", func(c *gin.Context) {
		value := c.Query("key")
		c.String(http.StatusOK, "Query value is: '%v'\n", value)
	})
	r.GET("/json-return-test", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, []any{"a", []string{"b1", "b2"}, map[string]any{
			"c1": 1,
			"c2": util.S1{
				Name: "s1",
				CounterPart: util.S2{
					Name: "s2",
				},
			},
		}})
	})
}

//func addRouterPost(r *gin.Engine) {
//	r.GET("/post", func(c *gin.Context) {
//		c.ShouldBindUri()
//	})
//}

//func addRouterOther(r *gin.Engine) {
//	r.GET("/other", func(c *gin.Context) {
//		func(a bool) {}
//	})
//}
