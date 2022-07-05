package framework

import (
	"net/http"
	"strings"
)

// Core represent core struct
//两级路由 一级get,post,put,delet；二级url路径
type Core struct {
	router map[string]map[string]ControllerHandler
}

//func NewCore() *Core {
//	return &Core{router: map[string]ControllerHandler{}}
//}
//
//func (c *Core) Get(url string, handler ControllerHandler) {
//	c.router[url] = handler
//}
//
//func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
//	log.Println("core.serveHTTP")
//	ctx := NewContext(request, response)
//
//	// 一个简单的路由选择器，这里直接写死为测试路由foo
//	router := c.router["foo"]
//	if router == nil {
//		return
//	}
//	log.Println("core.router")
//
//	router(ctx)
//
//}
//初始化框架核心结构
func NewCore() *Core {
	//定义二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	// 将二级map写入一级map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter
	return &Core{router: router}
}

//路由注册 URL    全部转换为大写  大小写不敏感
//对应method = GET
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

//对应Method = POST
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler

}

///对应Method = PUT
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

//对应Method =DELETE
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

//匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// url 和method全部转化为大写，保证大小写不敏感
	url := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUrl := strings.ToUpper(url)
	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		// 查找第二层map
		if handler, ok := methodHandlers[upperUrl]; ok {
			return handler
		}
	}
	return nil
}
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//封装自定义context
	ctx := NewContext(r, w)
	//寻找路由
	router := c.FindRouteByRequest(r)
	if router == nil {
		//如果没有找到，在这里打印日志
		ctx.Json(404, "not found")
		return
	}
	//调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
