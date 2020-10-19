package get

// HistoryResult type
type HistoryResult struct {
	Target string `json:"target" xml:"target" form:"target"`
	Last   int64  `json:"last" xml:"last" form:"last"`
	Text   string `json:"text" xml:"text" form:"text"`
}
