package designpattern

// Subject 定义代理和实际对象的共同接口
type Subject interface {
	Do() string
}

// RealSubject 是实际对象
type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "RealSubject: doing something"
}

// Proxy 是代理对象
type Proxy struct {
	realSubject *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Do() string {
	// 在调用实际对象之前的处理
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}

	// 调用实际对象
	result := p.realSubject.Do()

	// 在调用实际对象之后的处理
	return "Proxy: " + result
}

// 实际应用示例：带缓存的代理
type ImageProxy struct {
	realImage *Image
	filename  string
	cache     map[string]string
}

type Image struct {
	filename string
}

func (i *Image) Display() string {
	return "Displaying " + i.filename
}

func NewImageProxy(filename string) *ImageProxy {
	return &ImageProxy{
		filename: filename,
		cache:    make(map[string]string),
	}
}

func (p *ImageProxy) Display() string {
	// 检查缓存
	if result, ok := p.cache[p.filename]; ok {
		return "Cached: " + result
	}

	// 延迟加载
	if p.realImage == nil {
		p.realImage = &Image{p.filename}
	}

	// 执行实际操作并缓存结果
	result := p.realImage.Display()
	p.cache[p.filename] = result

	return result
}
