package Goods

import (
	"NFUShop/DbModel"
	"NFUShop/Result"
)

func GetGoodsSell(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += v.Sell
		}
	}
	return ret
}

func GetGoodsStoke(goodsId int) int {
	var ret int
	if ok, data := DbModel.SelectSubGoodsByGoodsId(goodsId); ok {
		for _, v := range data {
			ret += *v.Stoke
		}
	}
	return ret
}

/**
 * @Description: 搜索
 * @param title
 * @param limit
 * @param offset
 * @return Result.Result
 */
func SearchGoodsByTitle(title string, limit int, offset int) Result.Result {
	var ret Result.Result
	ret.Code = Result.UnKnow
	type name struct {
		DbModel.Goods
		SubGoods []int `json:"sub_goods"`
	}
	var retData []name
	if ok, data := DbModel.SelectGoodsLikeTitle(title, limit, offset); ok {
		var tmp name
		for _, v := range data {
			tmp.Goods = v
			var subGoodsSet []int
			if ok2, data2 := DbModel.SelectSubGoodsByGoodsId(v.Id); ok2 {
				for _, v2 := range data2 {
					subGoodsSet = append(subGoodsSet, v2.Id)
				}
			}
			tmp.SubGoods = subGoodsSet
			retData = append(retData, tmp)
		}

		ret.Code = Result.Ok
		ret.Data = retData
	}
	return ret
}