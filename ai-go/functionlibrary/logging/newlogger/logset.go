package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

const (
	setRollingDaily setRollingType = iota // 按天滚动
	setRollingSize                        // 按大小滚动
)

// 选项模式配置
type setRollingType int
type SetOption func(*setConfig)

type FileConfig struct {
	dir           string
	filename      string
	suffix        int
	isCover       bool
	date          time.Time
	mu            *sync.RWMutex
	logFile       *os.File
	log           *log.Logger
	maxFileSizes  int64
	maxFileCounts int64
	dailyRollings bool
	RollingFiles  bool
	*setConfig
}

type setConfig struct {
	isSourcePath     bool
	rollingType      setRollingType
	maxFiles         int64
	maxFileSize      int64
	grade            Units
	level            Levels
	console          bool
	consoleAppenders bool
	WriteFiles       bool
}

func LoggerSet(pathfile string, args ...any) *FileConfig {
	var maxFiles int64 = 10
	var grade Units = Mb
	var maxFileSize int64 = 10
	var cfg = &setConfig{
		rollingType:      setRollingDaily,
		maxFiles:         maxFiles,
		maxFileSize:      maxFileSize,
		level:            SUCCESSED,
		grade:            grade,
		WriteFiles:       true,
		console:          true,
		consoleAppenders: true,
	}
	var options []SetOption
	var otheroptions []any
	dir, filename := filepath.Split(pathfile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "create log dir error: %v\n", err)
			return nil
		}
	}
	if len(args) > 0 {
		for _, arg := range args {
			switch opt := arg.(type) {
			case SetOption:
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
	cfg.IsWriteFileSet(true)
	if cfg.rollingType == setRollingDaily {
		return SetRolling(dir, filename, cfg)
	} else {
		return SetRollingFileConfig(dir, filename, cfg)
	}
}

func WithMaxFileSizeSet(sizeobj any) SetOption {
	var size int64 = 1
	var grade Units = Mb
	switch sizevalue := sizeobj.(type) {
	case string:
		if sizeVal, gradestr, ok := extract(sizevalue); ok {
			size = sizeVal
			if gradevalue, ok := insizealone(gradestr); ok {
				grade = gradevalue
			}
		} else if gradevalue, ok := insizealone(sizevalue); ok {
			grade = gradevalue
		}
	case int:
		size = int64(sizevalue)
	}
	return func(c *setConfig) {
		c.rollingType = setRollingSize
		c.maxFileSize = size
		c.grade = grade
	}
}

func WithMaxFilesSet(n int64) SetOption {
	return func(c *setConfig) {
		c.maxFiles = n
	}
}

func WithSourcePathSet(enable bool) SetOption {
	return func(c *setConfig) {
		c.isSourcePath = enable
	}
}

func WithLevelSet(levelobj any) SetOption {
	var level Levels = SUCCESSED
	if levelVal, ok := inlevelalone(levelobj); ok {
		level = levelVal
	}
	return func(c *setConfig) {
		c.SetLevel(level)
	}
}

func WithConsoleSet(enable bool) SetOption {
	return func(c *setConfig) {
		c.SetConsole(enable)
	}
}

func (cfg *setConfig) parseOptions(options ...SetOption) {
	for _, optfunc := range options {
		optfunc(cfg)
	}
}

func (cfg *setConfig) parseOtherOptions(options ...any) {
	var size int64 = 1
	var grade Units = Mb
	var maxfile int64 = 5
	var boolnum = 0
	for _, opt := range options {
		switch value := opt.(type) {
		case string:
			if maxsize, gradestr, ok := extract(value); ok {
				cfg.rollingType = setRollingSize
				size = maxsize
				if gradevalue, ok := insizealone(gradestr); ok {
					grade = gradevalue
				}
			} else if gradevalue, ok := insizealone(value); ok {
				cfg.rollingType = setRollingSize
				grade = gradevalue
			} else if level, ok := inlevelalone(value); ok {
				cfg.SetLevel(level)
			}
		case int:
			maxfile = int64(value)
		case bool:
			boolnum++
			if boolnum == 1 {
				cfg.SetConsole(value)
			} else if boolnum == 2 {
				cfg.SetSourcePath(value)
			}
		}
	}
	cfg.maxFileSize = size
	cfg.grade = grade
	cfg.maxFiles = maxfile
}

func SetRolling(fileDir, fileName string, cfg *setConfig) *FileConfig {
	var logObjSet *FileConfig
	if cfg.WriteFiles {
		now := time.Now()
		logObjSet = &FileConfig{
			dir:           fileDir,
			filename:      fileName,
			date:          now,
			mu:            new(sync.RWMutex),
			RollingFiles:  false,
			dailyRollings: true,
			setConfig:     cfg,
		}
		logObjSet.mu.Lock()
		if !logObjSet.isMustRename() {
			logObjSet.openLogFile()
		} else {
			logObjSet.rename()
		}
		logObjSet.mu.Unlock()
	}
	return logObjSet
}

func SetRollingFileConfig(fileDir, fileName string, cfg *setConfig) *FileConfig {
	var logObjSet *FileConfig
	if cfg.WriteFiles {
		now := time.Now()
		logObjSet = &FileConfig{
			dir:           fileDir,
			filename:      fileName,
			date:          now,
			mu:            new(sync.RWMutex),
			RollingFiles:  true,
			dailyRollings: false,
			maxFileCounts: cfg.maxFiles,
			maxFileSizes:  cfg.maxFileSize * int64(cfg.grade),
			setConfig:     cfg,
		}
		logObjSet.mu.Lock()
		for i := 1; i <= int(cfg.maxFiles); i++ {
			if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
				logObjSet.suffix = i
			} else {
				break
			}
		}
		if !logObjSet.isMustRename() {
			logObjSet.openLogFile()
		} else {
			logObjSet.rename()
		}
		logObjSet.mu.Unlock()
		go logObjSet.fileMonitorSet()
	}
	return logObjSet
}

func (f *FileConfig) openLogFile() {
	filePath := f.dir + "/" + f.filename
	var err error
	f.logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open log file error: %v\n", err)
		return
	}
	f.log = log.New(f.logFile, "", 0)
}

func (f *FileConfig) isMustRename() bool {
	if f.dailyRollings {
		now := time.Now()
		if now.Day() != f.date.Day() {
			return true
		}
	} else if f.RollingFiles {
		if f.maxFileCounts > 1 {
			if fileSize(f.dir+"/"+f.filename) >= f.maxFileSizes {
				return true
			}
		}
	}
	return false
}

func (f *FileConfig) nextSuffix() int {
	return int(f.suffix)%int(f.maxFileCounts) + 1
}

func (f *FileConfig) rename() {
	if f.dailyRollings {
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

func (cfg *setConfig) IsWriteFileSet(isWrite bool) {
	cfg.WriteFiles = isWrite
}

func (f *FileConfig) fileMonitorSet() {
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			f.fileCheck()
		}
	}
}

func (f *FileConfig) fileCheck() {
	if f != nil && f.isMustRename() {
		f.mu.Lock()
		f.rename()
		f.mu.Unlock()
	}
}

func (cfg *setConfig) SetLevel(level Levels) {
	cfg.level = level
}

func (cfg *setConfig) SetConsole(isConsole bool) {
	cfg.consoleAppenders = isConsole
}

func (cfg *setConfig) SetSourcePath(enable bool) {
	isSourcePath = enable
}

func insizealone(grade string) (Units, bool) {
	switch strings.ToUpper(grade) {
	case "KB":
		return Kb, true
	case "MB":
		return Mb, true
	case "GB":
		return Gb, true
	case "TB":
		return Tb, true
	default:
		return Mb, false
	}
}

func inlevelalone(levelobj any) (Levels, bool) {
	switch level := levelobj.(type) {
	case string:
		switch strings.ToUpper(level) {
		case "SUCCESS":
			return SUCCESSED, true
		case "TRACE":
			return TRACEED, true
		case "DEBUG":
			return DEBUGING, true
		case "INFO":
			return INFOING, true
		case "WARN":
			return WARNING, true
		case "ERROR":
			return ERRORING, true
		case "FATAL":
			return FATALED, true
		case "OFF":
			return OFFING, true
		}
	case int:
		switch level {
		case 1:
			return SUCCESSED, true
		case 2:
			return TRACEED, true
		case 3:
			return DEBUGING, true
		case 4:
			return INFOING, true
		case 5:
			return WARNING, true
		case 6:
			return ERRORING, true
		case 7:
			return FATALED, true
		case 8:
			return OFFING, true
		}
	}
	return SUCCESSED, false
}

func (l *FileConfig) Debug(v ...any) {
	l.write(color_darkgreen, DEBUGING, debug, fmt.Sprint(v...))
}

func (l *FileConfig) Debugf(format string, v ...any) {
	l.write(color_darkgreen, DEBUGING, debug, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Info(v ...any) {
	l.write(color_white, INFOING, info, fmt.Sprint(v...))
}

func (l *FileConfig) Infof(format string, v ...any) {
	l.write(color_white, INFOING, info, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Warning(v ...any) {
	l.write(color_yellow, WARNING, warning, fmt.Sprint(v...))
}

func (l *FileConfig) Warningf(format string, v ...any) {
	l.write(color_yellow, WARNING, warning, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Error(v ...any) {
	l.write(color_red, ERRORING, err, fmt.Sprint(v...))
}

func (l *FileConfig) Errorf(format string, v ...any) {
	l.write(color_red, ERRORING, err, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Fatal(v ...any) {
	l.write(color_purple, FATALED, fatal, fmt.Sprint(v...))
}

func (l *FileConfig) Fatalf(format string, v ...any) {
	l.write(color_purple, FATALED, fatal, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Success(v ...any) {
	l.write(color_green, SUCCESSED, success, fmt.Sprint(v...))
}

func (l *FileConfig) Successf(format string, v ...any) {
	l.write(color_green, SUCCESSED, success, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Trace(v ...any) {
	l.write(color_blue, TRACEED, trace, fmt.Sprint(v...))
}

func (l *FileConfig) Tracef(format string, v ...any) {
	l.write(color_blue, TRACEED, trace, fmt.Sprintf(format, v...))
}

func (l *FileConfig) Off(v ...any) {
	l.write(color_white, OFFING, off, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *FileConfig) Offf(format string, v ...any) {
	l.write(color_white, OFFING, off, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *FileConfig) Print(v ...any) {
	var data string
	if l.isSourcePath {
		_, file, line, _ := runtime.Caller(1)
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
			}
		}
		file = short
		data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), trace, file+":"+strconv.Itoa(line), fmt.Sprint(v...))
	} else {
		data = fmt.Sprintf("[%v] [%v] >>> %v", time.Now().Format(Timeformat), trace, fmt.Sprint(v...))
	}
	l.consolealone(color_white, data)
}

func (l *FileConfig) Printf(format string, v ...any) {
	var data string
	if l.isSourcePath {
		_, file, line, _ := runtime.Caller(1)
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
			}
		}
		file = short
		data = fmt.Sprintf("[%v] [%v] [%v] >>> %v", time.Now().Format(Timeformat), trace, file+":"+strconv.Itoa(line), fmt.Sprintf(format, v...))
	} else {
		data = fmt.Sprintf("[%v] [%v] >>> %v", time.Now().Format(Timeformat), trace, fmt.Sprintf(format, v...))
	}
	l.consolealone(color_white, data)
}

func (l *FileConfig) write(color uint8, level Levels, logType, data string) {
	if l.dailyRollings {
		l.fileCheck()
	}
	if l.level <= level {
		if l.isSourcePath {
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
		if l.WriteFiles {
			defer catchError()
			l.mu.RLock()
			if l.log != nil {
				l.log.Output(3, data)
			}
			l.mu.RUnlock()
		}
		l.consolealone(color, data)
	}
}

func (l *FileConfig) consolealone(color uint8, data string) {
	if l.consoleAppenders {
		fmt.Printf("\x1b[%dm%s\x1b[0m\n", color, data)
	}
}
