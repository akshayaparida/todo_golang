package repository

import (
    "database/sql"
    "errors"
    "time"

    "github.com/akshayaparida/todo_golang/models"
    _ "github.com/lib/pq"
)

type TodoRepository struct {
    db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (r *TodoRepository) GetTodosByUserID(userID int) ([]models.Todo, error) {
    query := `
        SELECT id, title, description, is_completed, due_date, user_id, created_at, updated_at
        FROM todos
        WHERE user_id = $1
        ORDER BY created_at DESC
    `
    
    rows, err := r.db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        var dueDate sql.NullTime
        
        err := rows.Scan(
            &todo.ID,
            &todo.Title,
            &todo.Description,
            &todo.IsCompleted,
            &dueDate,
            &todo.UserID,
            &todo.CreatedAt,
            &todo.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }

        if dueDate.Valid {
            todo.DueDate = dueDate.Time
        }

        todos = append(todos, todo)
    }

    return todos, nil
}


func (r *TodoRepository) CreateTodo(todo models.Todo) (int, error) {
    query := `
        INSERT INTO todos (title, description, is_completed, due_date, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `
    
    var id int
    err := r.db.QueryRow(
        query,
        todo.Title,
        todo.Description,
        todo.IsCompleted,
        todo.DueDate,
        todo.UserID,
        time.Now(),
        time.Now(),
    ).Scan(&id)

    if err != nil {
        return 0, err
    }

    return id, nil
}

func (r *TodoRepository) UpdateTodo(todo models.Todo) error {
    query := `
        UPDATE todos
        SET title = $1,
            description = $2,
            is_completed = $3,
            due_date = $4,
            updated_at = $5
        WHERE id = $6 AND user_id = $7
        RETURNING id
    `
    
    var id int
    err := r.db.QueryRow(
        query,
        todo.Title,
        todo.Description,
        todo.IsCompleted,
        todo.DueDate,
        time.Now(),
        todo.ID,
        todo.UserID,
    ).Scan(&id)

    if err != nil {
        if err == sql.ErrNoRows {
            return errors.New("todo not found or not owned by user")
        }
        return err
    }

    return nil
}

func (r *TodoRepository) DeleteTodo(id, userID int) error {
    query := `
        DELETE FROM todos
        WHERE id = $1 AND user_id = $2
    `
    
    result, err := r.db.Exec(query, id, userID)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("todo not found or not owned by user")
    }

    return nil
}