/*
 * @Author: lisheng
 * @Date: 2022-10-14 23:34:41
 * @LastEditTime: 2022-11-18 09:57:18
 * @LastEditors: lisheng
 * @Description: 日志模块
 * @FilePath: /gitee.com/liqiyuworks/jf-go-kit/base/log.go
 */
package base

import (
	"fmt"
	"runtime"

	"gitee.com/liqiyuworks/jf-go-kit/config"

	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	Glog = logrus.New()
)

type output struct {
	fileOut io.Writer
	stdout  io.Writer
}

func (o *output) Write(p []byte) (n int, err error) {
	n, err = o.fileOut.Write(p)
	if err != nil {
		return n, err
	}

	return o.stdout.Write(p)
}

func CreateLogFile() {
	logFilePath := config.C.Log.Path
	logFileName := config.C.Log.Name
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(logFilePath, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logFilePath, err))
		}
	}

	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			Glog.Infoln(">>> create run.log failed!")
		}
		defer func() {
			// 2.关闭文件
			file.Close()
			Glog.Infoln(">>> create run.log ok and close!")
		}()
	}
}

func InitLogger() {
	logFilePath := config.C.Log.Path
	logFileName := config.C.Log.Name
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		Glog.Infoln("err", err)
	}

	//设置输出
	Glog.Out = src
	// RunLog.SetOutput(src)

	// 设置文件和行号
	Glog.SetReportCaller(true)

	//设置日志级别
	if config.C.Log.Level == "" || config.C.Log.Level == "error" {
		Glog.SetLevel(logrus.ErrorLevel)
	} else if config.C.Log.Level == "debug" {
		Glog.SetLevel(logrus.DebugLevel)
	}

	//设置日志格式
	Glog.SetFormatter(&logrus.TextFormatter{ //以下设置只是为了使输出更美观
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名
			fileName := path.Base(frame.File)
			return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	out := &output{
		fileOut: logWriter,
		stdout:  os.Stdout,
	}

	Glog.SetOutput(out)
}
