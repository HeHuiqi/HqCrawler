package controller

import (
	"HqCrawler/web/view"
	"github.com/olivere/elastic"
	"net/http"
	"strings"
	"strconv"
	"HqCrawler/web/model"
	"context"
	"reflect"
	"HqCrawler/engine"
	"regexp"
	"fmt"
)

type SearchResultHandler struct {
	 view view.SearchResultView
	 client *elastic.Client

}

func NewSearchResultHandler(template string) SearchResultHandler  {
	client,err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil  {
		panic(err)
	}
	return SearchResultHandler{
		view:view.CreateSearchResultView(template),
		client:client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// 例子 query= 女 Payload.Age:(<30)查询参数条件已空格隔开
	q := strings.TrimSpace(req.FormValue("q"))
	//fmt.Println("q==",q)

	from ,err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w,"q=%s,from=%d",q,from)

	var page model.SearchResult
	page,err = h.getSearchResult(q,from)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err =  h.view.Render(w,page)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}
func (h SearchResultHandler)getSearchResult(q string,from int) (model.SearchResult,error)  {


	var result model.SearchResult
	result.Query = q
	query := rewriteQueryString(q)
	fmt.Println("query=",query)
	resp,err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(query)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result,err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result,nil
}

func rewriteQueryString(q string) string  {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return  re.ReplaceAllString(q,"Payload.$1:")
}
