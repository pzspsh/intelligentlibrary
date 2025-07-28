package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 时间格式
const (
	Version    string = "1.0.0"
	Dateformat        = "2006-01-02"
	Timeformat        = "2006-01-02 15:04:05"
)

// 颜色常量
const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_purple
	color_darkgreen
	color_white
)

// 日志类型
const (
	err     = "ERROR"
	success = "SUCCESS"
	warning = "WARNING"
	trace   = "TRACE"
	fatal   = "FATAL"
	debug   = "DEBUG"
	info    = "INFO"
)

// 日志级别
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

// 单位常量
const (
	_       = iota
	KB Unit = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	// ZB // 注释掉 ZB 和 YB，避免 int64 溢出
	// YB
)

type (
	Level int32
	Unit  int64
)

var (
	logLevel        Level = SUCCESS
	maxFileSize     int64
	maxFileCount    int64
	isSourcePath    bool = false
	dailyRolling    bool = true
	consoleAppender bool = true
	RollingFile     bool = true
	WriteFile       bool = false
	logObj          *File
)

const (
	rollingDaily rollingType = iota // 按天滚动
	rollingSize                     // 按大小滚动
)

// 选项模式配置
type rollingType int
type Option func(*config)

// File 结构体
type File struct {
	dir      string
	filename string
	suffix   int
	isCover  bool
	date     time.Time
	mu       *sync.RWMutex
	logFile  *os.File
	log      *log.Logger
}

// 全局配置
type config struct {
	rollingType rollingType
	maxFiles    int64
	maxFileSize int64
	grade       Unit
	level       Level
	console     bool
}

// 初始化日志系统
func Logger(logPath string, args ...any) {
	var maxFiles int64 = 10
	var grade = MB
	var maxFileSize int64 = 10 // 默认 10MB
	var cfg = &config{
		rollingType: rollingDaily,
		maxFiles:    maxFiles,
		maxFileSize: maxFileSize,
		level:       SUCCESS,
		grade:       grade,
		console:     true,
	}
	var options []Option
	var otheroptions []any
	dir, filename := filepath.Split(logPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "create log dir error: %v\n", err)
			return
		}
	}
	if len(args) > 0 {
		for _, arg := range args {
			switch opt := arg.(type) {
			case Option:
				options = append(options, opt)
			default:
				otheroptions = append(otheroptions, opt)
			}
		}
	}
	if len(otheroptions) > 0 {
		cfg.parseOtherOptions(otheroptions...)
	}
	if len(options) > 0 {
		cfg.parseOptions(options...)
	}
	SetWriteFile(true)
	if cfg.rollingType == rollingDaily {
		SetRollingDaily(dir, filename)
	} else {
		SetRollingFile(dir, filename, cfg.maxFiles, cfg.maxFileSize, cfg.grade)
	}
}

func WithMaxFileSize(sizeobj any) Option {
	var size int64 = 1
	var grade Unit = MB
	switch sizevalue := sizeobj.(type) {
	case string:
		if sizeVal, gradestr, ok := extract(sizevalue); ok {
			size = sizeVal
			if gradevalue, ok := insize(gradestr); ok {
				grade = gradevalue
			}
		} else if gradevalue, ok := insize(sizevalue); ok {
			grade = gradevalue
		}
	case int:
		size = int64(sizevalue)
	}
	return func(c *config) {
		c.rollingType = rollingSize
		c.maxFileSize = size
		c.grade = grade
	}
}

func WithMaxFiles(n int64) Option {
	return func(c *config) {
		c.maxFiles = n
	}
}

func WithSourcePath(enable bool) Option {
	return func(c *config) {
		SetSourcePath(enable)
	}
}

func WithLevel(levelobj any) Option {
	var level Level = SUCCESS
	if levelVal, ok := inlevel(levelobj); ok {
		level = levelVal
	}
	return func(c *config) {
		SetLevel(level)
	}
}

func WithConsole(enable bool) Option {
	return func(c *config) {
		SetConsole(enable)
	}
}

func (cfg *config) parseOptions(options ...Option) {
	for _, optfunc := range options {
		optfunc(cfg)
	}
}

func (cfg *config) parseOtherOptions(options ...any) {
	var size int64 = 1
	var grade Unit = MB
	var maxfile int64 = 5
	var boolnum = 0
	for _, opt := range options {
		switch value := opt.(type) {
		case string:
			if maxsize, gradestr, ok := extract(value); ok {
				cfg.rollingType = rollingSize
				size = maxsize
				if gradevalue, ok := insize(gradestr); ok {
					grade = gradevalue
				}
			} else if gradevalue, ok := insize(value); ok {
				cfg.rollingType = rollingSize
				grade = gradevalue
			} else if level, ok := inlevel(value); ok {
				SetLevel(level)
			}
		case int:
			maxfile = int64(value)
		case bool:
			boolnum++
			if boolnum == 1 {
				SetConsole(value)
			} else if boolnum == 2 {
				SetSourcePath(value)
			}
		}
	}
	cfg.maxFileSize = size
	cfg.grade = grade
	cfg.maxFiles = maxfile
}

func SetRollingDaily(fileDir, fileName string) {
	if WriteFile {
		RollingFile = false
		dailyRolling = true
		now := time.Now()
		logObj = &File{
			dir:      fileDir,
			filename: fileName,
			date:     now,
			mu:       new(sync.RWMutex),
		}
		logObj.mu.Lock()
		// defer logObj.mu.Unlock()
		if !logObj.isMustRename() {
			logObj.openLogFile()
		} else {
			logObj.rename()
		}
		logObj.mu.Unlock()
	}
}

func SetRollingFile(fileDir, fileName string, maxNumber int64, maxSize int64, _unit Unit) {
	if WriteFile {
		RollingFile = true
		dailyRolling = false
		maxFileCount = maxNumber
		maxFileSize = maxSize * int64(_unit)
		now := time.Now()
		logObj = &File{
			dir:      fileDir,
			filename: fileName,
			date:     now,
			mu:       new(sync.RWMutex),
		}
		logObj.mu.Lock()
		// defer logObj.mu.Unlock()
		for i := 1; i <= int(maxNumber); i++ {
			if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
				logObj.suffix = i
			} else {
				break
			}
		}
		if !logObj.isMustRename() {
			logObj.openLogFile()
		} else {
			logObj.rename()
		}
		logObj.mu.Unlock()
		go fileMonitor()
	}
}

func (f *File) openLogFile() {
	filePath := f.dir + "/" + f.filename
	var err error
	f.logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open log file error: %v\n", err)
		return
	}
	f.log = log.New(f.logFile, "", 0)
}

func (f *File) isMustRename() bool {
	if dailyRolling {
		now := time.Now()
		if now.Day() != f.date.Day() {
			return true
		}
	} else if RollingFile {
		if maxFileCount > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
				return true
			}
		}
	}
	return false
}

func (f *File) nextSuffix() int {
	return int(f.suffix)%int(maxFileCount) + 1
}

func (f *File) rename() {
	if dailyRolling {
		dateStr := f.date.Format(Dateformat)
		newName := f.dir + "/" + f.filename + "." + dateStr
		if !isExist(newName) && f.isMustRename() {
			if f.logFile != nil {
				f.logFile.Close()
			}
			os.Rename(f.dir+"/"+f.filename, newName)
			f.openLogFile()
			f.date = time.Now()
		}
	} else {
		newSuffix := f.nextSuffix()
		if f.logFile != nil {
			f.logFile.Close()
		}
		oldName := f.dir + "/" + f.filename
		newName := oldName + "." + strconv.Itoa(newSuffix)
		if isExist(newName) {
			os.Remove(newName)
		}
		os.Rename(oldName, newName)
		f.suffix = newSuffix
		f.openLogFile()
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
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fileCheck()
		}
	}
}

func fileCheck() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()
	if logObj != nil && logObj.isMustRename() {
		logObj.mu.Lock()
		// defer logObj.mu.Unlock()
		logObj.rename()
		logObj.mu.Unlock()
	}
}

func SetConsole(enable bool) {
	consoleAppender = enable
}

func SetLevel(level Level) {
	logLevel = level
}

func SetSourcePath(enable bool) {
	isSourcePath = enable
}

func SetWriteFile(enable bool) {
	WriteFile = enable
}

func insize(grade string) (Unit, bool) {
	switch strings.ToUpper(grade) {
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	case "TB":
		return TB, true
	default:
		return MB, false
	}
}

func inlevel(levelobj any) (Level, bool) {
	switch level := levelobj.(type) {
	case string:
		switch strings.ToUpper(level) {
		case "SUCCESS":
			return SUCCESS, true
		case "TRACE":
			return TRACE, true
		case "DEBUG":
			return DEBUG, true
		case "INFO":
			return INFO, true
		case "WARN":
			return WARN, true
		case "ERROR":
			return ERROR, true
		case "FATAL":
			return FATAL, true
		case "OFF":
			return OFF, true
		}
	case int:
		switch level {
		case 1:
			return SUCCESS, true
		case 2:
			return TRACE, true
		case 3:
			return DEBUG, true
		case 4:
			return INFO, true
		case 5:
			return WARN, true
		case 6:
			return ERROR, true
		case 7:
			return FATAL, true
		case 8:
			return OFF, true
		}
	}
	return SUCCESS, false
}

func extract(data string) (int64, string, bool) {
	re := regexp.MustCompile(`(\d+)([a-zA-Z]+)`)
	match := re.FindStringSubmatch(data)
	if len(match) == 3 {
		if size, err := strconv.ParseInt(match[1], 10, 64); err == nil {
			return size, match[2], true
		}
	}
	return 0, "", false
}

func Debug(format string, v ...any) {
	write(color_darkgreen, DEBUG, debug, fmt.Sprintf(format, v...))
}

func Info(format string, v ...any) {
	write(color_white, INFO, info, fmt.Sprintf(format, v...))
}

func Warning(format string, v ...any) {
	write(color_yellow, WARN, warning, fmt.Sprintf(format, v...))
}

func Error(format string, v ...any) {
	write(color_red, ERROR, err, fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...any) {
	write(color_purple, FATAL, fatal, fmt.Sprintf(format, v...))
}

func Success(format string, v ...any) {
	write(color_green, SUCCESS, success, fmt.Sprintf(format, v...))
}

func Trace(format string, v ...any) {
	write(color_blue, TRACE, trace, fmt.Sprintf(format, v...))
}

func Off(format string, v ...any) {
	write(color_white, OFF, err, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Print(format string, v ...any) {
	var data string
	if isSourcePath {
		_, file, line, _ := runtime.Caller(1)
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
			}
		}
		file = short
		data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), info, file+":"+strconv.Itoa(line), fmt.Sprintf(format, v...))
	} else {
		data = fmt.Sprintf("[%v] [%v] >>> %v", time.Now().Format(Timeformat), info, fmt.Sprintf(format, v...))
	}
	console(color_white, data)
}

func write(color uint8, level Level, logType, data string) {
	if dailyRolling {
		fileCheck()
	}
	if logLevel <= level {
		if isSourcePath {
			_, file, line, _ := runtime.Caller(2)
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
				}
			}
			file = short
			data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), logType, file+":"+strconv.Itoa(line), data)
		} else {
			data = fmt.Sprintf("[%v] [%v] >>> %v", time.Now().Format(Timeformat), logType, data)
		}
		if WriteFile {
			// defer catchError()
			logObj.mu.RLock()
			if logObj.log != nil {
				logObj.log.Output(3, data)
			}
			logObj.mu.RUnlock()
		}
		console(color, data)
	}
}

func console(color uint8, data string) {
	if consoleAppender {
		fmt.Printf("\x1b[%dm%s\x1b[0m\n", color, data)
	}
}
