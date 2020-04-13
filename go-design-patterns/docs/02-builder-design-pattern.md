# Builder - 构造者模式

顾名思义, 构造者模式(Builder Design Pattern)帮助我们构造一些不方便直接实例化复杂的对象.
比如一些对象可能有十多个字段, 我们不想一个个编写这些对象的实例化逻辑, 只想使用这些对象.

实例化一个对象, 可以用"{}"简单初始化, 字段值都为0值; 也可以通过一定的API, 根据一定的逻辑构建对象.
还有一些对象由其他对象组成, 而Go语言不支持继承.

有时候要使用同样的方法构建一系列类似的对象, 比如按同样的方法构建一批汽车, 这些汽车只有车牌号等
少量特征不同, 其他方面都是一致的. 
这时候, 我们就要考虑使用相同的工序是生产汽车了.

这就是构造者模式的使用场景.

## 目标
构造者模式致力于:
- 抽象出复杂的构造方法, 构造方法与使用者分离;
- 通过填写字段和嵌入对象, 一步一步创建一个对象;
- 在不同对象中复用对象的构造算法.

## 示例 - 汽车生产商

### 要求与验收标准

### 示例代码
> builder/creational/creational.go
```go
package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (md *ManufacturingDirector) Construct() {
	md.builder.SetSeats().SetStructure().SetWheels()
}

func (md *ManufacturingDirector) SetBuilder(b BuildProcess) {
	md.builder = b
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

type CarBuilder struct {
	vp VehicleProduct
}

func (cb *CarBuilder) SetWheels() BuildProcess {
	cb.vp.Wheels = 4
	return cb
}

func (cb *CarBuilder) SetSeats() BuildProcess {
	cb.vp.Seats = 5
	return nil
}

func (cb *CarBuilder) SetStructure() BuildProcess {
	cb.vp.Structure = "Car"
	return cb
}

func (cb *CarBuilder) GetVehicle() VehicleProduct {
	return cb.vp
}

func (cb *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

type BikeBuilder struct {
	vp VehicleProduct
}

func (bb *BikeBuilder) SetWheels() BuildProcess {
	bb.vp.Wheels = 2
	return bb
}

func (bb *BikeBuilder) SetSeats() BuildProcess {
	bb.vp.Seats = 2
	return bb
}

func (bb *BikeBuilder) SetStructure() BuildProcess {
	bb.vp.Structure = "Motorbike"
	return bb
}

func (bb *BikeBuilder) GetVehicle() VehicleProduct {
	return bb.vp
}

func (bb *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

type BusBuilder struct {
	vp VehicleProduct
}

func (bb *BusBuilder) SetWheels() BuildProcess {
	bb.vp.Wheels = 4 * 2
	return bb
}

func (bb *BusBuilder) SetSeats() BuildProcess {
	bb.vp.Seats = 30
	return bb
}

func (bb *BusBuilder) SetStructure() BuildProcess {
	bb.vp.Structure = "Bus"
	return bb
}

func (bb *BusBuilder) GetVehicle() VehicleProduct {
	return bb.vp
}


```