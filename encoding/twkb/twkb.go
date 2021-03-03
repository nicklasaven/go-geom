// Package wkb implements Tiny Well Known Binary encoding and decoding.
//

package twkb

import (
	"fmt"
	"encoding/binary"
//	"fmt"
	"io"
	"bytes"
	"math"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkbcommon"
)

const (
	wkbXYID   = 0
	wkbXYZID  = 1000
	wkbXYMID  = 2000
	wkbXYZMID = 3000
)
type twkb struct{
	nDims int
	precision byte
	lastCoords []int64
	geomType byte
	hasBbox bool
	hasSize bool
	hasId bool
	isEmpty bool
	hasZ bool
	hasM bool
	zPrecision byte
	mPrecision byte
	cursor *uint
}

// Read reads an arbitrary geometry from r.
func Read(r io.ByteReader, opts ...wkbcommon.WKBOption) (geom.T, error) {


	var t twkb 

	t.readTWKBheader(r)
	t.lastCoords = make([]int64, t.nDims)
	layout := geom.NoLayout

	if !t.hasZ && !t.hasM && t.nDims == 2{		
		layout = geom.XY
	}else if t.hasZ && t.hasM && t.nDims == 4{
		layout = geom.XYZM
	}else if t.hasZ && t.nDims == 3{
		layout = geom.XYZ		
	}else if t.hasM && t.nDims == 3{
		layout = geom.XYM		
	}else {
		return nil, wkbcommon.ErrUnknownType(t.geomType)
	}


	return t.ReadWoHeader(r ,layout, opts ...)
}

func (t twkb) ReadWoHeader(r io.ByteReader,layout geom.Layout ,  opts ...wkbcommon.WKBOption) (geom.T, error) {
	params := wkbcommon.InitWKBParams(
	wkbcommon.WKBParams{
		EmptyPointHandling: wkbcommon.EmptyPointHandlingError,
	},
	opts...,
	)
	switch int(t.geomType) {
	case wkbcommon.PointID:
		flatCoords, err := t.ReadFlatCoords0(r)
		if err != nil {
			return nil, err
		}
		
		
		if params.EmptyPointHandling == wkbcommon.EmptyPointHandlingNaN {
			return geom.NewPointFlatMaybeEmpty(layout, flatCoords), nil
		}
		return geom.NewPointFlat(layout, flatCoords), nil
	case wkbcommon.LineStringID:
		flatCoords, err := t.ReadFlatCoords1(r)
		if err != nil {
			return nil, err
		}
		return geom.NewLineStringFlat(layout, flatCoords), nil
	case wkbcommon.PolygonID:
		flatCoords, ends, err := t.ReadFlatCoords2(r)
		if err != nil {
			return nil, err
		}
		return geom.NewPolygonFlat(layout, flatCoords, ends), nil
	case wkbcommon.MultiPointID:
		n, err := readUvarint(r)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[1]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 1, N: int(n), Limit: limit}
		}
		mp := geom.NewMultiPoint(layout)
		for i := uint64(0); i < n; i++ {
			g, err := t.ReadWoHeader(r,layout, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*geom.Point)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &geom.Point{}}
			}
			if err = mp.Push(p); err != nil {
				return nil, err
			}
		}
		return mp, nil
	case wkbcommon.MultiLineStringID:
		n, err := readUvarint(r)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[2]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 2, N: int(n), Limit: limit}
		}
		mls := geom.NewMultiLineString(layout)
		for i := uint64(0); i < n; i++ {
			g, err := t.ReadWoHeader(r,layout, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*geom.LineString)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &geom.LineString{}}
			}
			if err = mls.Push(p); err != nil {
				return nil, err
			}
		}
		return mls, nil
	case wkbcommon.MultiPolygonID:
		n, err := readUvarint(r)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[3]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 3, N: int(n), Limit: limit}
		}
		mp := geom.NewMultiPolygon(layout)
		for i := uint64(0); i < n; i++ {
			g, err := t.ReadWoHeader(r,layout, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*geom.Polygon)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &geom.Polygon{}}
			}
			if err = mp.Push(p); err != nil {
				return nil, err
			}
		}
		return mp, nil
	case wkbcommon.GeometryCollectionID:
		n, err := readUvarint(r)
		if err != nil {
			return nil, err
		}
		gc := geom.NewGeometryCollection()
		for i := uint64(0); i < n; i++ {
			g, err := Read(r, opts...)
			if err != nil {
				return nil, err
			}
			if err := gc.Push(g); err != nil {
				return nil, err
			}
		}
		return gc, nil
	default:
		return nil, wkbcommon.ErrUnsupportedType(uint32(t.geomType))
	}
}

func (t *twkb)readTWKBheader(r io.ByteReader)(error){
	

	byte1, err := r.ReadByte()
	if err != nil {
		return  err
	}
	
	t.geomType = byte(byte1 & 15)

	t.precision = byte1>>4
	extended := false
	byte2, err := r.ReadByte()
	if err != nil {
		return nil
	}
	
	if byte2 & 1 > 0{
		t.hasBbox = true
	}
	if byte2 & 2 > 0{
		t.hasSize = true
	}
	if byte2 & 4 > 0{
		t.hasId = true
	}
	if byte2 & 8 > 0{
		extended = true
	}
	if byte2 & 16 > 0{
		t.isEmpty = true
	}
	
	t.nDims= 2
	if(extended){
		byte3, err := r.ReadByte()
		if err != nil {
			return nil
		}
		if byte3 & 1 > 0{
			t.hasZ = true
			t.nDims++
		}
		if byte3 & 2 > 0{
			t.hasM = true
			t.nDims++
		}
		if t.hasZ{
			t.zPrecision = (byte3 & 28) >> 2
		}
		if t.hasM{
			t.mPrecision = (byte3 & 224) >> 5
		}
	}	
	return nil
}

type resbuf struct{
data *[]byte
}
// Unmarshal unmrshals an arbitrary geometry from a []byte.
func Unmarshal(data []byte, opts ...wkbcommon.WKBOption) (geom.T, error) {
	return Read(bytes.NewBuffer(data), opts...)
}

// Write writes an arbitrary geometry to w.
func Write(buf resbuf,precisionXY byte,precisionZ byte, precisionM byte, g geom.T) error {



	if g.Empty(){
		buf.resize(2)
		(*buf.data)[0] = 0
		(*buf.data)[1] = 16
		return nil
	}

var err error
/*
	params := wkbcommon.InitWKBParams(
		wkbcommon.WKBParams{
			EmptyPointHandling: wkbcommon.EmptyPointHandlingError,
		},
		opts...,
	)
	*/

	var t twkb

	switch g.(type) {
	case *geom.Point:
		t.geomType = byte(wkbcommon.PointID)
	case *geom.LineString:
		t.geomType = byte(wkbcommon.LineStringID)
	case *geom.Polygon:
		t.geomType = byte(wkbcommon.PolygonID)
	case *geom.MultiPoint:
		t.geomType = byte(wkbcommon.MultiPointID)
	case *geom.MultiLineString:
		t.geomType = byte(wkbcommon.MultiLineStringID)
	case *geom.MultiPolygon:
		t.geomType = byte(wkbcommon.MultiPolygonID)
	case *geom.GeometryCollection:
		t.geomType = byte(wkbcommon.GeometryCollectionID)
	default:
		return geom.ErrUnsupportedType{Value: g}
	}


	t.nDims = g.Stride()
	switch g.Layout() {
	case geom.NoLayout:
		// Special case for empty GeometryCollections
		if _, ok := g.(*geom.GeometryCollection); !ok || !g.Empty() {
			return geom.ErrUnsupportedLayout(g.Layout())
		}
	case geom.XY:
		t.hasZ = false
		t.hasM = false
		t.nDims = 2
		fmt.Printf("2D\n")
	case geom.XYZ:
		t.hasZ = true
		t.zPrecision = precisionZ
		t.hasM = false
		t.nDims = 3
		fmt.Printf("Has Z\n")
	case geom.XYM:
		t.hasZ = false
		t.hasM = true
		t.mPrecision = precisionM
		t.nDims = 3
		fmt.Printf("Has M\n")
	case geom.XYZM:
		t.hasZ = true
		t.zPrecision = precisionZ
		t.hasM = true
		t.mPrecision = precisionM
		t.nDims = 4

		fmt.Printf("Has Z and M\n")
	default:
		return geom.ErrUnsupportedLayout(g.Layout())
	}

	t.lastCoords = make([]int64, t.nDims)
	err = t.writeHeader(buf)
	if err != nil{
		return err
	}

	t.writeWoHeader(buf, g)
	return nil
}

func (t twkb) writeWoHeader(buf resbuf,g geom.T)(error){

	switch g := g.(type) {
	case *geom.Point:
		return t.writeFlatCoords0(buf, g.FlatCoords())
	case *geom.LineString:
		return t.writeFlatCoords1(buf, g.FlatCoords())
	case *geom.Polygon:
		return t.writeFlatCoords2(buf, g.FlatCoords(), g.Ends())
	case *geom.MultiPoint:
		n := g.NumPoints()
		buf.appendUvarint(uint64(n))
		for i := 0; i < n; i++ {
			if err := t.writeWoHeader(buf,g.Point(i)); err != nil {
				return err
			}
		}
		return nil
	case *geom.MultiLineString:
		n := g.NumLineStrings()
		buf.appendUvarint(uint64(n))
		for i := 0; i < n; i++ {
			if err := t.writeWoHeader(buf, g.LineString(i)); err != nil {
				return err
			}
		}
		return nil
	case *geom.MultiPolygon:
		n := g.NumPolygons()		
		buf.appendUvarint(uint64(n))
		for i := 0; i < n; i++ {
			if err := t.writeWoHeader(buf, g.Polygon(i)); err != nil {
				return err
			}
		}
		return nil
	case *geom.GeometryCollection:
		n := g.NumGeoms()		
		buf.appendUvarint(uint64(n))
		for i := 0; i < n; i++ {
			if err := Write(buf,t.precision, t.zPrecision, t.mPrecision, g.Geom(i)); err != nil {
				return err
			}
		}
		return nil
	default:
		return geom.ErrUnsupportedType{Value: g}
	}
}

// Marshal marshals an arbitrary geometry to a []byte.
func Marshal(g geom.T,precisionXY byte, precisionZ byte, precisionM byte, opts ...wkbcommon.WKBOption) ([]byte, error) {
	var res []byte
	
	if err := Write(resbuf{data:&res}, precisionXY, precisionZ, precisionM, g); err != nil {
		return nil, err
	}
	return []byte(res), nil
}


func (t twkb) writeHeader(buf resbuf)(error){
	
	fmt.Printf("1: buf cap: %d, buf len: %d\n", cap(*buf.data), len(*buf.data))
	buf.resize(2)
	var firstByte byte 
	firstByte &= byte(t.precision)
	firstByte = firstByte<< 4	
	firstByte |= byte(t.geomType)

	(*buf.data)[0] = firstByte

	/*TODO add support for bbox, size id-list and empty*/
	var secondByte byte = 0


	var thirdByte byte = 0
	
	if t.hasZ || t.hasM{
		fmt.Printf("t.hasZ: %v, secondByte: %d, thirdByte: %d\n", t.hasZ,secondByte, thirdByte)
		secondByte |= 8 // we set the fourth bit indicating extended precision information byte
		var zbyte, mbyte byte
		if t.hasZ{
			zbyte = t.zPrecision << 2
			zbyte |= 1
		}
		if t.hasM{
			mbyte = t.mPrecision << 5
			mbyte |= 2			
		}
		thirdByte = zbyte | mbyte

		fmt.Printf("t.hasZ: %v, secondByte: %d,mbyte: %d, zbyte: %d, thirdByte: %d\n", t.hasZ,secondByte,mbyte, zbyte, thirdByte)
	}

	
	(*buf.data)[1] = secondByte
	if thirdByte > 0{
		buf.resize(1)
		fmt.Printf("Writing third byte")
		(*buf.data)[2] = thirdByte	
	}
	
	fmt.Printf("2: buf cap: %d, buf len: %d\n", cap(*buf.data), len(*buf.data))
	return nil	
}



// ReadFlatCoords0 reads flat coordinates 0.
func (t twkb) ReadFlatCoords0(r io.ByteReader) ([]float64, error) {
	coord, err := t.readFloatArray(r, 1)
	if err != nil {
		return nil, err
	}
	return coord, nil
}

// ReadFlatCoords1 reads flat coordinates 1.
func  (t twkb) ReadFlatCoords1(r io.ByteReader) ([]float64, error) {
	n, err := readUvarint(r)
	if err != nil {
		return nil, err
	}
	if limit := wkbcommon.MaxGeometryElements[1]; limit >= 0 && int(n) > limit {
		return nil, wkbcommon.ErrGeometryTooLarge{Level: 1, N: int(n), Limit: limit}
	}

	flatCoords, err := t.readFloatArray(r, n)
	if err != nil{
		return nil, err
	}
	return flatCoords, nil
}

// ReadFlatCoords2 reads flat coordinates 2.
func (t twkb) ReadFlatCoords2(r io.ByteReader) ([]float64, []int, error) {
	n, err := readUvarint(r)
	if err != nil {
		return nil, nil, err
	}
	if limit := wkbcommon.MaxGeometryElements[2]; limit >= 0 && int(n) > limit {
		return nil, nil, wkbcommon.ErrGeometryTooLarge{Level: 2, N: int(n), Limit: limit}
	}
	var flatCoordss []float64
	var ends []int
	for i := 0; i < int(n); i++ {
		flatCoords, err := t.ReadFlatCoords1(r)
		if err != nil {
			return nil, nil, err
		}
		flatCoordss = append(flatCoordss, flatCoords...)
		ends = append(ends, len(flatCoordss))
	}
	return flatCoordss, ends, nil
}

func (t twkb) readFloatArray(r io.ByteReader, n uint64)([]float64, error){
	
	flatCoords := make([]float64, int(n)*t.nDims)
	for i:=0; i<int(n); i++{	
		for j:=0; j<t.nDims; j++{
			v, err := binary.ReadVarint(r)
			if err != nil{return nil, err}
			t.lastCoords[j] += v
			flatCoords[i*t.nDims + j] = float64(t.lastCoords[j])/math.Pow(10.0,float64(t.precision))
		}
	}
	return flatCoords, nil
}



func (t twkb) readVarint2float64(r io.ByteReader)(float64, error){
	intVal, err := binary.ReadVarint(r)
	if err != nil{
		return 0, err 
	}
	return float64(intVal)/math.Pow(10.0,float64(t.precision))	, nil
}




func readUvarint(r io.ByteReader)(uint64, error){
	
	uIntVal, err := binary.ReadUvarint(r)	
	if err != nil{
		return 0, err 
	}
	return uIntVal, nil
}

func readVarint(r io.ByteReader)(int64, error){
	
	intVal, err := binary.ReadVarint(r)	
	if err != nil{
		return 0, err 
	}
	return intVal, nil
}






// WriteFlatCoords0 writes flat coordinates 0.
func (t twkb) writeFlatCoords0(buf resbuf, coords []float64) error {
		t.writeFloatArray(buf, coords, false)
	return nil
}

// WriteFlatCoords1 writes flat coordinates 1.
func (t twkb) writeFlatCoords1(buf resbuf, coords []float64) error {	
	t.writeFloatArray(buf, coords, true)
	return nil}

// WriteFlatCoords2 writes flat coordinates 2.
func (t twkb) writeFlatCoords2(buf resbuf,  flatCoords []float64, ends []int) error {
	buf.appendUvarint(uint64(len(ends)))
	offset := 0
	for _, end := range ends {
		if err := t.writeFlatCoords1(buf, flatCoords[offset:end]); err != nil {
			return err
		}
		offset = end
	}
	return nil
}





func (t twkb)writeFloatArray(buf resbuf, coords []float64, writeNumPoints bool){

	nPoints := uint64(len(coords)/t.nDims)

	if writeNumPoints{
		buf.appendUvarint(uint64(nPoints))
	}

	nDims := t.nDims
	for i:=uint64(0);i<nPoints;i++{
		for j:=0;j<nDims;j++{
				newVal := int64(math.Round(coords[int(i)*nDims + j] * (float64(math.Pow(10.0,float64(t.precision))))))
				nextDelta := newVal - t.lastCoords[j]				
				buf.appendVarint(nextDelta)
				t.lastCoords[j] += nextDelta			
		}
	}
}


func (s resbuf)appendVarint(val int64){	
	buf := make([]byte, binary.MaxVarintLen64)	
	n := binary.PutVarint(buf, val)	
	
	m := s.resize(n)
	copy((*s.data)[m:m+n], buf)
//	*dst = append(*dst, buf[0:n] ...)
}

func (s resbuf)appendUvarint(val uint64){	
	buf := make([]byte, binary.MaxVarintLen64)	
	n := binary.PutUvarint(buf, val)	
	m := s.resize(n)	
	copy((*s.data)[m:m+n], buf)
	
	
//	*dst = append(*dst, buf[0:n] ...)
}
/*
func appendUVarint(dst *[]byte, val uint64){	
	buf := make([]byte, binary.MaxVarintLen64)	
	n := binary.PutUvarint(buf, val)	
	*dst = append(*dst, buf[0:n] ...)
}
*/

func (s resbuf)resize(n int)(int){
	fmt.Printf("resize1: buf cap: %d, buf len: %d\n", cap(*s.data), len(*s.data))
	var res []byte
	m := len(*s.data)
	if cap(*s.data) < m + n {
		var newSize int
		if cap(*s.data) == 0{
			newSize = 16
		}else{
			newSize = m
		}
		for newSize < m+n{
			newSize *= 2
			fmt.Printf("newSize: %d, m+n: %d\n", newSize, m+n)
		}

		res = make([]byte,m+n, newSize)
		copy(res, (*s.data)[:m])
		*s.data = res
		return m
	}
	(*s.data) = (*s.data)[0:m+n]
	
	fmt.Printf("resize2: buf cap: %d, buf len: %d\n", cap(*s.data), len(*s.data))
	return m
}