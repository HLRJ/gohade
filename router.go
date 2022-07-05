package main

import (
	"gohade/framework"
)

//注册路由规则
func registerRouter(core *framework.Core) {
	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	//需求 HTTP方法+静态路由匹配
	core.Get("/user/login", UserLoginController)
	//需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Get("/list", SubjectListController)
	}
}
