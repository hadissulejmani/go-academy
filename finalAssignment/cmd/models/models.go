package models

type List struct {
	ID   int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Task struct {
	ID        int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Text      string `json:"text,omitempty"`
	ListId    int64  `json:"listId,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
