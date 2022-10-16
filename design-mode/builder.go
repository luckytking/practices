package design_mode

type ResourcePool struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewResourcePool1(name string, maxTotal int, maxIdle int, minIdle int) *ResourcePool {
	this := &ResourcePool{
		name:     name,
		maxTotal: maxTotal,
		maxIdle:  maxIdle,
		minIdle:  minIdle,
	}
	return this
}

func NewResourcePool2(name string) *ResourcePool {
	this := &ResourcePool{
		name: name,
	}
	return this
}

func (this *ResourcePool) SetName(name string) *ResourcePool {
	if len(name) <= 0 {
		panic(any("name should not be empty."))
	}
	this.name = name
	return this
}

func (this *ResourcePool) SetMaxTotal(maxTotal int) *ResourcePool {
	if maxTotal <= 0 {
		panic(any("maxTotal should be positive."))
	}
	this.maxTotal = maxTotal
	return this
}

func (this *ResourcePool) SetMinIdle(minIdle int) *ResourcePool {
	if minIdle < 0 {
		panic(any("minIdle should not be negative."))
	}
	this.minIdle = minIdle
	return this
}

func NewResourcePool(builder *ResourcePoolBuilder) *ResourcePool {
	this := &ResourcePool{}
	this.name = builder.name
	this.maxTotal = builder.maxTotal
	this.maxIdle = builder.maxIdle
	this.minIdle = builder.minIdle
	return this
}

type ResourcePoolBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func Builder() *ResourcePoolBuilder {
	builder := &ResourcePoolBuilder{}
	builder.maxTotal = 16
	builder.minIdle = 5
	return builder
}

func (b *ResourcePoolBuilder) Build() *ResourcePool {
	// 校验逻辑放到这里来做，包括必填项校验、依赖关系校验、约束条件校验等
	if len(b.name) <= 0 {
		panic(any("name should not be empty."))
	}
	if b.maxIdle > b.maxTotal {
		panic(any("maxIdle should bigger than maxTotal."))
	}
	return NewResourcePool(b)
}

func (this *ResourcePoolBuilder) SetName(name string) *ResourcePoolBuilder {
	if len(name) <= 0 {
		panic(any("name should not be empty."))
	}
	this.name = name
	return this
}

func (this *ResourcePoolBuilder) SetMaxTotal(maxTotal int) *ResourcePoolBuilder {
	if maxTotal <= 0 {
		panic(any("maxTotal should be positive."))
	}
	this.maxTotal = maxTotal
	return this
}

func (this *ResourcePoolBuilder) SetMinIdle(minIdle int) *ResourcePoolBuilder {
	if minIdle < 0 {
		panic(any("minIdle should not be negative."))
	}
	this.minIdle = minIdle
	return this
}
