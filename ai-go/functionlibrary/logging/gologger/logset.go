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

var (
	dailyRollingSet bool = true
	RollingFileSet  bool = true
	WriteFileSet    bool = false
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
	var grade Unit
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
					if gradevalue, ok := insize(gradestr); ok {
						isrolling = true
						grade = gradevalue
					}
				} else if level, ok := inlevel(arg); ok {
					SetLevel(level)
				}
			case int:
				maxNumber = int64(arg)
			case bool:
				SetConsole(arg)
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

func SetRollingFileConfig(fileDir, fileName string, maxNumber int64, maxSize int64, _unit Unit) *FileConfig {
	var logObjSet *FileConfig
	if WriteFileSet {
		maxFileCount = maxNumber
		maxFileSize = maxSize * int64(_unit)
		RollingFileSet = true
		dailyRollingSet = false
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
	WriteFileSet = isWrite
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
	if WriteFileSet {
		RollingFileSet = false
		dailyRollingSet = true
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
	if dailyRollingSet {
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

func (f *FileConfig) nextSuffix() int {
	return int(f.suffix%int(maxFileCount) + 1)
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
	if dailyRollingSet {
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

func (l *FileConfig) Debug(format string, v ...interface{}) {
	l.write(color_darkgreen, DEBUG, debug, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Info(format string, v ...interface{}) {
	l.write(color_white, INFO, info, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Warning(format string, v ...interface{}) {
	l.write(color_yellow, WARN, warning, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Error(format string, v ...interface{}) {
	l.write(color_red, ERROR, err, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Fatal(format string, v ...interface{}) {
	l.write(color_purple, FATAL, fatal, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Success(format string, v ...interface{}) {
	l.write(color_green, SUCCESS, success, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Trace(format string, v ...interface{}) {
	l.write(color_blue, TRACE, trace, fmt.Sprintf(format, v...))
}

func (l *FileConfig) write(color uint8, level Level, logType, data string) {
	if dailyRollingSet {
		l.fileCheck()
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
		if WriteFileSet {
			defer catchError()
			l.mu.RLock()
			defer l.mu.RUnlock()
			l.log.Output(3, data) // 3 Indicates the path of running files
			console(color, data)
		} else {
			console(color, data)
		}
	}
}
