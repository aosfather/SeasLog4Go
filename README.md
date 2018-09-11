# SeasLog4Go
SeasLog for Golang

Examples:
'''Go
   import （seaslog 'github.com/SeasX/SeasLog4Go'）

   //1、config the logger
   config:=LoggerConfig{"/var/log/test",""}
   seaslog.ConfigLogger(config)
   logger:=seaslog.GetLogger("firstModule")
   logger.Info("the first line for %s,%s","test","done")
'''