-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES ($1, $2, $3);

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, price, category_id) VALUES ($1, $2, $3, $4, $5);