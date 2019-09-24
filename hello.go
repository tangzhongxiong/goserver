package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// fmt.Println("Hello,World")
	// fmt.Println("Hello" + " World")

	// var flag bool
	// flag = true
	// bar := false
	// fmt.Println(flag, bar)

	// var a, b int = 1, 2
	// fmt.Println(a, b)

	// var c string
	// fmt.Println(c)

	// var d *int
	// fmt.Println(d)

	// var q, w, e = 1, 2, 3
	// fmt.Println(q, w, e)

	// const length int = 10
	// const height int = 10
	// fmt.Println("area is %d", length*height)

	// const (
	// 	Unknown = 0
	// 	Female  = 1
	// 	Male    = 2
	// )

	// fmt.Println(Unknown)

	// var foo int = 10
	// var ptr *int
	// ptr = &foo
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	// fmt.Println(max(3, 2))

	// str1, str2 := swap("A", "B")
	// fmt.Println(str1, str2)

	// var arr1 [3]int
	// arr1[1] = 10
	// arr1[2] = 20
	// var arr2 = [2]int{1, 2}
	// var arr3 = [...]int{1, 2, 3}

	// fmt.Println(arr1[0])
	// fmt.Println(arr1[1])
	// fmt.Println(arr2[1])
	// fmt.Println(arr3[2])

	// var book1 Books
	// book1.title = "大问题"
	// fmt.Println(book1.title)

	// var book2 Books = Books{"大问题", "au"}
	// fmt.Println(book2.author)

	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

type Books struct {
	title  string
	author string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}
