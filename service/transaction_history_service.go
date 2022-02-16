package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/repository"
)

type TransactionHistoryService interface {
	CreateTransaction(transactionInput input.InputTransaction, IDUser int) (entity.TransactionHistory, error)
	GetMyTransaction(IDUser int) ([]response.TransactionHistoryResponse, error)
	GetUserTransaction(IDUser int) ([]response.UserTransactionHistoryResponse, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
	productRepository            repository.ProductRepository
	userRepository               repository.UserRepository
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository, productRepository repository.ProductRepository, userRepository repository.UserRepository) *transactionHistoryService {
	return &transactionHistoryService{transactionHistoryRepository, productRepository, userRepository}
}

func (s *transactionHistoryService) CreateTransaction(transactionInput input.InputTransaction, IDUser int) (entity.TransactionHistory, error) {
	newTransactionHistory := entity.TransactionHistory{}

	newTransactionHistory.ProductID = transactionInput.ProductID
	newTransactionHistory.Quantity = transactionInput.Quantity

	// query product
	product, err := s.productRepository.GetProductByID(transactionInput.ProductID)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if product.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	// ketika jumlah tidak mencukupi
	if product.Stock < transactionInput.Quantity {
		return entity.TransactionHistory{}, errors.New("Jumlah tidak mencukupi")
	}

	// query user
	datauser, err := s.userRepository.GetByID(IDUser)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	if datauser.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	// saldo tidak mencukupi
	if datauser.Balance < (product.Price * transactionInput.Quantity) {
		return entity.TransactionHistory{}, errors.New("Saldo tidak mencukupi")
	}

	// pastikan balance tersedia
	buyAmount := product.Price * transactionInput.Quantity
	newTransactionHistory.TotalPrice = buyAmount
	newTransactionHistory.UserID = IDUser

	// kurangi stock
	productUpdate := entity.Product{
		ID:    transactionInput.ProductID,
		Stock: product.Stock - transactionInput.Quantity,
	}

	_, err = s.productRepository.Update(productUpdate)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	// store data ke transactions history
	transactionCreated, err := s.transactionHistoryRepository.Save(newTransactionHistory)

	if err != nil {
		return entity.TransactionHistory{}, err
	}

	_, err = s.userRepository.UpdateSaldo(IDUser, buyAmount)

	if transactionCreated.ID == 0 {
		return entity.TransactionHistory{}, err
	}

	return transactionCreated, nil
}

func (s *transactionHistoryService) GetMyTransaction(IDUser int) ([]response.TransactionHistoryResponse, error) {
	myTransaction, err := s.transactionHistoryRepository.GetTransactionByIDUser(IDUser)

	if err != nil {
		return []response.TransactionHistoryResponse{}, err
	}

	if len(myTransaction) < 1 {
		return []response.TransactionHistoryResponse{}, err
	}

	var myTransactionResponse []response.TransactionHistoryResponse

	for _, item := range myTransaction {
		productTemp, _ := s.productRepository.GetProductByID(item.ProductID)

		temp := response.TransactionHistoryResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			UserID:     item.UserID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product:    productTemp,
		}
		myTransactionResponse = append(myTransactionResponse, temp)
	}

	return myTransactionResponse, nil
}

func (s *transactionHistoryService) GetUserTransaction(IDUser int) ([]response.UserTransactionHistoryResponse, error) {

	userdata, err := s.userRepository.GetByID(IDUser)
	if userdata.ID == 0 {
		return []response.UserTransactionHistoryResponse{}, errors.New("users not found!")
	}

	if userdata.Role != "admin" {
		return []response.UserTransactionHistoryResponse{}, errors.New("Unauthorized user!")
	}

	myTransaction, err := s.transactionHistoryRepository.GetAllTransaction()

	if err != nil {
		return []response.UserTransactionHistoryResponse{}, err
	}

	if len(myTransaction) < 1 {
		return []response.UserTransactionHistoryResponse{}, err
	}

	var myTransactionResponse []response.UserTransactionHistoryResponse

	for _, item := range myTransaction {
		productTemp, _ := s.productRepository.GetProductByID(item.ProductID)
		userTemp, _ := s.userRepository.GetByID(item.UserID)

		temp := response.UserTransactionHistoryResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			UserID:     item.UserID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product:    productTemp,
			User:       userTemp,
		}
		myTransactionResponse = append(myTransactionResponse, temp)
	}

	return myTransactionResponse, nil
}
