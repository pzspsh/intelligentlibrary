/*
@File   : main.go
@Author : pan
@Time   : 2024-01-10 10:23:44
*/
package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// 启动文件监听
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go watchForChanges(watcher)

	// 主循环
	for {
		select {
		case event := <-watcher.Events:
			// 处理文件变化事件
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File modified:", event.Name)
				restartApp()
			}
		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}

func watchForChanges(watcher *fsnotify.Watcher) {
	// 监听当前目录下的所有Go文件
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func restartApp() {
	log.Println("Restarting application...")
	cmd := exec.Command("go", "run", ".") // 这里使用了简单的重新运行方式，实际生产中可以替换为编译命令并重启
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error restarting application:", err)
	}
	// 在实际生产中，你可能需要等待一段时间，确保新的进程完全启动
	time.Sleep(2 * time.Second)
	os.Exit(0) // 退出当前进程，让新的进程接管
}
