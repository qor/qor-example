package cart

type CartBucket interface {
	Restore() (map[uint]*CartItem, error)
	Save(map[uint]*CartItem) error
}
