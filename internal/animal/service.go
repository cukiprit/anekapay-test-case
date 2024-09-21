package animal

type Repository interface {
	Create(animal *Animal) error
	Update(animal *Animal) error
	Delete(id int) error
	GetAll() ([]Animal, error)
	GetByID(id int) (*Animal, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(animal *Animal) error {
	return s.repo.Create(animal)
}

func (s *Service) Update(animal *Animal) error {
	return s.repo.Update(animal)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) GetAll() ([]Animal, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id int) (*Animal, error) {
	return s.repo.GetByID(id)
}
