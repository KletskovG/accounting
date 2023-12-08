package common

type Transaction struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Date      string `bson:"date,omitempty" json:"date"`
	Expense   int32  `bson:"expense,omitempty" json:"expense"`
	Category  string `bson:"category,omitempty" json:"category"`
	Note      string `bson:"note,omitempty" json:"note"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
}
