package SeasLog4Go

import (
	"container/list"
	"fmt"
	"log"
	"runtime"
	"sync"
)

//使用接口
/**
  Logger 的配置
   1、文件路径
   2、日期格式
*/
type LoggerConfig struct {
	BasePath   string
	DateFormat string
}

//工厂方法获取模块特定的logger
func GetLogger(name string) Logger {
	return &loggerImpl{name}
}

//配置logger工厂
func ConfigLogger(config LoggerConfig) {

}

type logRecord struct {
	format string
	objs   []interface{}
}

//工厂
type _LogFactory struct {
	loglevel int
	logfile  *RollingFile
	l        *log.Logger
	running  bool
	queue    *Queue
}

func (this *_LogFactory) Write(level int, prefix string, fmt string, obj ...interface{}) {
	if level >= this.loglevel { //判断loglevel是否大于指定的level，如果大于则输出，否则直接抛弃
		content := this.formatHeader(prefix, level) + fmt + "\n"
		this.queue.push(&logRecord{content, obj})

	}
}

func (this *_LogFactory) formatHeader(prefix string, level int) string {
	_, thefile, line, _ := runtime.Caller(3)
	short := thefile
	for i := len(thefile) - 1; i > 0; i-- {
		if thefile[i] == '/' {
			short = thefile[i+1:]
			break
		}
	}
	return fmt.Sprintf("[%s][%s][%s(%d)]- ", LEVEL_NAMES[level-1], prefix, short, line)
}

//the logger interface
type Logger interface {
	Info(format string, parameters ...interface{})

	Debug(format string, parameters ...interface{})

	Notice(format string, parameters ...interface{})

	Warning(format string, parameters ...interface{})

	Error(format string, parameters ...interface{})

	Critical(format string, parameters ...interface{})

	Alert(format string, parameters ...interface{})

	Emergency(format string, parameters ...interface{})
	//手工关闭logger输出流
	Close()
}

//具体实现，不对外暴露
type loggerImpl struct {
	name    string       //模块名称
	factory *_LogFactory //工厂
}

func (this *loggerImpl) log(level int, format string, parameters ...interface{}) {
	this.factory.Write(level, this.name, format, parameters...)
}

func (this *loggerImpl) Info(format string, parameters ...interface{}) {
	this.log(INFO, format, parameters...)
}

func (this *loggerImpl) Debug(format string, parameters ...interface{}) {
	this.log(DEBUG, format, parameters...)
}

func (this *loggerImpl) Notice(format string, parameters ...interface{}) {
	this.log(NOTICE, format, parameters...)
}

func (this *loggerImpl) Warning(format string, parameters ...interface{}) {
	this.log(WARNING, format, parameters...)
}

func (this *loggerImpl) Error(format string, parameters ...interface{}) {
	this.log(ERROR, format, parameters...)
}

func (this *loggerImpl) Critical(format string, parameters ...interface{}) {
	this.log(CRITICAL, format, parameters...)
}

func (this *loggerImpl) Alert(format string, parameters ...interface{}) {
	this.log(ALERT, format, parameters...)
}

func (this *loggerImpl) Emergency(format string, parameters ...interface{}) {
	this.log(EMERGENCY, format, parameters...)
}

//手工关闭logger输出流
func (this *loggerImpl) Close() {

}

/**
  队列实现
*/
type Queue struct {
	data *list.List
	sync.Mutex
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) push(v interface{}) {
	q.Lock()
	q.data.PushFront(v)
	q.Unlock()
}

func (q *Queue) pop() interface{} {
	q.Lock()
	iter := q.data.Back()
	if iter != nil {
		v := iter.Value
		q.data.Remove(iter)
		q.Unlock()
		return v
	}

	q.Unlock()
	return nil
}

func (q *Queue) dump() {
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}
