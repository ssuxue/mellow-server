package initData

import (
	"fmt"
	"github.com/casbin/casbin"
)

// Web应用，可以将路由作为obj，请求方式Method作为act
// 登录用户的角色作为sub，在每一次请求时，把这3个参数传递给e.Enforce
// 就可以实现对Web页面和请求接口的权限控制管理
// TODO 这个最后在解决
func SetCasbin() {
	//通过策略文件和模型配置穿件一个casbin访问控制实例
	e := casbin.NewEnforcer("static/conf/rbac_model.conf", "rbac_policy.csv")

	//定义各种sub，obj和act的数组
	subs := []string{"bob", "zeta"}
	objs := []string{"data1", "data2"}
	acts := []string{"read", "write"}

	//遍历组合sub，obj，act并打印出对应策略匹配结果。
	for _, sub := range subs {
		for _, obj := range objs {
			for _, act := range acts {
				fmt.Println(sub, "/", obj, "/", act, "=", e.Enforce(sub, obj, act))
			}
		}
	}
}
