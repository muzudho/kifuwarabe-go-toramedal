package kernel

// MutableKo - コウ
type MutableKo struct {
	// コウの番地。無ければ Cell_Pass と同値
	place Point
}

// GetPlace - コウの番地を取得
func (k *MutableKo) GetPlace() Point {
	return k.place
}

// SetPlace - コウの番地を設定
func (k *MutableKo) SetPlace(placeKo Point) {
	k.place = placeKo
}

// ClearPlace - コウの番地をクリアー
func (k *MutableKo) ClearPlace() {
	k.place = Cell_Pass
}
