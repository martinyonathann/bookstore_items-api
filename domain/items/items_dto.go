package items

type Item struct {
	ID          int64  `json:"id"`
	BookName    string `json:"book_name"`
	Detail      string `json:"detail_book"`
	Price       string `json:"price"`
	Writer      string `json:"writer"`
	YearCreated string `json:"year_created"`
	FlagActive  string `json:"flag_active"`
}
type Items []Item
