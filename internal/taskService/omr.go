package taskService

type TaskSt struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
}

type RequestBody struct {
	Task string `json:"task"`
}
