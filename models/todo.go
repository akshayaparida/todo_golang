
package models

import "time"

type Todo struct {
    ID          int       `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description,omitempty" db:"description"` // omitempty means don't include in JSON if empty
    IsCompleted bool      `json:"is_completed" db:"is_completed"`
    DueDate     time.Time `json:"due_date,omitempty" db:"due_date"`
    UserID      int       `json:"user_id" db:"user_id"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
    
    User        *User     `json:"user,omitempty" db:"-"`
}