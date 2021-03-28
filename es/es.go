package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v6" //这里使用的是版本5，最新的是6，有改动
	"reflect"
)

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

type TimeInfo struct {
	Time  string `json:"time"`
}

//初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/

//创建
func create() {

	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func delete() {

	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func update() {
	res, err := client.Update().
		Index("megacorp").
		Type("employee").
		Id("2").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("megacorp").Type("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

//搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	printEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	if res.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {

			var t Employee
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}

			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("megacorp").Type("employee").Query(boolQ).Do(context.Background())
	printEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)

}

//简单分页
func list(size,page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}

	res,err := client.Search("megacorp").
		Type("employee").
		Size(size).
		From((page-1)*size).
		Do(context.Background())
	printEmployee(res, err)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	//create()
	//delete()
	//update()
	//gets()
	//query()
	//list(10,1)

	//res,_ := client.Search("megacorp").
	//	Type("employee").
	//	Size(10).
	//	From(9999).
	//	Sort("age",true).
	//	Do(context.Background())
	//var typ Employee
	//for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
	//	t := item.(Employee)
	//	fmt.Printf("%#v\n", t)
	//}

	//for i:=20;i<30 ;i++  {
	//	itoa := strconv.Itoa(i)
	//	e1 := Employee{"Jane"+itoa, "Smith"+itoa, 32+i, "I like to collect rock albums"+itoa, []string{"music"+itoa}}
	//	put1, err := client.Index().
	//		Index("megacorp").
	//		Type("employee").
	//		Id(fmt.Sprintf("%d",time.Now().UnixNano())).
	//		BodyJson(e1).
	//		Do(context.Background())
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	//}

	res,_ := client.Search("megacorp").
		Type("employee").
		//Sort("age",true).
		Size(10000).
		Do(context.Background())
	var typ Employee
	eachs := res.Each(reflect.TypeOf(typ))
	for _, item := range eachs { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}

	//mu := make(map[string]interface{})
	//mu["front"] = "10.34.4.10"
	//mu["back"] = "10.34.4.11"
	//data := make(map[string]interface{})
	//data["terminal.client_ip"] = mu
	//
	//e1 := AssetChangesInfo{
	//	Time:       time.Now().Unix(),
	//	Person:     "xx",
	//	Group:      "adsf",
	//	HostName:   "aaasdf",
	//	ClientIp:   "10.34.4.11",
	//	Details: data,
	//}
	//_, err := client.Index().
	//	Index("mutation").
	//	Type("doc").
	//	BodyJson(e1).
	//	Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
}

type AssetChangesInfo struct {
	Time int64 `json:"time"`  //时间
	Person string `json:"person"` //责任人
	Group string `json:"group"`   //部门
	HostName string `json:"host_name"` //计算机名称
	ClientIp string `json:"client_ip"` //终端IP
	Details map[string]interface{} `json:"details"` //变更信息
}