package bitpacking_test

import (
	"fmt"
	"testing"

	bitpacking "github.com/nikolaydubina/go-bit-packing"
)

func ExamplePack2x4b() {
	vs := [2]byte{
		0b0000_0111,
		0b0000_1110,
	}
	packed := bitpacking.Pack2x4b(vs)
	fmt.Printf("%08b\n", packed)
	// Output:
	// 11100111
}

func ExampleUnpack2x4b() {
	var packed byte = 0b11100111
	upacked := bitpacking.Unpack2x4b(packed)
	for _, q := range upacked {
		fmt.Printf("%08b\n", q)
	}
	// Output:
	// 00000111
	// 00001110
}

func FuzzExamplePack2x4b(f *testing.F) {
	f.Fuzz(func(t *testing.T, v0, v1 byte) {
		vs := [2]byte{v0, v1}

		for i := range vs {
			vs[i] = vs[i] & 0x0F
		}

		packed := bitpacking.Pack2x4b(vs)
		unpacked := bitpacking.Unpack2x4b(packed)

		if unpacked != vs {
			t.Errorf("exp(%v) != got (%v)", vs, unpacked)
		}
	})
}

func Benchmark_2x4b(b *testing.B) {
	unpacked := [...][2]byte{
		{0x0A, 0x0B},
		{0x02, 0x03},
		{0x0B, 0x05},
		{0x02, 0x07},
		{0x0F, 0x0C},
	}

	b.Run("pack", func(b *testing.B) {
		b.ResetTimer()
		var r byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Pack2x4b(unpacked[n%len(unpacked)])
			if r == 0 {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N*2)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})

	packed := [...]byte{
		0xAB,
		0x23,
		0xB5,
		0x27,
		0xFC,
	}

	b.Run("unpack", func(b *testing.B) {
		b.ResetTimer()
		var r [2]byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Unpack2x4b(packed[n%len(packed)])
			if r == [2]byte{} {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})
}

func ExamplePack4x6b() {
	vs := [4]byte{
		0b0010_0111,
		0b0001_1010,
		0b0011_1110,
		0b0001_0001,
	}
	packed := bitpacking.Pack4x6b(vs)
	for _, q := range packed {
		fmt.Printf("%08b\n", q)
	}
	// Output:
	// 10011010
	// 01111110
	// 11010001
}

func ExampleUnpack4x6b() {
	packed := [3]byte{
		0b1001_1010,
		0b0111_1110,
		0b1101_0001,
	}
	upacked := bitpacking.Unpack4x6b(packed)
	for _, q := range upacked {
		fmt.Printf("%08b\n", q)
	}
	// Output:
	// 00100111
	// 00011010
	// 00111110
	// 00010001
}

func FuzzExamplePack4x6b(f *testing.F) {
	f.Fuzz(func(t *testing.T, v0, v1, v2, v3 byte) {
		vs := [4]byte{v0, v1, v2, v3}

		// clear most significant bits
		for i := range vs {
			vs[i] = vs[i] & 0x3F
		}

		packed := bitpacking.Pack4x6b(vs)
		unpacked := bitpacking.Unpack4x6b(packed)

		if unpacked != vs {
			t.Errorf("exp(%v) != got (%v)", vs, unpacked)
		}
	})
}

func Benchmark_4x6b(b *testing.B) {
	unpacked := [...][4]byte{
		{0x1A, 0x2B, 0x3C, 0x0D},
		{0x02, 0x03, 0x04, 0x35},
		{0x02, 0x03, 0x04, 0x15},
		{0x1F, 0x3C, 0x0D, 0x0E},
	}

	b.Run("pack", func(b *testing.B) {
		b.ResetTimer()
		var r [3]byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Pack4x6b(unpacked[n%len(unpacked)])
			if r == [3]byte{} {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N*4)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})

	packed := [...][3]byte{
		{0xAB, 0x23, 0xFF},
		{0x23, 0x45, 0x67},
		{0x12, 0x34, 0x56},
		{0x78, 0x9A, 0xBC},
	}

	b.Run("unpack", func(b *testing.B) {
		b.ResetTimer()
		var r [4]byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Unpack4x6b(packed[n%len(packed)])
			if r == [4]byte{} {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N*3)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})
}

func ExamplePack8x7b() {
	vs := [8]byte{
		0b0001_0010,
		0b0010_0101,
		0b0011_0110,
		0b0100_1001,
		0b0101_1010,
		0b0110_1100,
		0b0111_1110,
		0b0001_1111,
	}
	packed := bitpacking.Pack8x7b(vs)
	for _, p := range packed {
		fmt.Printf("%08b\n", p)
	}
	// Output:
	// 00100101
	// 00110110
	// 11001001
	// 01011010
	// 01101100
	// 11111110
	// 00011111
}

func ExampleUnpack8x7b() {
	vs := [7]byte{
		0b00100101,
		0b00110110,
		0b11001001,
		0b01011010,
		0b01101100,
		0b11111110,
		0b00011111,
	}
	packed := bitpacking.Unpack8x7b(vs)
	for _, p := range packed {
		fmt.Printf("%08b\n", p)
	}
	// Output:
	// 00010010
	// 00100101
	// 00110110
	// 01001001
	// 01011010
	// 01101100
	// 01111110
	// 00011111
}

func FuzzExamplePack8x7b(f *testing.F) {
	f.Fuzz(func(t *testing.T, v0, v1, v2, v3, v4, v5, v6, v7 byte) {
		vs := [8]byte{v0, v1, v2, v3, v4, v5, v6, v7}

		// clear most significant bit
		for i := range vs {
			vs[i] = vs[i] & 0x7F
		}

		packed := bitpacking.Pack8x7b(vs)
		unpacked := bitpacking.Unpack8x7b(packed)

		if unpacked != vs {
			t.Errorf("exp(%v) != got (%v)", vs, unpacked)
		}
	})
}

func Benchmark_8x7b(b *testing.B) {
	unpacked := [...][8]byte{
		{0x1A, 0x2B, 0x3C, 0x1D, 0x1E, 0x2F, 0x3A, 0x7B},
		{0x11, 0x22, 0x33, 0x14, 0x15, 0x26, 0x37, 0x78},
	}

	b.Run("pack", func(b *testing.B) {
		b.ResetTimer()
		var r [7]byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Pack8x7b(unpacked[n%len(unpacked)])
			if r == [7]byte{} {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N*8)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})

	packed := [...][7]byte{
		{0xAB, 0x23, 0xFF, 0xF1, 0xA2, 0x92, 0x91},
		{0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78},
		{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE},
	}

	b.Run("unpack", func(b *testing.B) {
		b.ResetTimer()
		var r [8]byte
		for n := 0; n < b.N; n++ {
			r = bitpacking.Unpack8x7b(packed[n%len(packed)])
			if r == [8]byte{} {
				b.Error(r)
			}
		}
		b.ReportMetric(float64(b.N*7)/b.Elapsed().Seconds()/(1<<30), "GB/s")
	})
}
