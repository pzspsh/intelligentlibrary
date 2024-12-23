/*
@File   : logset.go
@Author : pan
@Time   : 2023-10-20 17:03:04
*/
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	_        = iota
	Kb Units = 1 << (iota * 10)
	Mb
	Gb
	Tb
)

const (
	All Levels = iota
	SUCCESSED
	TRACEED
	DEBUGING
	INFOING
	WARNING
	ERRORING
	FATALED
	OFFING
)

type (
	Units  int64
	Levels int32
)

var (
	logLevels        Levels = 1
	maxFileSizes     int64
	maxFileCounts    int64
	consoleAppenders bool = true
	dailyRollings    bool = true
	RollingFiles     bool = true
	WriteFiles       bool = false
)

type FileConfig struct {
	dir      string
	filename string
	suffix   int
	isCover  bool
	date     *time.Time
	mu       *sync.RWMutex
	logFile  *os.File
	log      *log.Logger
}

func LoggerSet(pathfile string, args ...interface{}) *FileConfig {
	var maxNumber int64
	var maxSize int64
	var isrolling bool
	var grade Units
	logpath, logfile := filepath.Split(pathfile)
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		err = os.MkdirAll(logpath, os.ModePerm)
		if err != nil {
			fmt.Println("create log file error:", err)
			return nil
		}
	}
	IsWriteFileSet(true)
	if len(args) > 0 {
		for _, arg := range args {
			switch arg := arg.(type) {
			case string:
				if maxsize, gradestr, ok := extract(arg); ok {
					maxSize = maxsize
					if gradevalue, ok := insizealone(gradestr); ok {
						isrolling = true
						grade = gradevalue
					}
				} else if level, ok := inlevelalone(arg); ok {
					SetLevelSet(level)
				}
			case int:
				maxNumber = int64(arg)
			case bool:
				SetConsoleAlone(arg)
			default:
				return SetRolling(logpath, logfile)
			}
		}
		if isrolling {
			return SetRollingFileConfig(logpath, logfile, maxNumber, maxSize, grade)
		} else {
			return SetRolling(logpath, logfile)
		}
	} else {
		return SetRolling(logpath, logfile)
	}
}

func SetRollingFileConfig(fileDir, fileName string, maxNumber int64, maxSize int64, _unit Units) *FileConfig {
	var logObjSet *FileConfig
	if WriteFiles {
		maxFileCounts = maxNumber
		maxFileSizes = maxSize * int64(_unit)
		RollingFiles = true
		dailyRollings = false
		logObjSet = &FileConfig{dir: fileDir, filename: fileName, isCover: false, mu: new(sync.RWMutex)}
		logObjSet.mu.Lock()
		defer logObjSet.mu.Unlock()
		for i := 1; i <= int(maxNumber); i++ {
			if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
				logObjSet.suffix = i
			} else {
				break
			}
		}
		if !logObjSet.isMustRename() {
			logObjSet.logFile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
			logObjSet.log = log.New(logObjSet.logFile, "", 0)
		} else {
			logObjSet.rename()
		}
		go logObjSet.fileMonitorSet()
	}
	return logObjSet
}

func IsWriteFileSet(isWrite bool) {
	WriteFiles = isWrite
}

func (f *FileConfig) fileMonitorSet() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			f.fileCheck()
		default:
			continue
		}
	}
}

func SetRolling(fileDir, fileName string) *FileConfig {
	var logObjSet *FileConfig
	if WriteFiles {
		RollingFiles = false
		dailyRollings = true
		t, _ := time.Parse(Dateformat, time.Now().Format(Dateformat))
		logObjSet = &FileConfig{dir: fileDir, filename: fileName, date: &t, isCover: false, mu: new(sync.RWMutex)}
		logObjSet.mu.Lock()
		defer logObjSet.mu.Unlock()
		if !logObjSet.isMustRename() {
			logObjSet.logFile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
			logObjSet.log = log.New(logObjSet.logFile, "", 0)
		} else {
			logObjSet.rename()
		}
	}
	return logObjSet
}

func (f *FileConfig) isMustRename() bool {
	if dailyRollings {
		t, _ := time.Parse(Dateformat, time.Now().Format(Dateformat))
		if t.After(*f.date) {
			return true
		}
	} else {
		if maxFileCounts > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSizes {
				return true
			}
		}
	}
	return false
}

func SetLevelSet(level Levels) {
	logLevels = level
}

func SetConsoleAlone(isConsole bool) {
	consoleAppenders = isConsole
}

func (f *FileConfig) nextSuffix() int {
	return int(f.suffix%int(maxFileCounts) + 1)
}

func (f *FileConfig) coverNextOne() {
	f.suffix = f.nextSuffix()
	if f.logFile != nil {
		f.logFile.Close()
	}
	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(f.suffix)) {
		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(f.suffix))
	}
	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(f.suffix))
	f.logFile, _ = os.Create(f.dir + "/" + f.filename)
	f.log = log.New(f.logFile, "", 0)
}

func (f *FileConfig) rename() {
	if dailyRollings {
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
			f.log = log.New(f.logFile, "\n", log.Ldate|log.Ltime|log.Llongfile)
		}
	} else {
		f.coverNextOne()
	}
}

func (f *FileConfig) fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if f != nil && f.isMustRename() {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.rename()
	}
}

func insizealone(grade string) (Units, bool) {
	switch grade {
	case "KB", "kb":
		return Kb, true
	case "MB", "mb":
		return Mb, true
	case "GB", "gb":
		return Gb, true
	case "TB", "tb":
		return Tb, true
	default:
		return 0, false
	}
}

func inlevelalone(level string) (Levels, bool) {
	switch level {
	case "SUCCESS", "success":
		return SUCCESSED, true
	case "TRACE", "trace":
		return TRACEED, true
	case "DEBUG", "debug":
		return DEBUGING, true
	case "INFO", "info":
		return INFOING, true
	case "WARN", "warn":
		return WARNING, true
	case "ERROR", "error":
		return ERRORING, true
	case "FATAL", "fatal":
		return FATALED, true
	case "OFF", "off":
		return OFFING, true
	default:
		return 0, false
	}
}

func (l *FileConfig) Debug(format string, v ...interface{}) {
	l.write(color_darkgreen, DEBUGING, debug, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Info(format string, v ...interface{}) {
	l.write(color_white, INFOING, info, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Warning(format string, v ...interface{}) {
	l.write(color_yellow, WARNING, warning, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Error(format string, v ...interface{}) {
	l.write(color_red, ERRORING, err, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Fatal(format string, v ...interface{}) {
	l.write(color_purple, FATALED, fatal, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Success(format string, v ...interface{}) {
	l.write(color_green, SUCCESSED, success, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Trace(format string, v ...interface{}) {
	l.write(color_blue, TRACEED, trace, fmt.Sprintf(format, v...))
}

func (l *FileConfig) write(color uint8, level Levels, logType, data string) {
	if dailyRollings {
		l.fileCheck()
	}
	if logLevels <= level {
		_, file, line, _ := runtime.Caller(2)
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
			}
		}
		file = short
		data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), logType, file+":"+strconv.Itoa(line), data)
		if WriteFiles {
			defer catchError()
			l.mu.RLock()
			defer l.mu.RUnlock()
			l.log.Output(3, data) // 3 Indicates the path of running files
			consolealone(color, data)
		} else {
			consolealone(color, data)
		}
	}
}

func consolealone(color uint8, data string) {
	if consoleAppenders {
		data = setcolor(color, data)
		fmt.Println(data)
	}
}
