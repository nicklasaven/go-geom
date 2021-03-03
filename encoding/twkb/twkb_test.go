package twkb

import (
//	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
	"github.com/twpayne/go-geom/internal/geomtest"
//	"github.com/twpayne/go-geom/internal/testdata"
)

func test(t *testing.T, g geom.T, precision, twkb []byte, opts ...wkbcommon.WKBOption) {

		t.Run("twkb", func(t *testing.T) {
			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(twkb, opts...)
				require.NoError(t, err)
				require.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, precision[0], precision[1], precision[2], opts...)
				require.NoError(t, err)
				require.Equal(t, twkb, got)
			})
		})


/*
	t.Run("scan", func(t *testing.T) {
		switch g := g.(type) {
		case *geom.Point:
			p := Point{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, p.Scan(xdr))
					require.Equal(t, Point{g, opts}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, p.Scan(ndr))
					require.Equal(t, Point{g, opts}, p)
				})
			}
		case *geom.LineString:
			ls := LineString{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, ls.Scan(xdr))
					require.Equal(t, LineString{g, opts}, ls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, ls.Scan(ndr))
					require.Equal(t, LineString{g, opts}, ls)
				})
			}
		case *geom.Polygon:
			p := Polygon{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, p.Scan(xdr))
					require.Equal(t, Polygon{g, opts}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, p.Scan(ndr))
					require.Equal(t, Polygon{g, opts}, p)
				})
			}
		case *geom.MultiPoint:
			mp := MultiPoint{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, mp.Scan(xdr))
					require.Equal(t, MultiPoint{g, opts}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, mp.Scan(ndr))
					require.Equal(t, MultiPoint{g, opts}, mp)
				})
			}
		case *geom.MultiLineString:
			mls := MultiLineString{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, mls.Scan(xdr))
					require.Equal(t, MultiLineString{g, opts}, mls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, mls.Scan(ndr))
					require.Equal(t, MultiLineString{g, opts}, mls)
				})
			}
		case *geom.MultiPolygon:
			mp := MultiPolygon{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, mp.Scan(xdr))
					require.Equal(t, MultiPolygon{g, opts}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, mp.Scan(ndr))
					require.Equal(t, MultiPolygon{g, opts}, mp)
				})
			}
		case *geom.GeometryCollection:
			gc := GeometryCollection{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					require.NoError(t, gc.Scan(xdr))
					require.Equal(t, GeometryCollection{g, opts}, gc)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					require.NoError(t, gc.Scan(ndr))
					require.Equal(t, GeometryCollection{g, opts}, gc)
				})
			}
		}
	})
	*/
}

func Test(t *testing.T) {
	for _, tc := range []struct {
		g    geom.T
		opts []wkbcommon.WKBOption
		precision []byte
		twkb  []byte
	}{
/*
		{
			g:    geom.NewPointEmpty(geom.XY),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  geomtest.MustHexDecode("00000000017ff80000000000007ff8000000000000"),
			ndr:  geomtest.MustHexDecode("0101000000000000000000f87f000000000000f87f"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYM),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  geomtest.MustHexDecode("00000007d17ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  geomtest.MustHexDecode("01d1070000000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYZ),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  geomtest.MustHexDecode("00000003e97ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  geomtest.MustHexDecode("01e9030000000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYZM),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  geomtest.MustHexDecode("0000000bb97ff80000000000007ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  geomtest.MustHexDecode("01b90b0000000000000000f87f000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    geom.NewGeometryCollection().MustPush(geom.NewPointEmpty(geom.XY)),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  geomtest.MustHexDecode("00000000070000000100000000017ff80000000000007ff8000000000000"),
			ndr:  geomtest.MustHexDecode("0107000000010000000101000000000000000000f87f000000000000f87f"),
		},
*/
		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{1, 2}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("01000204"),
		},

		{
			g:   geom.NewPoint(geom.XYZ).MustSetCoords(geom.Coord{1, 2, 3}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("010801020406"),
		},
		{
			g:   geom.NewPoint(geom.XYM).MustSetCoords(geom.Coord{1, 2, 3}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("010802020406"),
		},
		{
			g:   geom.NewPoint(geom.XYZM).MustSetCoords(geom.Coord{1, 2, 3, 4}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("01080302040608"),
		},
		{
			g:   geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{1, 2}, {3, 4}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("02000202040404"),
		},
		{
			g:   geom.NewLineString(geom.XYZ).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("02080102020406060606"),
		},
		{
			g:   geom.NewLineString(geom.XYM).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("02080202020406060606"),
		},
		{
			g:   geom.NewLineString(geom.XYZM).MustSetCoords([]geom.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("020803020204060808080808"),
		},
		{
			g:   geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("030001040204040404040707"),
		},
		{
			g:   geom.NewPolygon(geom.XYZ).MustSetCoords([][]geom.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("03080101040204060606060606060b0b0b"),
		},
		{
			g:   geom.NewPolygon(geom.XYM).MustSetCoords([][]geom.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("03080201040204060606060606060b0b0b"),
		},
		{
			g:   geom.NewPolygon(geom.XYZM).MustSetCoords([][]geom.Coord{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("03080301040204060808080808080808080f0f0f0f"),
		},
/*		{
			g:   geom.NewMultiPoint(geom.XY).MustSetCoords([]geom.Coord{{1, 2}, {3, 4}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("04000202040404"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYZ).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("04080102020406060606"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYM).MustSetCoords([]geom.Coord{{1, 2, 3}, {4, 5, 6}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("04080202020406060606"),
		},
		{
			g:   geom.NewMultiPoint(geom.XYZM).MustSetCoords([]geom.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("040803020204060808080808"),
		},
		{
			g:    geom.NewMultiPoint(geom.XY).MustSetCoords([]geom.Coord{nil, {1, 2}, {3, 4}}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("0104000000030000000101000000000000000000f87f000000000000f87f0101000000000000000000f03f0000000000000040010100000000000000000008400000000000001040"),
		},
		{
			g:   geom.NewGeometryCollection(),
			xdr: geomtest.MustHexDecode("000000000700000000"),
			ndr: geomtest.MustHexDecode("010700000000000000"),
		},
		{
			g: geom.NewGeometryCollection().MustPush(
				geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{-79.3698576, 43.6456613}),
				geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{-79.3707986, 43.6453697}, {-79.3704747, 43.6454819}, {-79.370186, 43.6455592}, {-79.3699323, 43.6456385}, {-79.3698576, 43.6456613}}),
				geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{{-79.3698576, 43.6456613}, {-79.3698057, 43.6455265}}),
			),
			xdr: geomtest.MustHexDecode("0000000007000000030000000001c053d7abbf360b554045d2a5078be57c000000000200000005c053d7bb2a0d19c44045d29b796daa28c053d7b5db841fb54045d29f26a15479c053d7b1209edbf94045d2a1af11d0e3c053d7acf8868efb4045d2a4484944edc053d7abbf360b554045d2a5078be57c000000000200000002c053d7abbf360b554045d2a5078be57cc053d7aae586d7f64045d2a09cc319c6"),
			ndr: geomtest.MustHexDecode("0107000000030000000101000000550B36BFABD753C07CE58B07A5D24540010200000005000000C4190D2ABBD753C028AA6D799BD24540B51F84DBB5D753C07954A1269FD24540F9DB9E20B1D753C0E3D011AFA1D24540FB8E86F8ACD753C0ED444948A4D24540550B36BFABD753C07CE58B07A5D24540010200000002000000550B36BFABD753C07CE58B07A5D24540F6D786E5AAD753C0C619C39CA0D24540"),
		},
*/
	} {
		t.Run(fmt.Sprintf("twkb:%s", tc.twkb), func(t *testing.T) {
			test(t, tc.g, tc.precision,tc.twkb, tc.opts...)
		})
	}
	/*
	t.Run("errors when encoding empty point WKBs by default", func(t *testing.T) {
		_, err := Marshal(geom.NewPointEmpty(geom.XY), binary.LittleEndian)
		matchStr := "cannot encode empty Point in WKB"
		if err == nil || err.Error() != matchStr {
			t.Errorf("expected error matching %s, got %#v", matchStr, err)
		}
	})

	t.Run("errors when encoding empty MultiPoint WKBs by default", func(t *testing.T) {
		_, err := Marshal(geom.NewMultiPoint(geom.XY).MustSetCoords([]geom.Coord{nil, {1, 2}}), binary.LittleEndian)
		matchStr := "cannot encode empty Point in WKB"
		if err == nil || err.Error() != matchStr {
			t.Errorf("expected error matching %s, got %#v", matchStr, err)
		}
	})
*/
}
/*
func TestRandom(t *testing.T) {
	for _, tc := range testdata.Random {
		test(t, tc.G, nil, tc.WKB)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range testdata.Random {
			if _, err := Unmarshal(tc.WKB); err != nil {
				b.Errorf("unmarshal error %v", err)
			}
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tc := range testdata.Random {
			if _, err := Marshal(tc.G, NDR); err != nil {
				b.Errorf("marshal error %v", err)
			}
		}
	}
}

func TestCrashes(t *testing.T) {
	// FIXME this test modifies a global variable. It will be racy if tests are
	// run in parallel.
	savedMaxGeometryElements := wkbcommon.MaxGeometryElements
	defer func() {
		wkbcommon.MaxGeometryElements = savedMaxGeometryElements
	}()
	wkbcommon.MaxGeometryElements[1] = 1 << 20
	for _, tc := range []struct {
		s    string
		want error
	}{
		{
			s: "\x01\x03\x00\x00\x00\x04\x00\x00\x00\a\x00\x00tٽ&\xf2\xa6\xd0\x1a" +
				"\xce\xc7\x1a\xfd67\xa3\x98Y.\xa5\xfbH\x1b\xe7|\xbe\xac\xfd%" +
				";\x05\\\x90c\x83\xe9g\x01\xcbk\xa3\xc8\xdb\x0f\xae\x16bYl" +
				"\x1b\x1a\xae\xe0\x95=o\x85/\xec\xd2~\xf3\xce\xe7\xad\x04\x92\xc3\xea" +
				"r\xacE\xe3A\u008cR\x86sb\xd5sҙ\u007f\x82\xec\x88\xff" +
				"\x8aM\xa7\u007f;\x9b\x93\xa2tٽ&\xf2\xa6\xd0\x1a\xce\xc7\x1a\xfd" +
				"67\xa3\x98\x05\x00\x00\x004\xed\x19\x9c/\x8ej\ue643\x018" +
				"?\x01|\x02\xa2\xad\x18Wyʡ\xb4h\xc1j\xf6\xbb\xf0=\xbf" +
				"\x03d%\xe6PsyQ\xce4pѹ\x1dcR\xadr\x14\t" +
				"\x02pm\x86=_\xfb%\x81\"\xde\xdf4\xed\x19\x9c/\x8ej\xee" +
				"\x99\x83\x018?\x01|\x02\x05\x00\x00\x00\xfb#\xbf\xc8\xe2i\xe9'" +
				"<(\xa3\u05ccz\x06a\x8e\x17<\x956\xa4\\K\xccy\u05f7" +
				"\xcc\xdfԴp.\x9b\xce\xef0nx}\xe9\xfc\x10\xf7?\xc9\xcc" +
				"!,\xab\x15}*;\x84K\xeco\u07b6$_\xea\xfb#\xbf\xc8" +
				"\xe2i\xe9'<(\xa3\u05ccz\x06a\x04\x00\x00\x00\x8f\x8a\x9f9" +
				"\x81\x10h!N\xdcf\n\xf0-\xeaL\x02\xba\xe9\x03\xd6/G\xc2" +
				"\x1cj\r\xd8 \xbc\xd6r\x05աTS\xb3\xa5\xdc\xd8\xfb\")" +
				"\xab\x19\xf7̏\x8a\x9f9\x81\x10h!N\xdcf\n\xf0-\xeaL",
			want: wkbcommon.ErrGeometryTooLarge{Level: 1, N: 1946157063, Limit: wkbcommon.MaxGeometryElements[1]},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			_, err := Unmarshal([]byte(tc.s))
			require.Equal(t, tc.want, err)
		})
	}
}
*/
