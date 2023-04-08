package entity

type Category struct {
	ID       int
	Name     string
	CourseID []string
}

func (c *Category) AddCourse(id string) {
	c.CourseID = append(c.CourseID, id)
}

type Course struct {
	ID         int
	Name       string
	CategoryID string
}
