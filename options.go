package main

type Key int64

func (k Key) Int64() int64 {
	return int64(k)
}

type Opts interface {
	Int64() int64
	GetKey() Key
}

type Opt struct {
	i    Key
	Func func(*Files)
}

func (o Opt) Int64() int64 {
	return int64(o.i)
}

func (o Opt) GetKey() Key {
	return o.i
}

type SortOpt struct {
	i         Key
	Func      func(*Files)
	Direction SortDirection
}

func (o SortOpt) Int64() int64 {
	return int64(o.i)
}
func (o SortOpt) GetKey() Key {
	return o.i
}
