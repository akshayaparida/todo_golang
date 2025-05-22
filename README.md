# ToDo List SQL Project

A hands-on project to master essential SQL skills by designing and querying a ToDo list database. This project is perfect for beginners and intermediates who want to solidify their understanding of relational databases, data modeling, and SQL queries.

---

##  Project Overview
This project guides you through designing a robust database schema for a ToDo list application. You'll learn how to model data, define relationships, enforce data integrity, and write powerful queries to manage and analyze your tasks.

---

##  Database Schema Design

### Main Table: `todos`
```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each todo
    title VARCHAR(255) NOT NULL,            -- Task title (required)
    description TEXT,                       -- Optional detailed description
    is_completed BOOLEAN DEFAULT FALSE,     -- Completion status
    due_date TIMESTAMP,                     -- Optional due date
    user_id INTEGER NOT NULL,               -- Reference to the user (assumes a users table)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Last update timestamp
);
```

### Key Concepts Illustrated
- **Primary Key**: Ensures each todo is uniquely identifiable.
- **Foreign Key**: `user_id` links to a user (not shown, but recommended for multi-user systems).
- **Constraints**: `NOT NULL`, `DEFAULT`, and appropriate data types for each field.
- **Timestamps**: Track when todos are created and updated.

---

## Example SQL Queries

### Insert a New Todo
```sql
INSERT INTO todos (title, description, user_id, due_date)
VALUES ('Learn SQL', 'Master the basics of SQL', 1, '2024-07-01');
```

### Fetch Todos for a User
```sql
SELECT * FROM todos WHERE user_id = 1 ORDER BY created_at DESC;
```

### Update a Todo
```sql
UPDATE todos SET is_completed = TRUE, updated_at = NOW() WHERE id = 1;
```

### Delete a Todo
```sql
DELETE FROM todos WHERE id = 1;
```

### Count Completed Todos
```sql
SELECT COUNT(*) FROM todos WHERE is_completed = TRUE;
```

### Aggregation & Filtering
```sql
-- Find overdue todos
SELECT * FROM todos WHERE due_date < NOW() AND is_completed = FALSE;

-- Group todos by completion status
SELECT is_completed, COUNT(*) FROM todos GROUP BY is_completed;
```

---

##  What You Will Learn
- How to design normalized, efficient database tables
- The importance of keys and constraints for data integrity
- Writing CRUD (Create, Read, Update, Delete) operations in SQL
- Using aggregation and filtering to analyze your data
- Best practices for schema design and query writing
