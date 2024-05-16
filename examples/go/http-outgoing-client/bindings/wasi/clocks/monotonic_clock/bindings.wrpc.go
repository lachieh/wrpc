// Generated by `wit-bindgen-wrpc-go` 0.1.0. DO NOT EDIT!
package monotonic_clock

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	errors "errors"
	fmt "fmt"
	wasi__io__poll "github.com/wrpc/wrpc/examples/go/http-outgoing-client/bindings/wasi/io/poll"
	wrpc "github.com/wrpc/wrpc/go"
	io "io"
	slog "log/slog"
	utf8 "unicode/utf8"
)

type Pollable = wasi__io__poll.Pollable

// An instant in time, in nanoseconds. An instant is relative to an
// unspecified initial value, and can only be compared to instances from
// the same monotonic-clock.
type Instant = uint64

// A duration of time, in nanoseconds.
type Duration = uint64

// Read the current value of the clock.
//
// The clock is monotonic, therefore calling this function repeatedly will
// produce a sequence of non-decreasing values.
func Now(ctx__ context.Context, wrpc__ wrpc.Client) (r0__ uint64, close__ func() error, err__ error) {
	if err__ = wrpc__.Invoke(ctx__, "wasi:clocks/monotonic-clock@0.2.0", "now", func(w__ wrpc.IndexWriter, r__ wrpc.IndexReadCloser) error {
		close__ = r__.Close
		_, err__ = w__.Write(nil)
		if err__ != nil {
			return fmt.Errorf("failed to write empty parameters: %w", err__)
		}
		r0__, err__ = func() (Instant, error) {
			v, err := func(r io.ByteReader) (uint64, error) {
				var x uint64
				var s uint
				for i := 0; i < 10; i++ {
					slog.Debug("reading u64 byte", "i", i)
					b, err := r.ReadByte()
					if err != nil {
						if i > 0 && err == io.EOF {
							err = io.ErrUnexpectedEOF
						}
						return x, fmt.Errorf("failed to read u64 byte: %w", err)
					}
					if b < 0x80 {
						if i == 9 && b > 1 {
							return x, errors.New("varint overflows a 64-bit integer")
						}
						return x | uint64(b)<<s, nil
					}
					x |= uint64(b&0x7f) << s
					s += 7
				}
				return x, errors.New("varint overflows a 64-bit integer")
			}(r__)
			return (Instant)(v), err
		}()

		if err__ != nil {
			return fmt.Errorf("failed to read result 0: %w", err__)
		}
		return nil
	}); err__ != nil {
		err__ = fmt.Errorf("failed to invoke `now`: %w", err__)
		return
	}
	return
}

// Query the resolution of the clock. Returns the duration of time
// corresponding to a clock tick.
func Resolution(ctx__ context.Context, wrpc__ wrpc.Client) (r0__ uint64, close__ func() error, err__ error) {
	if err__ = wrpc__.Invoke(ctx__, "wasi:clocks/monotonic-clock@0.2.0", "resolution", func(w__ wrpc.IndexWriter, r__ wrpc.IndexReadCloser) error {
		close__ = r__.Close
		_, err__ = w__.Write(nil)
		if err__ != nil {
			return fmt.Errorf("failed to write empty parameters: %w", err__)
		}
		r0__, err__ = func() (Duration, error) {
			v, err := func(r io.ByteReader) (uint64, error) {
				var x uint64
				var s uint
				for i := 0; i < 10; i++ {
					slog.Debug("reading u64 byte", "i", i)
					b, err := r.ReadByte()
					if err != nil {
						if i > 0 && err == io.EOF {
							err = io.ErrUnexpectedEOF
						}
						return x, fmt.Errorf("failed to read u64 byte: %w", err)
					}
					if b < 0x80 {
						if i == 9 && b > 1 {
							return x, errors.New("varint overflows a 64-bit integer")
						}
						return x | uint64(b)<<s, nil
					}
					x |= uint64(b&0x7f) << s
					s += 7
				}
				return x, errors.New("varint overflows a 64-bit integer")
			}(r__)
			return (Duration)(v), err
		}()

		if err__ != nil {
			return fmt.Errorf("failed to read result 0: %w", err__)
		}
		return nil
	}); err__ != nil {
		err__ = fmt.Errorf("failed to invoke `resolution`: %w", err__)
		return
	}
	return
}

// Create a `pollable` which will resolve once the specified instant
// occured.
func SubscribeInstant(ctx__ context.Context, wrpc__ wrpc.Client, when uint64) (r0__ wrpc.Own[Pollable], close__ func() error, err__ error) {
	if err__ = wrpc__.Invoke(ctx__, "wasi:clocks/monotonic-clock@0.2.0", "subscribe-instant", func(w__ wrpc.IndexWriter, r__ wrpc.IndexReadCloser) error {
		close__ = r__.Close
		var buf__ bytes.Buffer
		writes__ := make(map[uint32]func(wrpc.IndexWriter) error, 1)
		write0__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v uint64, w io.Writer) (err error) {
			b := make([]byte, binary.MaxVarintLen64)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing u64")
			_, err = w.Write(b[:i])
			return err
		}(when, &buf__)
		if err__ != nil {
			return fmt.Errorf("failed to write `when` parameter: %w", err__)
		}
		if write0__ != nil {
			writes__[0] = write0__
		}
		_, err__ = w__.Write(buf__.Bytes())
		if err__ != nil {
			return fmt.Errorf("failed to write parameters: %w", err__)
		}
		r0__, err__ = func(r interface {
			io.ByteReader
			io.Reader
		}) (wrpc.Own[Pollable], error) {
			var x uint32
			var s uint
			for i := 0; i < 5; i++ {
				slog.Debug("reading owned resource ID length byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return "", fmt.Errorf("failed to read owned resource ID length byte: %w", err)
				}
				if b < 0x80 {
					if i == 4 && b > 1 {
						return "", errors.New("owned resource ID length overflows a 32-bit integer")
					}
					x = x | uint32(b)<<s
					buf := make([]byte, x)
					slog.Debug("reading owned resource ID bytes", "len", x)
					_, err = r.Read(buf)
					if err != nil {
						return "", fmt.Errorf("failed to read owned resource ID bytes: %w", err)
					}
					if !utf8.Valid(buf) {
						return "", errors.New("owned resource ID is not valid UTF-8")
					}
					return wrpc.Own[Pollable](buf), nil
				}
				x |= uint32(b&0x7f) << s
				s += 7
			}
			return "", errors.New("owned resource ID length overflows a 32-bit integer")
		}(r__)
		if err__ != nil {
			return fmt.Errorf("failed to read result 0: %w", err__)
		}
		return nil
	}); err__ != nil {
		err__ = fmt.Errorf("failed to invoke `subscribe-instant`: %w", err__)
		return
	}
	return
}

// Create a `pollable` which will resolve once the given duration has
// elapsed, starting at the time at which this function was called.
// occured.
func SubscribeDuration(ctx__ context.Context, wrpc__ wrpc.Client, when uint64) (r0__ wrpc.Own[Pollable], close__ func() error, err__ error) {
	if err__ = wrpc__.Invoke(ctx__, "wasi:clocks/monotonic-clock@0.2.0", "subscribe-duration", func(w__ wrpc.IndexWriter, r__ wrpc.IndexReadCloser) error {
		close__ = r__.Close
		var buf__ bytes.Buffer
		writes__ := make(map[uint32]func(wrpc.IndexWriter) error, 1)
		write0__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v uint64, w io.Writer) (err error) {
			b := make([]byte, binary.MaxVarintLen64)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing u64")
			_, err = w.Write(b[:i])
			return err
		}(when, &buf__)
		if err__ != nil {
			return fmt.Errorf("failed to write `when` parameter: %w", err__)
		}
		if write0__ != nil {
			writes__[0] = write0__
		}
		_, err__ = w__.Write(buf__.Bytes())
		if err__ != nil {
			return fmt.Errorf("failed to write parameters: %w", err__)
		}
		r0__, err__ = func(r interface {
			io.ByteReader
			io.Reader
		}) (wrpc.Own[Pollable], error) {
			var x uint32
			var s uint
			for i := 0; i < 5; i++ {
				slog.Debug("reading owned resource ID length byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return "", fmt.Errorf("failed to read owned resource ID length byte: %w", err)
				}
				if b < 0x80 {
					if i == 4 && b > 1 {
						return "", errors.New("owned resource ID length overflows a 32-bit integer")
					}
					x = x | uint32(b)<<s
					buf := make([]byte, x)
					slog.Debug("reading owned resource ID bytes", "len", x)
					_, err = r.Read(buf)
					if err != nil {
						return "", fmt.Errorf("failed to read owned resource ID bytes: %w", err)
					}
					if !utf8.Valid(buf) {
						return "", errors.New("owned resource ID is not valid UTF-8")
					}
					return wrpc.Own[Pollable](buf), nil
				}
				x |= uint32(b&0x7f) << s
				s += 7
			}
			return "", errors.New("owned resource ID length overflows a 32-bit integer")
		}(r__)
		if err__ != nil {
			return fmt.Errorf("failed to read result 0: %w", err__)
		}
		return nil
	}); err__ != nil {
		err__ = fmt.Errorf("failed to invoke `subscribe-duration`: %w", err__)
		return
	}
	return
}
