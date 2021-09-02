package singleton

type Counter struct {
	count int
}

func (c *Counter) AddOne() {
	c.count++
}
func (c *Counter) Get() int {
	return c.count
}

var instance *Counter

func GetInstance() *Counter {
	if instance == nil {
		instance = new(Counter)
	}
	return instance
}
