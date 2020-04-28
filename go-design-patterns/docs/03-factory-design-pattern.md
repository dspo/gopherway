# 工厂方法模式
工厂模式可能是除了单例模式以外, 最有名的设计模式了.
工厂模式致力于将用户的需求抽象出来, 使得创建对象和使用对象是分离的, 用户总是引用抽象工厂和抽象产品.

比如从web或数据库取得数据, 用户需要的仅仅是一个输出数据的接口, 而不要管提供接口者的实现.
 用户不需要知道数据是从web取得的, 还是数据库取得的.
 
## 描述
当我们使用工厂模式, 我们就获得了一个额外的封装层, 程序的增长更加可控.

通过工程模式, 我们可以将相似的对象的创建方法分离到不同的package或结构体中.
想象一下, 雇佣一个行程代理人来管理你的旅游, 你不用自己考虑理酒店和路线这些琐碎的事情, 
你只要告诉代理人你的目的地就可以了.
你的代理人会提供一切你需要的. 
这个代理人就是旅行工厂.

再比如你要使用一个键值对缓存结构, 你不关心数据被存放到了哪里, 
可以是在内存中, 可以是在ETCD中, 也可以是在Redis中.
你只关心这个结构提供了
- `Store(key, value interface{})` 存放键值对 
- `Load(key interface{}) (value interface{})` 取出键值对
- `Delete(key interface{})` 删除键值对

等方法即可.

通过上面的描述, 对于工厂方法模式, 你应当清楚:
- 新建实例的方法对于用户是处于分离状态的
- 用户只使用interface, 而不是用具体的实现, 一般情况下, 用户也不应该了解具体的实现

在Jave语言中, 如果用户想得到具体的实现结构, 可以使用静态方法`getFactory()`来返回真正的子类.
在Go语言中, 用户可以使用类型断言获取真正的子类(如果创建方法的设计者允许的话).

## 示例代码
以下示例代码描述了某支付场景. 
首先定义了支付方式接口, 再实现了两种支付方式.
用户可以选择使用任意一种支付方式, 只需在调用`GetPaymentMethod`函数时, 传入要使用的支付方式即可.
当然, 当业务场景更加丰富时, 用户也可以实现自己的支付方式, 
如使要新增用微信付款, 只需定义一个`type WechatPM struct`.

> /go-design-patterns/code/factory/factory.go
```go
package factory

import (
	"errors"
	"strconv"
)

const (
	Cash = iota+1
	DebitCard
)

// 支付方式
type PaymentMethod interface {
	Pay(amount int) string
}

type CashPM struct {

}

func (c *CashPM) Pay(amount int) string {
	return "you pay" + strconv.Itoa(amount) + "by CASH payment"
}

type DebitCardPM struct {

}

func (d *DebitCardPM) Pay(amount int) string {
	return "you pay" + strconv.Itoa(amount) + "by DEBIT CARD payment"
}

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, errors.New("payment method not recognized")
	}
}

```

利用工厂模式, 可以实现插件化开发.