package main

import (
	"practice/crawler/engine"
	"practice/crawler/idoctor/parser"
	"practice/crawler/repository"
	"practice/crawler/scheduler"
)

//func main() {
//resp, err := http.Get("http://www.zhenai.com/zhenghun")
//if err!=nil {
//	panic(err)
//}
//defer  resp.Body.Close()
//
//if resp.StatusCode!=http.StatusOK {
//	fmt.Println("Error: status code ",resp.StatusCode)
//	return
//}
//
////encode:=determineEncoding(resp.Body)
////encodeReader:=transform.NewReader(resp.Body,
////	encode.NewDecoder())
////all, err := ioutil.ReadAll(encodeReader)
//all, err := ioutil.ReadAll(resp.Body)
//if err!=nil {
//	panic(err)
//}
//
//
//printCityList(all)

//}

////检测编码
//func determineEncoding(r io.Reader) encoding.Encoding  {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err!=nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}
//
////城市
//func  printCityList(contents []byte)  {
//	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//	//match:=re.FindAllString(text,-1)
//	matchs:=re.FindAllSubmatch(contents,-1)
//
//	for _,m:=range matchs{
//		fmt.Printf("City:%s  , URL:%s \n",m[2],m[1])
//	}
//
//	fmt.Printf("累计获取城市：%d ",len(matchs))
//}

func main() {
	itemChan, err := repository.ItemSave("idoctor")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		//Url:        "http://www.zhenai.com/zhenghun",
		//Url:        "http://www.daxuecn.com/chaxun/",
		//Url:        "https://cd.julive.com/project/s",
		Url:        "https://ysk.99.com.cn/department/all/",
		ParserFunc: parser.ParserKeshiList,
	})
}
