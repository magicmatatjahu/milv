package milv

type Hooks interface {
	onBefore()
	d(namespace string)
	onAfter(stats)
}