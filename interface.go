package goutil

import "io"

type writer struct{ write func([]byte) (int, error) }

func NewWriter(write func([]byte) (int, error)) io.Writer { return &writer{write: write} }
func (x *writer) Write(p []byte) (n int, err error)       { return x.write(p) }

var _ io.Writer = (*writer)(nil)

type reader struct{ read func([]byte) (int, error) }

func NewReader(read func([]byte) (int, error)) io.Reader { return &reader{read: read} }
func (x *reader) Read(p []byte) (n int, err error)       { return x.read(p) }

var _ io.Reader = (*reader)(nil)

type closer struct{ close func() error }

func NewCloser(close func() error) io.Closer { return &closer{close: close} }
func (x *closer) Close() error               { return x.close() }

var _ io.Closer = (*closer)(nil)
