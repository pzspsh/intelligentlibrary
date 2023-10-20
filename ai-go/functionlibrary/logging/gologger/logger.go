package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	Version    string = "1.0.0"
	Dateformat string = "2006-01-02"
	Timeformat string = "2006-01-02 15:04:05"
)

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_purple
	color_darkgreen
	color_white
	err     = "ERROR"
	success = "SUCCESS"
	warning = "WARNING"
	trace   = "TRACE"
	fatal   = "FATAL"
	debug   = "DEBUG"
	info    = "INFO"
)

const (
	ALL Level = iota
	SUCCESS
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

const (
	_       = iota
	KB Unit = 1 << (iota * 10)
	MB
	GB
	TB
)

type (
	Unit  int64
	Level int32
)

var (
	logLevel        Level = 1
	maxFileSize     int64
	maxFileCount    int64
	dailyRolling    bool = true
	consoleAppender bool = true
	RollingFile     bool = true
	WriteFile       bool = true
	logObj          *File
)

type File struct {
	dir      string
	filename string
	suffix   int
	isCover  bool
	date     *time.Time
	mu       *sync.RWMutex
	logFile  *os.File
	log      *log.Logger
}

func Logger(pathfile string, args ...interface{}) {
	var maxNumber int64
	var maxSize int64
	var isrolling bool
	var grade Unit
	logpath, logfile := filepath.Split(pathfile)
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		err = os.MkdirAll(logpath, os.ModePerm)
		if err != nil {
			fmt.Println("create log file error:", err)
			return
		}
	}
	if len(args) > 0 {
		for _, arg := range args {
			switch arg := arg.(type) {
			case string:
				if maxsize, gradestr, ok := extract(arg); ok {
					maxSize = maxsize
					if gradevalue, ok := insize(gradestr); ok {
						isrolling = true
						grade = gradevalue
					}
				} else if level, ok := inlevel(arg); ok {
					SetLevel(level)
				}
			case int:
				maxNumber = int64(arg)
			default:
				SetRollingDaily(logpath, logfile)
			}
		}
		if isrolling {
			SetRollingFile(logpath, logfile, maxNumber, maxSize, grade)
		} else {
			SetRollingDaily(logpath, logfile)
		}
	} else {
		SetRollingDaily(logpath, logfile)
	}
}

func SetConsole(isConsole bool) {
	consoleAppender = isConsole
}

func SetLevel(level Level) {
	logLevel = level
}

func SetWriteFile(isWrite bool) {
	WriteFile = isWrite
}

func SetRollingFile(fileDir, fileName string, maxNumber int64, maxSize int64, _unit Unit) {
	if WriteFile {
		maxFileCount = maxNumber
		maxFileSize = maxSize * int64(_unit)
		RollingFile = true
		dailyRolling = false
		logObj = &File{dir: fileDir, filename: fileName, isCover: false, mu: new(sync.RWMutex)}
		logObj.mu.Lock()
		defer logObj.mu.Unlock()
		for i := 1; i <= int(maxNumber); i++ {
			if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
				logObj.suffix = i
			} else {
				break
			}
		}
		if !logObj.isMustRename() {
			logObj.logFile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
			logObj.log = log.New(logObj.logFile, "", 0)
		} else {
			logObj.rename()
		}
		go fileMonitor()
	}
}

func SetRollingDaily(fileDir, fileName string) {
	if WriteFile {
		RollingFile = false
		dailyRolling = true
		t, _ := time.Parse(Dateformat, time.Now().Format(Dateformat))
		logObj = &File{dir: fileDir, filename: fileName, date: &t, isCover: false, mu: new(sync.RWMutex)}
		logObj.mu.Lock()
		defer logObj.mu.Unlock()
		if !logObj.isMustRename() {
			logObj.logFile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
			logObj.log = log.New(logObj.logFile, "", 0)
			// logObj.log = log.New(logObj.logFile, "",log.Ldate|log.Ltime|log.Llongfile)
		} else {
			logObj.rename()
		}
	}
}

func (f *File) isMustRename() bool {
	if dailyRolling {
		t, _ := time.Parse(Dateformat, time.Now().Format(Dateformat))
		if t.After(*f.date) {
			return true
		}
	} else {
		if maxFileCount > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
				return true
			}
		}
	}
	return false
}

func (f *File) nextSuffix() int {
	return int(f.suffix%int(maxFileCount) + 1)
}

func (f *File) coverNextOne() {
	f.suffix = f.nextSuffix()
	if f.logFile != nil {
		f.logFile.Close()
	}
	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(f.suffix)) {
		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(f.suffix))
	}
	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(f.suffix))
	f.logFile, _ = os.Create(f.dir + "/" + f.filename)
	f.log = log.New(logObj.logFile, "", 0)
}

func (f *File) rename() {
	if dailyRolling {
		fn := f.dir + "/" + f.filename + "." + f.date.Format(Dateformat)
		if !isExist(fn) && f.isMustRename() {
			if f.logFile != nil {
				f.logFile.Close()
			}
			err := os.Rename(f.dir+"/"+f.filename, fn)
			if err != nil {
				f.log.Println("rename err", err.Error())
			}
			t, _ := time.Parse(Dateformat, time.Now().Format(Dateformat))
			f.date = &t
			f.logFile, _ = os.Create(f.dir + "/" + f.filename)
			f.log = log.New(logObj.logFile, "\n", log.Ldate|log.Ltime|log.Llongfile)
		}
	} else {
		f.coverNextOne()
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func fileSize(file string) int64 {
	f, err := os.Stat(file)
	if err != nil {
		return 0
	}
	return f.Size()
}

func catchError() {
	if err := recover(); err != nil {
		log.Println("err", err)
	}
}

func fileMonitor() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fileCheck()
		default:
			continue
		}
	}
}

func fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if logObj != nil && logObj.isMustRename() {
		logObj.mu.Lock()
		defer logObj.mu.Unlock()
		logObj.rename()
	}
}

func Write(color uint8, level Level, logType, data string) {
	if dailyRolling {
		fileCheck()
	}
	if logLevel <= level {
		_, file, line, _ := runtime.Caller(2)
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
			}
		}
		file = short
		data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), logType, file+":"+strconv.Itoa(line), data)
		if WriteFile {
			defer catchError()
			logObj.mu.RLock()
			defer logObj.mu.RUnlock()
			logObj.log.Output(3, data) // 3 Indicates the path of running files
			console(color, data)
		} else {
			console(color, data)
		}
	}
}

func console(color uint8, data string) {
	if consoleAppender {
		data = setcolor(color, data)
		fmt.Println(data)
	}
}

func Debug(format string, v ...interface{}) {
	Write(color_darkgreen, DEBUG, debug, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	Write(color_white, INFO, info, fmt.Sprintf(format, v...))
}

func Warning(format string, v ...interface{}) {
	Write(color_yellow, WARN, warning, fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	Write(color_red, ERROR, err, fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	Write(color_purple, FATAL, fatal, fmt.Sprintf(format, v...))
}

func Success(format string, v ...interface{}) {
	Write(color_green, SUCCESS, success, fmt.Sprintf(format, v...))
}

func Trace(format string, v ...interface{}) {
	Write(color_blue, TRACE, trace, fmt.Sprintf(format, v...))
}

func setcolor(color uint8, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
}

func insize(grade string) (Unit, bool) {
	switch grade {
	case "KB", "kb":
		return KB, true
	case "MB", "mb":
		return MB, true
	case "GB", "gb":
		return GB, true
	case "TB", "tb":
		return TB, true
	default:
		return 0, false
	}
}

func inlevel(level string) (Level, bool) {
	switch level {
	case "SUCCESS", "success":
		return SUCCESS, true
	case "TRACE", "trace":
		return TRACE, true
	case "DEBUG", "debug":
		return DEBUG, true
	case "INFO", "info":
		return INFO, true
	case "WARN", "warn":
		return WARN, true
	case "ERROR", "error":
		return ERROR, true
	case "FATAL", "fatal":
		return FATAL, true
	case "OFF", "off":
		return OFF, true
	default:
		return 0, false
	}
}

func extract(data string) (int64, string, bool) {
	re := regexp.MustCompile("([0-9]+)([a-zA-Z]+)")
	values := re.FindStringSubmatch(data)
	if len(values) == 3 {
		if s, err := strconv.Atoi(values[1]); err == nil {
			return int64(s), values[2], true
		}
	}
	return 0, "", false
}
