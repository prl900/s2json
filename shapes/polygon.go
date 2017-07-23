package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Polygon struct {
	s2.Polygon
}

type geoJSONPolygon struct {
	Type   string        `json:"type"`
	BBox   []float64     `json:"bbox,omitempty"`
	Coords [][][]float64 `json:"coordinates"`
}

func (p *S2Polygon) Coverage(maxLevel, maxCells int) []s2.CellID {
	rc := &s2.RegionCoverer{MaxLevel: maxLevel, MaxCells: maxCells}
	return rc.InteriorCellUnion(p.Polygon.Loops()[0])
}

func (p *S2Polygon) MarshalJSON() ([]byte, error) {
	plnglats := [][][]float64{}
	lo := p.Polygon.RectBound().Lo().Normalized()
	hi := p.Polygon.RectBound().Hi().Normalized()
	for _, loop := range p.Polygon.Loops() {
		lnglats := [][]float64{}
		for _, point := range loop.Vertices() {
			ll := s2.LatLngFromPoint(point)
			lnglats = append(lnglats, []float64{ll.Lng.Degrees(), ll.Lat.Degrees()})
		}
		plnglats = append(plnglats, lnglats)
	}

	gP := &geoJSONPolygon{"Polygon", []float64{lo.Lng.Degrees(), lo.Lat.Degrees(), hi.Lng.Degrees(), hi.Lat.Degrees()}, plnglats}
	return json.Marshal(gP)
}

func (p *S2Polygon) UnmarshalJSON(in []byte) error {
	pView := &geoJSONPolygon{}

	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	loops := []*s2.Loop{}
	for _, lll := range pView.Coords {
		points := []s2.Point{}
		for _, ll := range lll {
			p := s2.PointFromLatLng(s2.LatLngFromDegrees(ll[1], ll[0]))
			points = append(points, p)
		}
		loop := s2.LoopFromPoints(points)
		loops = append(loops, loop)
	}

	*p = S2Polygon{*s2.PolygonFromLoops(loops)}

	return err
}
