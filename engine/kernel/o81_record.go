package kernel

// Record - 棋譜
type Record struct {
	// Items - 棋譜項目の配列
	Items []RecordItem
}

// NewRecord - 新規作成
func NewRecord() *Record {
	var r = new(Record)
	return r
}

// GetItems - 棋譜項目の配列
func (r *Record) GetItems() []RecordItem {
	return r.Items
}
