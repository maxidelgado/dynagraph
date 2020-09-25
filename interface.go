package dynagraph

type Graph interface {
	Node(id ...string) Node
	Query(filter Filter) Query
	Batch() Batch
	Transaction() Transaction
}

type Node interface {
	Put(value interface{}) (string, error)
	Update(value interface{}) error
	Edge(value interface{}) error
	Prop(value interface{}) error
	Ref(id string) error
	Delete(filter Filter) error
}

type Batch interface {
	Get(out interface{}, keys KeysInput) error
	Put(inputs WriteItemsInput) (int, error)
	Delete(keys KeysInput) (int, error)
}

type Query interface {
	One(out interface{}) error
	All(out interface{}) error
}

type Transaction interface {
	Put(inputs WriteItemsInput) Transaction
	Update(inputs WriteItemsInput) Transaction
	Delete(keys KeysInput) Transaction
	Run() error
}
