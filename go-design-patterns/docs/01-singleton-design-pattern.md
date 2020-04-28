# Singleton - 单例模式
> 整个程序中某类型的唯一实例

一个有趣的现象, 当面试官问及设计模式的时候, 80%的面试者都会提及"单例模式".
为啥呢? 也许是因为这个最常用的一种设计模式, 也可能是因为这个最简单的一种设计模式.
我们(go design patterns)之所以把单例模式放到第一个来讲,
就是出于后一个原因.

## 描述
单例模式很容易记.
正如其名, 单例模式提供某类的唯一实例, 并且保证这个实例的唯一性.

在第一次使用这个实例时, 这个实例就已经被创建了,
接下来这个实例可复用于程序的任何部分.

你可以将单例模式用于不同的情景, 比如:
- 在同一个数据库连接上处理每一个请求时;
- 用Secure Shell(SSH)连接到服务器上执行一些任务, 但并不想为每一个任务都单独打开一个连接;
- 想要限制对某些变量或空间的访问, 把单例作为这个变量的"门"(不过在Go语言中更推荐使用channel处理这种场景);
- ...

## Objectives
当满是以下条件时, 我们可以考虑单例模式:
- 我们需要一个某特定类型的独一的、共享的值;
- 我们要求某个对象在整个程序中是唯一的。

### 示例 - 唯一的计数器
我们编写一个计数器来计算某个方法被执行的次数, 这个计数器要保持全局的唯一性.

编写这样的计数器有以下要求:
- 如果计数器还没有被创建, 那么第一次使用时会创建, 且计数器的初始值为0;
- 如果计数器已经被创建了, 返回的实例包含其真实的值;
- 当调用`AddOne`方法时, 计数器的值要加1 .

### 代码
Go语言实现单例模式的方法与纯OO(object-oriented, 面向对象)的语言(如Java, C++, 有静态成员)有所不同,
Go语言中, 并没有静态成员的概念, 但是Go语言中的包级别变量, 可以做到类似的效果.

我创建了一个go module: [github.com/dspo/gopherway/go-design-patterns](/go-design-paterns),
module下建了一个目录[code/singleton](/go-design-patterns/code/singleton),
用来编写单例模式的示例代码.

> /go-design-patterns/code/singleton/singleton.go
```go
package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
} 

var instance *singleton

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
```

### sync.Once
`sync.Once` 是Go语言并发组件中常用工具, 其参数是传入一个函数, Go语言层面保证这个函数只被执行一次.
> /go-design-patterns/code/singleton/once.go
```go
package singleton

import (
	"sync"
)

var o sync.Once

func GetInstance2() Singleton {
	o.Do(func() {
		instance = new(singleton)
	})
	return instance
}
```
如上代码中, GetInstance2也能保证获取是示例永远是同一个示例. 
 两种实现单例的方式可任选一种, 但不要混用, 混用后就不能保证单例了.
 
注意示例代码中用`o`命名变量, 不代表推荐这种明明风格.