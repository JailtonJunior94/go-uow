-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES ($1, $2, $3);

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4);