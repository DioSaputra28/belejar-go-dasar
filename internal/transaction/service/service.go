package service

import (
	"errors"
	"time"

	"github.com/DioSaputra28/belejar-go-dasar/internal/transaction/dto"
	"github.com/DioSaputra28/belejar-go-dasar/internal/transaction/model"
	"github.com/DioSaputra28/belejar-go-dasar/internal/transaction/repository"
)

type TransactionService interface {
	Checkout(items []dto.CheckoutItem) (model.Transaction, error)
	GetReportToday() (model.ReportResponse, error)
	GetReportByDateRange(startDate, endDate string) (model.ReportResponse, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) Checkout(items []dto.CheckoutItem) (model.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *transactionService) GetReportToday() (model.ReportResponse, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	nextDay := startOfDay.Add(24 * time.Hour)
	return s.repo.GetReport(startOfDay, nextDay)
}

func (s *transactionService) GetReportByDateRange(startDate, endDate string) (model.ReportResponse, error) {
	const layout = "2006-01-02"

	parsedStart, err := time.ParseInLocation(layout, startDate, time.Local)
	if err != nil {
		return model.ReportResponse{}, errors.New("invalid start_date format, use YYYY-MM-DD")
	}

	parsedEnd, err := time.ParseInLocation(layout, endDate, time.Local)
	if err != nil {
		return model.ReportResponse{}, errors.New("invalid end_date format, use YYYY-MM-DD")
	}

	if parsedEnd.Before(parsedStart) || parsedEnd.Equal(parsedStart) {
		return model.ReportResponse{}, errors.New("end_date must be after start_date")
	}

	normalizedStart := time.Date(parsedStart.Year(), parsedStart.Month(), parsedStart.Day(), 0, 0, 0, 0, parsedStart.Location())
	normalizedEnd := time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 0, 0, 0, 0, parsedEnd.Location()).Add(24 * time.Hour)

	return s.repo.GetReport(normalizedStart, normalizedEnd)
}
