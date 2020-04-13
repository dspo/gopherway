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



