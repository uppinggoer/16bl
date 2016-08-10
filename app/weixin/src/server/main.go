package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	. "global"
	"http"
	mw "http/middleware"
	. "util"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
}

var serverLogFile *os.File
var pid int

func init() {
	// 获取pid
	pid = os.Getpid()
}

func main() {
	// serverLog
	serverLogFileName := LOG_PATH + "/server/" + filepath.Base(os.Args[0]) + ".log." + time.Now().Format("20060102")
	var err error
	serverLogFile, err = os.OpenFile(serverLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer serverLogFile.Close()

	// 重定向标准输出
	// syscall.Dup2(int(serverLogFile.Fd()), int(os.Stderr.Fd()))
	// syscall.Dup2(int(serverLogFile.Fd()), int(os.Stdout.Fd()))

	startCallback()
	defer stopCallback()

	// 初始化框架
	e := echo.New()

	// 设置路由
	e.Static("/static/", BASE_PATH+"/static")
	e.File("/favicon.ico", BASE_PATH+"/static/img/16bl.ico")

	// 配置server的log
	var loggerConfig = middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "cost_time":${latency},` +
			`"cost_time_human":"${latency_human}","rx_bytes":${rx_bytes},` +
			`"tx_bytes":${tx_bytes}}` + "\n",
		Output: serverLogFile,
	}
	genralGroup := e.Group("", middleware.LoggerWithConfig(loggerConfig), mw.UserInfo())
	http.RegisterRoutes(genralGroup)

	// 默认监听 127.0.0.1：8080
	var serverConf = Conf("server")
	host := serverConf.MustValue("http.host", "127.0.0.1").(string)
	port := serverConf.MustValue("http.port", "8080").(string)

	// 注册框架
	std := standard.New(host + ":" + port)
	std.SetHandler(e)
	// 开启服务
	serverErr := gracehttp.Serve(std.Server)
	appendServerLog(serverErr)
}

// server 开始前的回调
func startCallback() {
	pidFileName := VAR_PATH + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
	// fmt.Print(pidFileName)

	ioutil.WriteFile(pidFileName, []byte(strconv.Itoa(pid)), 0755)

	startLog := "server " + strconv.Itoa(pid) + " started @ " + time.Now().Format("15:04:05")
	appendServerLog(startLog)
}

// server关闭后的回调
func stopCallback() {
	pidFileName := VAR_PATH + "/pid/" + filepath.Base(os.Args[0]) + ".pid"
	err := os.Remove(pidFileName)
	var stopLog string
	if nil != err {
		stopLog = "remove server err @ " + time.Now().Format("15:04:05") + err.Error()
	} else {
		stopLog = "server " + strconv.Itoa(pid) + " stopped @ " + time.Now().Format("15:04:05")
	}
	appendServerLog(stopLog)
}

// 追回serverlog  如server的启动、关闭,端口监听等等
func appendServerLog(logLine interface{}) {
	// fmt.Println(logLine)
	serverLogFile.WriteString(fmt.Sprintln(logLine))
}
