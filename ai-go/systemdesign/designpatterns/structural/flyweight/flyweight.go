package main

import "fmt"

/*
# 享元模式
	享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元，使多个对象共享，从而节省内存以及减少对象数量。

解释：
	享元模式（英语：Flyweight Pattern）是一种软件设计模式。它使用共享物件，用来尽可能减少内存使用量以及分享资讯给尽可能
多的相似物件；它适合用于只是因重复而导致使用无法令人接受的大量内存的大量物件。通常物件中的部分状态是可以分享。常见做法是
把它们放在外部数据结构，当需要使用时再将它们传递给享元。

概念
	使用共享对象可有效地支持大量的细粒度的对象。
	对象的信息分为两个部分：内部状态（intrinsic）与外部状态（extrinsic）。
	内部状态：内部状态是对象可共享出来的信息，存储在享元对象内部并且不会随环境改变而改变。
	外部状态：外部状态是对象得以依赖的一个标记，是随环境改变而改变的、不可以共享的状态。

定义：
	享元模式（FlyWeight），运用共享技术有效的支持大量细粒度的对象。
	典型的享元模式的例子为文书处理器中以图形结构来表示字符。一个做法是，每个字形有其字型外观, 字模 metrics，和其它格式资
讯，但这会使每个字符就耗用上千字节。取而代之的是，每个字符参照到一个共享字形物件，此物件会被其它有共同特质的字符所分享；
只有每个字符（文件中或页面中）的位置才需要另外储存。

使用场景
	系统中存在大量的相似对象。
	细粒度的对象都具备较接近的外部状态，而且内部状态与环境无关，也就是说对象没有特定身份。
	需要缓冲池的场景。

结构
	Flyweight——抽象享元角色：它简单地说就是一个产品的抽象类，同时定义出对象的外部状态和内部状态的接口或实现。
	ConcreteFlyweight——具体享元角色：具体的一个产品类，实现抽象角色定义的业务。该角色中需要注意的是内部状态处理应该与
环境无关，不应该出现一个操作改变了内部状态，同时修改了外部状态，这是绝对不允许的。
	unsharedConcreteFlyweight——不可共享的享元角色：不存在外部状态或者安全要求（如线程安全）不能够使用共享技术的对象，
该对象一般不会出现在享元工厂中。
	FlyweightFactory——享元工厂：职责非常简单，就是构造一个池容器，同时提供从池中获得对象的方法。

结构：
	内蕴状态存储在享元内部，不会随环境的改变而有所不同，是可以共享的。
	外蕴状态是不可以共享的，它随环境的改变而改变的，因此外蕴状态是由客户端来保持（因为环境的变化是由客户端引起的）。
	(1) 抽象享元角色：为具体享元角色规定了必须实现的方法，而外蕴状态就是以参数的形式通过此方法传入。在Java中可以由抽象类、接口来担当。
	(2) 具体享元角色：实现抽象角色规定的方法。如果存在内蕴状态，就负责为内蕴状态提供存储空间。
	(3) 享元工厂角色：负责创建和管理享元角色。要想达到共享的目的，这个角色的实现是关键！
	(4) 客户端角色：维护对所有享元对象的引用，而且还需要存储对应的外蕴状态。

使用场景：
	如果一个应用程序使用了大量的对象，而这些对象造成了很大的存储开销的时候就可以考虑是否可以使用享元模式。
	例如,如果发现某个对象的生成了大量细粒度的实例，并且这些实例除了几个参数外基本是相同的，如果把那些共享参数移到类外面，
在方法调用时将他们传递进来，就可以通过共享大幅度单个实例的数目。
*/

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}

func main() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()

	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight == viewer2.ImageFlyweight {
		fmt.Println("模式成功")
	}
}
