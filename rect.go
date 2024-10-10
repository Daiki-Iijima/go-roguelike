package main

// Rect 左上のXYと右下のXYの座標を持つ
type Rect struct {
	X1 int // 左上 X座標
	X2 int // 右下 X座標
	Y1 int // 左上 Y座標
	Y2 int // 右下 Y座標
}

func NewRect(x int, y int, width int, height int) Rect {
	return Rect{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

// Center 四角形の中心
func (r *Rect) Center() (int, int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2
	return centerX, centerY
}

// Intersect 四角形同士が重なっているか
func (r *Rect) Intersect(other Rect) bool {
	return r.X1 <= other.X2 &&
		r.X2 >= other.X1 &&
		r.Y1 <= other.Y1 &&
		r.Y2 >= other.Y1
}
