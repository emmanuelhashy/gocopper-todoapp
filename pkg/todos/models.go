package todos

type Todo struct {
	Name string
	Rank int64 `gorm:"-"`
}
