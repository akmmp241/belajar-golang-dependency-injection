package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	Repository *BarRepository
}

func NewBarService(repository *BarRepository) *BarService {
	return &BarService{Repository: repository}
}
