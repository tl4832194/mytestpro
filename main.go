// 每个工程必须有且只有一个主包，且包名必须为main
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mytestpro/demo"
	"net/http"

	// 引用要调用的包，格式：工程名/包名；比如下面的mytestpro就为我们的工程名，test为包名

	"mytestpro/test"
)

// 每个工程有且只有一个主函数，且函数名必须为main，而且必须在main包下，这个包和函数为这个工程的唯一入口
func main() {
	// 这里调用test包下面的Hello函数，必须在前面先引用test包
	s := test.Hello()
	fmt.Printf("s: %v\n", s)

	// 统计输入的行数据，并统计重复的输入行
	// a := test.SoComan()
	// fmt.Printf("a: %v\n", a)

	// b := test.Dup()
	// fmt.Printf("b: %v\n", b)

	// 输入文件名，循环读取文件内容，统计重复的行
	/* 	c := test.Dup1()
	   	fmt.Printf("c: %v\n", c) */

	// 输入文件名，读取文件所有内容，再通过分割切片的方式统计文件中重复的行
	/* 	d := test.Dup2()
	   	fmt.Printf("d: %v\n", d) */

	// 画图
	/* 	rand.Seed(time.Now().UTC().UnixNano())
	   	test.Lissajous(os.Stdout) */

	// 获取单个url
	// test.GetUrl()

	// 获取多个url
	/*	start := time.Now()
		// 定义一个channel通道
		ch := make(chan string)
		// 循环获取每个参数
		for _, url := range os.Args[1:] {
			// 将每个参数传给通道执行GetUrlAll()函数
			go test.GetUrlAll(url, ch)
			// 接收GetUrlAll()函数通过ch通道传递过来的信息
			fmt.Println(<-ch)
		}
		//for range os.Args[1:] {
		//	fmt.Println(<-ch)
		//}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	*/

	//访问http://localhost:8000/test地址，会返回地址后面的信息，比如：URL.Path = "/test"
	/*	http.HandleFunc("/", test.Handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	*/

	/*
		作用：主要用于阴历与阳历之间的转换
		参数：第一个参数为想要转换的日期，格式为：2023-04-09；第二个参数为：阳历转阴历，还是阴历转阳历，1:阳历转阴历，其它为阴历转阳历。
		如果不输入参数，那么获取当前日期进行转换，默认将阳历转换为阴历；
		如果只输入一个参数，即日期，那么就是将制定的日期默认转换为阴历
	*/
	/* aa := os.Args[1:]
	d := ""
	n := ""
	// 判断输入的参数小于1个
	if len(aa) < 1 {
		// 获取当前时间
		d = time.Now().Format("2006-01-02")
		// 默认将阳历转换为阴历
		n = "1"
	} else if len(aa) < 2 {
		d = aa[0]
		n = "1"
	} else {
		d = aa[0]
		n = aa[1]
	}

	// 校验输入的日期字符串格式，比如：2023-4-1，会自动转换为：2023-04-01
	dd := demo.DateFormatCheck(d)
	// 字符串转换为整数
	num, _ := strconv.Atoi(n)
	// 开始转换为阴历或者阳历
	demo.Solarlunar(dd, num)
	*/

	/*	abc := ""
		abc = demo.SolarlunarSwitch()
		fmt.Printf("abc: %v\n", abc)
	*/

	// =========================== 【开始redis操作】 ===============================
	// 初始化 Redis 客户端
	/*	client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis 服务器地址
			Password: "",               // 密码，如果有的话
			DB:       0,                // 使用默认的数据库
			PoolSize: 20,
		})

		// 关闭 Redis 连接
		//defer client.Close()

		// 要写入 Redis 的字符串
		//myString := "Hello, Redis!"

		// 使用 context.Background() 创建上下文
		//ctx := context.Background()
		// 设置键值对到 Redis 中，键为 "myKey"，值为 myString
		err := client.Set("myKey", abc, time.Hour).Err()
		//err := client.set
		if err != nil {
			fmt.Println("Error writing to Redis:", err)
			return
		}

		fmt.Println("String written to Redis successfully!")

		value := client.Get("myKey")
		fmt.Println("my Key is:", value)

		// 关闭 Redis 连接
		defer client.Close()
		// =========================== 【结束redis操作】 ===============================
	*/

	//================================ 操作html页面开始 =================================
	// 初始化 Gin 路由
	router := gin.Default()

	// 设置模板路径
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLGlob("/Users/tanglong/work/go_workspace/mytestpro/templates/*")

	// 设置静态文件路径
	router.Static("/static", "static")

	// 定义路由
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/convert", demo.HandleConvert)

	// 启动服务器
	router.Run(":8081")
	//demo.SetRedis()

	//================================ 操作html页面结束 =================================

}
