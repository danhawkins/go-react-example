package todos

type TodoService struct {
	repo *TodoRepository
}

func NewTodoService(repo *TodoRepository) *TodoService {
	svc := &TodoService{
		repo: repo,
	}
	return svc
}
