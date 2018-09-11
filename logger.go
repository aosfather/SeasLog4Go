package SeasLog4Go

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
	return nil
}

//配置logger工厂
func ConfigLogger(config LoggerConfig) {

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
	name string //模块名称
}

func (this *loggerImpl) log(level string, format string, parameters ...interface{}) {

}

func (this *loggerImpl) Info(format string, parameters ...interface{}) {
	this.log(SEASLOG_INFO, format, parameters...)
}

func (this *loggerImpl) Debug(format string, parameters ...interface{}) {
	this.log(SEASLOG_DEBUG, format, parameters...)
}

func (this *loggerImpl) Notice(format string, parameters ...interface{}) {
	this.log(SEASLOG_NOTICE, format, parameters...)
}

func (this *loggerImpl) Warning(format string, parameters ...interface{}) {
	this.log(SEASLOG_WARNING, format, parameters...)
}

func (this *loggerImpl) Error(format string, parameters ...interface{}) {
	this.log(SEASLOG_ERROR, format, parameters...)
}

func (this *loggerImpl) Critical(format string, parameters ...interface{}) {
	this.log(SEASLOG_CRITICAL, format, parameters...)
}

func (this *loggerImpl) Alert(format string, parameters ...interface{}) {
	this.log(SEASLOG_ALERT, format, parameters...)
}

func (this *loggerImpl) Emergency(format string, parameters ...interface{}) {
	this.log(SEASLOG_EMERGENCY, format, parameters...)
}

//手工关闭logger输出流
func (this *loggerImpl) Close() {

}
