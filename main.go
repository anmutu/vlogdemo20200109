/*
  author='du'
  date='2020/1/9 16:25'
*/
package main

import "net/http"


func main() {
	//第二步，注册，分发
	http.HandleFunc("/sayHello",sayHello)
	//第三步，启动web服务
	http.ListenAndServe("8090",nil)
}

func sayHello(w http.ResponseWriter,r *http.Request)  {
	//第一步，写方法
	w.Write([]byte("hello world"))
}



