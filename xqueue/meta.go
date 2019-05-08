package xqueue

import "errors"

type Queue interface {
	Push(topic, data []byte) error
	Pull(topic []byte) (data []byte, err error)
}



var ErrNil = errors.New("queue: nil returned")