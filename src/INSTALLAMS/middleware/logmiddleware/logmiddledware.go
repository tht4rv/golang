package logmiddleware

import(
	"net/http"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

func LogMiddleware( handler http.Handler) http.Handler{
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request){
		var logfile = `.\log\operation\logfile.log`
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		Formatter := new(log.TextFormatter)
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FullTimestamp = true
		log.SetFormatter(Formatter)
		log.SetReportCaller(true)
		if err != nil {
			fmt.Println(err)
		} else {
			log.SetOutput(f)
		}
		log.Info(request)
		handler.ServeHTTP(response,request)
	})
}