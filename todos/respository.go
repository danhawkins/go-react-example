package todos

import "gorm.io/gorm"

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Migrate() error {
	return r.db.AutoMigrate(&Todo{})
}
