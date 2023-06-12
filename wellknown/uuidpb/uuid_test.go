package uuidpb_test

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/dogmatiq/interopspec/wellknown/uuidpb"
	"google.golang.org/protobuf/proto"
)

func TestUUID_New(t *testing.T) {
	t.Parallel()

	uuid := New()

	if proto.Equal(uuid, &UUID{}) {
		t.Fatal("got 'nil' UUID (all zeroes), want random UUID")
	}

	omni := &UUID{
		Upper: 0xffffffffffffffff,
		Lower: 0xffffffffffffffff,
	}

	if proto.Equal(uuid, omni) {
		t.Fatal("got 'omni' UUID (all ones), want random UUID")
	}

	expect := uint64(0x40)
	actual := (uuid.Upper >> 8) & 0xf0

	if actual != expect {
		t.Fatalf("got version %d, want %d", actual, expect)
	}

	expect = 0x80
	actual = (uuid.Lower >> 56) & 0xc0

	if actual != expect {
		t.Fatalf("got variant %d, want %d (RFC 4122)", actual, expect)
	}
}

func TestUUID_FromBytes(t *testing.T) {
	t.Parallel()

	data := []byte{
		0xa9, 0x67, 0xa8, 0xb9,
		0x3f, 0x9c, 0x49, 0x18,
		0x9a, 0x41, 0x19, 0x57,
		0x7b, 0xe5, 0xfe, 0xc5,
	}

	actual := FromBytes(data)
	expect := &UUID{
		Upper: 0xa967a8b93f9c4918,
		Lower: 0x9a4119577be5fec5,
	}

	if !proto.Equal(actual, expect) {
		t.Fatalf("got %s, want %s", actual, expect)
	}
}

func TestUUID_ToBytes(t *testing.T) {
	t.Parallel()

	uuid := &UUID{
		Upper: 0xa967a8b93f9c4918,
		Lower: 0x9a4119577be5fec5,
	}

	actual := uuid.ToBytes()
	expect := []byte{
		0xa9, 0x67, 0xa8, 0xb9,
		0x3f, 0x9c, 0x49, 0x18,
		0x9a, 0x41, 0x19, 0x57,
		0x7b, 0xe5, 0xfe, 0xc5,
	}

	if !bytes.Equal(actual, expect) {
		t.Fatalf("got %#v, want %#v", actual, expect)
	}
}

func TestUUID_ToString(t *testing.T) {
	t.Parallel()

	uuid := &UUID{
		Upper: 0xa967a8b93f9c4918,
		Lower: 0x9a4119577be5fec5,
	}

	expect := "a967a8b9-3f9c-4918-9a41-19577be5fec5"
	actual := uuid.ToString()

	if actual != expect {
		t.Fatalf("got %q, want %q", actual, expect)
	}
}

func TestUUID_Format(t *testing.T) {
	t.Parallel()

	uuid := &UUID{
		Upper: 0xa967a8b93f9c4918,
		Lower: 0x9a4119577be5fec5,
	}

	expect := uuid.ToString()
	actual := fmt.Sprintf("%s", uuid)

	if actual != expect {
		t.Fatalf("got %q, want %q", actual, expect)
	}
}
