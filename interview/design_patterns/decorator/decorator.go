package decorator

import (
	"fmt"
	"net/http"
)

//装饰器模式：允许向一个现有的对象添加新的功能，同时又不改变其结构。
//这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。
//这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。
//我们使用最为频繁的场景就是http请求的处理：对http请求做cookie、header校验。

func AutoAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		//cookie, err := r.Cookie("Auth")
		//if err != nil || cookie.Value != "Authentic" {
		//	w.WriteHeader(http.StatusForbidden)
		//	return
		//}
		if headers.Get("Token") != "123" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}
