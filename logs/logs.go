package logs

import "log"

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started...")
}

func RunTrialLogs() {
	log.Println("main started...")
	log.Fatalln("fatal message")
	// log.Panicln("panic message")

}
