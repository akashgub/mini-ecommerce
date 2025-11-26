package product

type ProductService interface {
	CreateProduct(req CreateProductRequest) (*Product, error)
	GetAllProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
	UpdateProduct(id int, req UpdateProductRequest) (*Product, error)
	DeleteProduct(id int) error
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(req CreateProductRequest) (*Product, error) {
	product := &Product{
		Name:        req.Name,
		Price:       req.Price,
		Weight:      req.Weight,
		Colour:      req.Colour,
		Description: req.Description,
	}
	err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetAllProducts() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id int) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(id int, req UpdateProductRequest) (*Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Weight > 0 {
		product.Weight = req.Weight
	}
	if req.Colour != "" {
		product.Colour = req.Colour
	}
	if req.Description != "" {
		product.Description = req.Description
	}

	err = s.repo.Update(id, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
