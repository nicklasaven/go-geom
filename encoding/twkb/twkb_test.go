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
				got, err := Unmarshal(twkb)
				require.NoError(t, err)
				require.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, precision[0], precision[1], precision[2])
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

		{
			g:    geom.NewPointEmpty(geom.XY),
			precision: []byte{0,0,0},
			twkb:  geomtest.MustHexDecode("0110"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYM),
			precision: []byte{0,0,0},
			twkb:  geomtest.MustHexDecode("011802"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYZ),
			precision: []byte{0,0,0},
			twkb:  geomtest.MustHexDecode("011801"),
		},
		{
			g:    geom.NewPointEmpty(geom.XYZM),
			precision: []byte{0,0,0},
			twkb:  geomtest.MustHexDecode("011803"),
		},
/*		{
			g:    geom.NewGeometryCollection().MustPush(geom.NewPointEmpty(geom.XY)),
			precision: []byte{0,0,0},
			twkb:  geomtest.MustHexDecode("011802"),
		},
*/
		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{1, 2}),
			precision: []byte{0,0,0},
			twkb: geomtest.MustHexDecode("01000204"),
		},

		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{1, 2}),
			precision: []byte{1,0,0},
			twkb: geomtest.MustHexDecode("21001428"),
		},
		{
			g:   geom.NewPoint(geom.XYZ).MustSetCoords(geom.Coord{1.2, 2.3, 3.45}),
			precision: []byte{1,2,0},
			twkb: geomtest.MustHexDecode("210809182eb205"),
		},
		{
			g:   geom.NewPoint(geom.XYM).MustSetCoords(geom.Coord{1.2, 2.3, 3.456}),
			precision: []byte{1,0,3},
			twkb: geomtest.MustHexDecode("210862182e8036"),
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
		{
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

/*		{
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
			precision: []byte{7,0,0},
			twkb: geomtest.MustHexDecode("e70003e1009f84f7f405cab29ea003e20005a397f8f40582859ea003ce32c4118e2d8a0cd227b20cd60bc803e200029f84f7f405cab29ea0038e088715"),
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
		for _, tc := range testdata.Random {/*
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
