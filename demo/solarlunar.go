package demo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/nosixtools/solarlunar"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name           string //姓名
	Birthday_Solar string //阳历生日
	Birthday_Lunar string //阴历生日
}

/*
作用：主要用于阴历与阳历之间的转换

	参数：第一个参数为想要转换的日期，格式为：2023-04-09；第二个参数为：阳历转阴历，还是阴历转阳历，1:阳历转阴历，其它为阴历转阳历。
	如果不输入参数，那么获取当前日期进行转换，默认将阳历转换为阴历；
	如果只输入一个参数，即日期，那么就是将指定的日期默认转换为阴历

author：tanglong
date：2023-04-10
*/
func SolarlunarSwitch(dateStr string, flag string) (switchResult string) {
	tmp := ""
	/*
		d := ""
		n := ""
		tmp := ""

		if len(dateStr) == 0 {
			// 获取当前时间
			d = time.Now().Format("2006-01-02")
			// 默认将阳历转换为阴历
			n = "1"
		} else if len(flag) == 0 {
			d = dateStr
			n = "1"
		} else {
			d = dateStr
			n = flag
		}
	*/
	// 字符串转换为整数
	num, _ := strconv.Atoi(flag)
	// 校验输入的日期字符串格式，比如：2023-4-1，会自动转换为：2023-04-01
	dd := DateFormatCheck(dateStr)
	if num == 1 {
		switchResult = fmt.Sprintf("%s(%s)", solarlunar.SolarToSimpleLuanr(dd), solarlunar.SolarToChineseLuanr(dd))
		// fmt.Printf("数据类型为： %T\n", d)
		tmp = fmt.Sprintf("[%s]通过阳历转换成阴历为: %s", dateStr, switchResult)
	} else {
		switchResult = fmt.Sprintf("%s", solarlunar.LunarToSolar(dd, false))
		tmp = fmt.Sprintf("[%s]通过阴历转换成阳历为：%s", dateStr, switchResult)
	}
	fmt.Printf("tmp: %v\n", tmp)
	return switchResult
}

/*
作用：校验输入的日期字符串格式，比如：2023-4-1，会自动转换为：2023-04-01
author：tanglong
date：2023-04-10
*/
func DateFormatCheck(dstr string) (ddstr string) {
	str := "-"
	// 匹配更多的日期格式，比如：2023-4-9，自动转换成2023-04-09
	for _, line := range strings.Split(string(dstr), str) {
		if len(line) < 2 {
			// 不够2位，在前面补0
			line = fmt.Sprintf("%0*s", 2, line)
			// fmt.Printf("line: %v\n", line)
		}
		// 拼接字符串与符号
		ddstr += line + str
		// fmt.Printf("line123: %v\n", line)
		// fmt.Printf("d111: %v\n", dd)
	}
	// 去掉右边的“-”字符串
	ddstr = strings.TrimRight(ddstr, "-")
	return ddstr
}

/*
操作redis，将转换后的日期写入redis
*/
func SetRedis(name string, dataStr1 string, flag string) (dataStr string) {
	date_Solar := ""
	date_lunar := ""
	dataStr = SolarlunarSwitch(dataStr1, flag)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:9736", // Redis 服务器地址
		Password: "YunJia#NJ@2017", // 密码，如果有的话
		DB:       0,                // 使用默认的数据库
		PoolSize: 20,
	})

	if flag == "1" {
		date_Solar = dataStr1
		date_lunar = dataStr

	} else {
		date_Solar = dataStr
		date_lunar = dataStr1
	}
	person := Person{
		Name:           name,
		Birthday_Solar: date_Solar,
		Birthday_Lunar: date_lunar,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("JSON编码失败：", err)
		return
	}

	// 设置键值对到 Redis 中，键为 "myKey"，值为 myString
	err = client.Set("person", jsonData, time.Hour).Err()

	if err != nil {
		fmt.Println("Error writing to Redis:", err)
		return
	}

	fmt.Println("String written to Redis successfully!")

	value := client.Get("person")
	fmt.Println("my Key is:", value)

	// 关闭 Redis 连接
	defer client.Close()

	return dataStr
}

/*
从页面读取数据并返回给页面展示
*/
func HandleConvert(c *gin.Context) {
	date_str := ""
	// 从表单中获取日期
	name := c.PostForm("name")
	dateStr := c.PostForm("date")
	flagStr := c.PostForm("flag")

	SolarlunarSwitch(dateStr, flagStr)
	date_str = SetRedis(name, dateStr, flagStr)
	// 返回结果
	if flagStr == "1" {
		c.JSON(http.StatusOK, gin.H{"姓名": name, "农历日期": date_str, "阳历日期": dateStr})
	} else {
		c.JSON(http.StatusOK, gin.H{"姓名": name, "农历日期": dateStr, "阳历日期": date_str})
	}
}
