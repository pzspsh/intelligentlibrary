package main

import "fmt"

/*
# 模版方法模式
	模版方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。使得实现符合开闭原则。
	如实例代码中通用步骤在父类中实现（`准备`、`下载`、`保存`、`收尾`）下载和保存的具体实现留到子类中，并且提供 `保存`方法的默认实现。
	因为Golang不提供继承机制，需要使用匿名组合模拟实现继承。
	此处需要注意：因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类需要持有子类的引用。

解释：
	模板方法模式定义了一个算法的步骤，并允许子类别为一个或多个步骤提供其实践方式。让子类别在不改变算法架构的情况下，重新
定义算法中的某些步骤。在软件工程中，它是一种软件设计模式，和C++模板没有关连。



简介：
	模板方法模式定义了一个算法的步骤，并允许子类别为一个或多个步骤提供其实践方式。让子类别在不改变算法架构的情况下，重新
定义算法中的某些步骤。在软件工程中，它是一种软件设计模式，和C++模板没有关连。

概念
	定义一个操作中的算法的框架，而将一些步骤延迟到子类中。使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

使用场景
	多个子类有公有的方法，并且逻辑基本相同时。
	重要、复杂的算法，可以把核心算法设计为模板方法，周边的相关细节功能则由各个子类实现。
	重构时，模板方法模式是一个经常使用的模式，把相同的代码抽取到父类中，然后通过钩子函数（见“模板方法模式的扩展”）约束其行为。

结构
	抽象模板：AbstractClass为抽象模板，它的方法分为两类：
	1、基本方法：也叫做基本操作，是由子类实现的方法，并且在模板方法被调用。
	2、模板方法：可以有一个或几个，一般是一个具体方法，也就是一个框架，实现对基本方法的调度，完成固定的逻辑。
	注意： 为了防止恶意的操作，一般模板方法都加上final关键字，不允许被覆写。
	具体模板：实现父类所定义的一个或多个抽象方法，也就是父类定义的基本方法在子类中得以实现。
*/

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

func main() {
	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

	var downloader1 Downloader = NewFTPDownloader()
	downloader1.Download("ftp://example.com/abc.zip")
}
