package tasks_domain

type Task struct {
	ID          string `bson:"_id,omitempty" json:"id" validate:""`
	Name        string `bson:"name" json:"name" validate:"required"`
	Description string `bson:"description" json:"description" validate:"required"`
	Status      string `bson:"status" json:"status" validate:"required"`
	DueDate     string `bson:"dueDate" json:"dueDate" validate:"required"`
	Priority    string `bson:"priority" json:"priority" validate:"required"`
}
