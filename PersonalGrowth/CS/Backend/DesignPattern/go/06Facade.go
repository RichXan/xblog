package designpattern

// 子系统1：CPU
type CPU struct{}

func (c *CPU) Start() {
	println("CPU is starting")
}

func (c *CPU) Execute() {
	println("CPU is executing")
}

func (c *CPU) Shutdown() {
	println("CPU is shutting down")
}

// 子系统2：Memory
type Memory struct{}

func (m *Memory) Load() {
	println("Memory is loading")
}

func (m *Memory) Unload() {
	println("Memory is unloading")
}

// 子系统3：HardDrive
type HardDrive struct{}

func (h *HardDrive) Read() {
	println("HardDrive is reading")
}

func (h *HardDrive) Write() {
	println("HardDrive is writing")
}

// ComputerFacade 提供了一个统一的接口
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

// NewComputerFacade 创建外观
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Start 提供了一个简单的接口来启动计算机
func (c *ComputerFacade) Start() {
	c.cpu.Start()
	c.memory.Load()
	c.hardDrive.Read()
	c.cpu.Execute()
}

// Shutdown 提供了一个简单的接口来关闭计算机
func (c *ComputerFacade) Shutdown() {
	c.cpu.Shutdown()
	c.memory.Unload()
}
