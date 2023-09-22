package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Request() {
	response, err := http.Get("https://m.zhipin.com/wapi/moment/salaryQuery/analysis?positionCode=100116&cityCode=101010100&selfSalary=0&_t=1694616879501")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()
	b, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println("read body fail")
		return
	}
	fmt.Printf("body content: [%s]\n", string(b))
}

// GetPositionConfigList 获取positionConfigList
func GetPositionConfigList() {

}

// GetCityJson 获取city-json
func GetCityJson(processFunc func([]map[string]interface{})) {
	//now := time.Now().Unix()
	//// 将时间戳转换为字符串
	//timestampStr := strconv.FormatInt(now, 10) // 10 表示十进制表示
	//cityUrl := "https://m.zhipin.com/wapi/zpboss/h5/marketpay/newCitys.json?_t=" + string(timestampStr)
	//log.Println(cityUrl)
	//response, err := http.Get(cityUrl)
	//if err != nil {
	//	panic("request error")
	//}
	//defer response.Body.Close()
	//b, error := io.ReadAll(response.Body)
	//if error != nil {
	//	panic("body error")
	//}
	//// 定义一个map，用来存储解析后的json数据
	//var data map[string]interface{}
	//if err := json.Unmarshal(b, &data); err != nil {
	//	panic("json解析失败")
	//}
	data := []map[string]interface{}{
		{
			"name": "一线城市",
			"subLevelModelList": []map[string]interface{}{
				{
					"name":      "北京",
					"levelType": 0,
					"code":      101010100,
				},
				{
					"name":      "上海",
					"levelType": 0,
					"code":      101020100,
				},
				{
					"name":      "广州",
					"levelType": 0,
					"code":      101280100,
				},
				{
					"name":      "深圳",
					"levelType": 0,
					"code":      101280600,
				},
			},
			"code": 1,
		},
		{
			"name": "新一线城市",
			"subLevelModelList": []map[string]interface{}{
				{
					"name":      "杭州",
					"levelType": 1,
					"code":      101210100,
				},
				{
					"name":      "成都",
					"levelType": 1,
					"code":      101270100,
				},
				{
					"name":      "长沙",
					"levelType": 1,
					"code":      101250100,
				},
				{
					"name":      "武汉",
					"levelType": 1,
					"code":      101200100,
				},
				{
					"name":      "南京",
					"levelType": 1,
					"code":      101190100,
				},
				{
					"name":      "郑州",
					"levelType": 1,
					"code":      101180100,
				},
				{
					"name":      "沈阳",
					"levelType": 1,
					"code":      101070100,
				},
				{
					"name":      "西安",
					"levelType": 1,
					"code":      101110100,
				},
				{
					"name":      "天津",
					"levelType": 1,
					"code":      101030100,
				},
				{
					"name":      "青岛",
					"levelType": 1,
					"code":      101120200,
				},
				{
					"name":      "重庆",
					"levelType": 1,
					"code":      101040100,
				},
				{
					"name":      "苏州",
					"levelType": 1,
					"code":      101190400,
				},
				{
					"name":      "东莞",
					"levelType": 1,
					"code":      101281600,
				},
				{
					"name":      "合肥",
					"levelType": 1,
					"code":      101220100,
				},
				{
					"name":      "佛山",
					"levelType": 1,
					"code":      101280800,
				},
				{
					"name":      "宁波",
					"levelType": 1,
					"code":      101210400,
				},
			},
			"code": 2,
		},
		{
			"name": "二线城市",
			"subLevelModelList": []map[string]interface{}{
				{
					"name":      "贵阳",
					"levelType": 2,
					"code":      101260100,
				},
				{
					"name":      "南昌",
					"levelType": 2,
					"code":      101240100,
				},
				{
					"name":      "长春",
					"levelType": 2,
					"code":      101060100,
				},
				{
					"name":      "福州",
					"levelType": 2,
					"code":      101230100,
				},
				{
					"name":      "哈尔滨",
					"levelType": 2,
					"code":      101050100,
				},
				{
					"name":      "石家庄",
					"levelType": 2,
					"code":      101090100,
				},
				{
					"name":      "南宁",
					"levelType": 2,
					"code":      101300100,
				},
				{
					"name":      "太原",
					"levelType": 2,
					"code":      101100100,
				},
				{
					"name":      "兰州",
					"levelType": 2,
					"code":      101160100,
				},
				{
					"name":      "济南",
					"levelType": 2,
					"code":      101120100,
				},
				{
					"name":      "大连",
					"levelType": 2,
					"code":      101070200,
				},
				{
					"name":      "无锡",
					"levelType": 2,
					"code":      101190200,
				},
				{
					"name":      "厦门",
					"levelType": 2,
					"code":      101230200,
				},
				{
					"name":      "温州",
					"levelType": 2,
					"code":      101210700,
				},
				{
					"name":      "徐州",
					"levelType": 2,
					"code":      101190800,
				},
				{
					"name":      "嘉兴",
					"levelType": 2,
					"code":      101210300,
				},
				{
					"name":      "珠海",
					"levelType": 2,
					"code":      101280700,
				},
				{
					"name":      "泉州",
					"levelType": 2,
					"code":      101230500,
				},
				{
					"name":      "常州",
					"levelType": 2,
					"code":      101191100,
				},
				{
					"name":      "绍兴",
					"levelType": 2,
					"code":      101210500,
				},
				{
					"name":      "昆明",
					"levelType": 2,
					"code":      101290100,
				},
				{
					"name":      "南通",
					"levelType": 2,
					"code":      101190500,
				},
				{
					"name":      "烟台",
					"levelType": 2,
					"code":      101120500,
				},
				{
					"name":      "金华",
					"levelType": 2,
					"code":      101210900,
				},
				{
					"name":      "台州",
					"levelType": 2,
					"code":      101210600,
				},
				{
					"name":      "惠州",
					"levelType": 2,
					"code":      101280300,
				},
				{
					"name":      "中山",
					"levelType": 2,
					"code":      101281700,
				},
				{
					"name":      "保定",
					"levelType": 2,
					"code":      101090200,
				},
				{
					"name":      "廊坊",
					"levelType": 2,
					"code":      101090600,
				},
			},
			"code": 3,
		},
	}

	processFunc(data)

}

func ReadLocalPosition(callBack func(m []map[string]interface{})) {
	content, err := ioutil.ReadFile("./datasource/stream-response.txt")
	if err != nil {
		log.Println("read file error")
		return
	}
	jsonMap := map[string]interface{}{}
	err = json.Unmarshal([]byte(content), &jsonMap)
	if err != nil {
		log.Println("json load fail")
		return
	}
	zpData := jsonMap["zpData"].([]interface{})
	var convertedSlice []map[string]interface{}
	for _, item := range zpData {
		if m, ok := item.(map[string]interface{}); ok {
			convertedSlice = append(convertedSlice, m)
		}
	}
	if len(zpData) > 0 {
		callBack(convertedSlice)
	} else {
		callBack(convertedSlice)
	}

}
