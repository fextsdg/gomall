package service

import (
	"context"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//业务逻辑--返回前端页面动态数据
	res = make(map[string]any)
	res["Name"] = "Hot Sole" //网站名称
	//items := []map[string]any{
	//	{
	//		"Id":          1,
	//		"Name":        "华为智选优畅享50Plus 仅￥9.99包邮啦！",
	//		"Picture":     "/static/images/phone.avif",
	//		"Description": "华为智选优畅享50Plus 5G全网通 6400万AI影像 8GB+128GB海雾蓝【充电套装】",
	//		"Price":       "9.99",
	//	},
	//	{
	//		"Id":          2,
	//		"Name":        "python编程 ￥108.00！",
	//		"Picture":     "/static/images/book.avif",
	//		"Description": "python编程从入门三剑客Python编程从入门到实践+快速上手+极客编程+编程实战 （第3版） 三剑客全套4册",
	//		"Price":       "108.00",
	//	},
	//	{
	//		"Id":          3,
	//		"Name":        "秋冬新款长袖t恤 ￥12.90！",
	//		"Picture":     "/static/images/Tshirt.avif",
	//		"Description": "秋冬新款长袖t恤",
	//		"Price":       "12.90",
	//	},
	//	{
	//		"Id":          4,
	//		"Name":        "哪吒元素手表 ￥9.90另付邮费！ ",
	//		"Picture":     "/static/images/watch.avif",
	//		"Description": "哪吒元素手表青少年学生休闲时尚石英手表儿童礼物国潮手表 魔童降世 40mm",
	//		"Price":       "9.90",
	//	},
	//	{
	//		"Id":          5,
	//		"Name":        "华为智选优畅享50Plus 仅￥9.99包邮啦！",
	//		"Picture":     "/static/images/phone.avif",
	//		"Description": "华为智选优畅享50Plus 5G全网通 6400万AI影像 8GB+128GB海雾蓝【充电套装】",
	//		"Price":       "9.99",
	//	},
	//	{
	//		"Id":          6,
	//		"Name":        "python编程 ￥108.00！",
	//		"Picture":     "/static/images/book.avif",
	//		"Description": "python编程从入门三剑客Python编程从入门到实践+快速上手+极客编程+编程实战 （第3版） 三剑客全套4册",
	//		"Price":       "108.00",
	//	},
	//	{
	//		"Id":          7,
	//		"Name":        "秋冬新款长袖t恤 ￥12.90！",
	//		"Picture":     "/static/images/Tshirt.avif",
	//		"Description": "秋冬新款长袖t恤女修身显瘦打底衫内搭上衣高级感欧货 愚仁1045白色 L",
	//		"Price":       "12.90",
	//	},
	//	{
	//		"Id":          8,
	//		"Name":        "哪吒元素手表 ￥9.90另付邮费！ ",
	//		"Picture":     "/static/images/watch.avif",
	//		"Description": "哪吒元素手表青少年学生休闲时尚石英手表儿童礼物国潮手表 魔童降世 40mm",
	//		"Price":       "9.90",
	//	},
	//}
	var items []map[string]any
	result, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	products := result.GetProducts()
	for _, p := range products {
		items = append(items, map[string]any{
			"Id":          p.GetId(),
			"Name":        p.Name,
			"Picture":     p.Picture,
			"Description": p.Description,
			"Price":       p.Price,
		})
	}

	res["Items"] = items
	//res["user_id"] = 123
	var cartNum int
	res["cart_num"] = cartNum
	return res, nil
}
