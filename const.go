package SeasLog4Go

//常量定义
//日志级别
const (
	SEASLOG_ALL       = "ALL"
	SEASLOG_DEBUG     = "DEBUG"
	SEASLOG_INFO      = "INFO"
	SEASLOG_NOTICE    = "NOTICE"
	SEASLOG_WARNING   = "WARNING"
	SEASLOG_ERROR     = "ERROR"
	SEASLOG_CRITICAL  = "CRITICAL"
	SEASLOG_ALERT     = "ALERT"
	SEASLOG_EMERGENCY = "EMERGENCY"
)

//排序方式
const (
	SEASLOG_DETAIL_ORDER_ASC  = 1
	SEASLOG_DETAIL_ORDER_DESC = 2
)

const (
	SEASLOG_CLOSE_LOGGER_STREAM_MOD_ALL    = 1
	SEASLOG_CLOSE_LOGGER_STREAM_MOD_ASSIGN = 2
)
