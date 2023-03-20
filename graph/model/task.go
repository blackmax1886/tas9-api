package model

type Task struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Content  *string    `json:"content"`
	Done     bool       `json:"done"`
	Due      *string    `json:"due"`
	Start    *string    `json:"start"`
	End      *string    `json:"end"`
	Group    *string    `json:"group"`
	Type     *string    `json:"type"`
	Priority *string    `json:"priority"`
	Archived bool       `json:"archived"`
	UserID   *string    `json:"user_id" gorm:"size:64"`
	User     *User      `json:"user" gorm:"foreignKey:UserID"`
	Subtasks []*Subtask `json:"subtasks"`
}

type Subtask struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	TaskID   *string `json:"task_id" gorm:"size:64"`
	Task     *Task   `json:"task" gorm:"foreignKey:TaskID"`
	Content  *string `json:"content"`
	Done     bool    `json:"done"`
	Due      *string `json:"due"`
	Start    *string `json:"start"`
	End      *string `json:"end"`
	Priority *string `json:"priority"`
	Archived bool    `json:"archived"`
}
