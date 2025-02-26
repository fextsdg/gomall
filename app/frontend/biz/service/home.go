package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "gomall/app/frontend/hertz_gen/frontend/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (res map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//业务逻辑--返回前端页面动态数据
	res = make(map[string]any)
	res["Name"] = "Hot Sole" //网站名称
	items := []map[string]any{
		{
			"Name":  "华为智选优畅享50Plus 仅￥9.99包邮啦！",
			"Image": "/static/images/phone.avif",
			"title": "华为智选优畅享50Plus 5G全网通 6400万AI影像 8GB+128GB海雾蓝【充电套装】",
		},
		{
			"Name":  "python编程 ￥108.00！",
			"Image": "/static/images/book.avif",
			"title": "python编程从入门三剑客Python编程从入门到实践+快速上手+极客编程+编程实战 （第3版） 三剑客全套4册",
		},
		{
			"Name":  "秋冬新款长袖t恤 ￥12.90！",
			"Image": "/static/images/Tshirt.avif",
			"title": "秋冬新款长袖t恤女修身显瘦打底衫内搭上衣高级感欧货 愚仁1045白色 L",
		},
		{
			"Name":  "哪吒元素手表 ￥9.90另付邮费！ ",
			"Image": "/static/images/watch.avif",
			"title": "哪吒元素手表青少年学生休闲时尚石英手表儿童礼物国潮手表 魔童降世 40mm",
		},
		{
			"Name":  "华为智选优畅享50Plus 仅￥9.99包邮啦！",
			"Image": "/static/images/phone.avif",
			"title": "华为智选优畅享50Plus 5G全网通 6400万AI影像 8GB+128GB海雾蓝【充电套装】",
		},
		{
			"Name":  "python编程 ￥108.00！",
			"Image": "/static/images/book.avif",
			"title": "python编程从入门三剑客Python编程从入门到实践+快速上手+极客编程+编程实战 （第3版） 三剑客全套4册",
		},
		{
			"Name":  "秋冬新款长袖t恤 ￥12.90！",
			"Image": "/static/images/Tshirt.avif",
			"title": "秋冬新款长袖t恤女修身显瘦打底衫内搭上衣高级感欧货 愚仁1045白色 L",
		},
		{
			"Name":  "哪吒元素手表 ￥9.90另付邮费！ ",
			"Image": "/static/images/watch.avif",
			"title": "哪吒元素手表青少年学生休闲时尚石英手表儿童礼物国潮手表 魔童降世 40mm",
		},
	}
	res["Items"] = items

	return res, nil
}
