package main

import (
	"log"
	"os"
)

//logger files break into three sections - prefix, automatic generated
//information, and log message
//prefix is controlled with second log.New arguement
//put in a whitespace ffs
//here we set the date and file and line info
//Ldate - date. Ltime - timestamp Lmicrosends - precisie time
//LstdFlags turns on data and time
//Llongfile - file path and line no
//Lsthortfile - filename and line number

func main() {
	//creates a log file and closes it in the end
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	//creates a logger
	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	//logs to file. last one will never get called as there is a fatal.
	logger.Println("This is a regualar message.")
	logger.Fatalln("this is a fatal error")
	logger.Println("this tis the end of the function.")
}
