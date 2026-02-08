package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/DioSaputra28/belejar-go-dasar/internal/transaction/dto"
	"github.com/DioSaputra28/belejar-go-dasar/internal/transaction/model"
)

type TransactionRepository interface {
	CreateTransaction(items []dto.CheckoutItem) (model.Transaction, error)
	GetReport(startDate, endDate time.Time) (model.ReportResponse, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(items []dto.CheckoutItem) (model.Transaction, error) {
	if len(items) == 0 {
		return model.Transaction{}, errors.New("no items in checkout")
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return model.Transaction{}, errors.New("quantity must be greater than zero")
		}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return model.Transaction{}, err
	}
	defer tx.Rollback()

	var totalAmount int
	createdAt := time.Now()

	var transactionID int
	queryInsertTransaction := `INSERT INTO transactions (total_amount, created_at) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(queryInsertTransaction, 0, createdAt).Scan(&transactionID)
	if err != nil {
		return model.Transaction{}, err
	}

	for _, item := range items {
		var productID int
		var productName string
		var productPrice int
		var productStock int
		var productCreatedAt time.Time

		queryGetProduct := `SELECT id, name, price, stock, created_at FROM product WHERE id = $1 FOR UPDATE`
		err = tx.QueryRow(queryGetProduct, item.ProductID).Scan(
			&productID, &productName, &productPrice, &productStock, &productCreatedAt,
		)
		if err == sql.ErrNoRows {
			return model.Transaction{}, errors.New("product not found")
		}
		if err != nil {
			return model.Transaction{}, err
		}

		if productStock < item.Quantity {
			return model.Transaction{}, errors.New("insufficient stock")
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		queryUpdateStock := `UPDATE product SET stock = stock - $1 WHERE id = $2`
		_, err = tx.Exec(queryUpdateStock, item.Quantity, item.ProductID)
		if err != nil {
			return model.Transaction{}, err
		}

		queryInsertDetail := `INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)`
		_, err = tx.Exec(queryInsertDetail, transactionID, item.ProductID, item.Quantity, subtotal)
		if err != nil {
			return model.Transaction{}, err
		}
	}

	queryUpdateTotal := `UPDATE transactions SET total_amount = $1 WHERE id = $2`
	_, err = tx.Exec(queryUpdateTotal, totalAmount, transactionID)
	if err != nil {
		return model.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.Transaction{}, err
	}

	return model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		CreatedAt:   createdAt,
	}, nil
}

func (r *transactionRepository) GetReport(startDate, endDate time.Time) (model.ReportResponse, error) {
	var report model.ReportResponse

	err := r.db.QueryRow(
		"SELECT COALESCE(SUM(total_amount), 0) FROM transactions WHERE created_at >= $1 AND created_at < $2",
		startDate, endDate,
	).Scan(&report.TotalRevenue)
	if err != nil {
		return model.ReportResponse{}, err
	}

	err = r.db.QueryRow(
		"SELECT COUNT(id) FROM transactions WHERE created_at >= $1 AND created_at < $2",
		startDate, endDate,
	).Scan(&report.TotalTransaksi)
	if err != nil {
		return model.ReportResponse{}, err
	}

	var productName string
	var qtyTerjual int
	err = r.db.QueryRow(
		`SELECT p.name, SUM(td.quantity) as total_qty
		 FROM transaction_details td
		 JOIN transactions t ON td.transaction_id = t.id
		 JOIN product p ON td.product_id = p.id
		 WHERE t.created_at >= $1 AND t.created_at < $2
		 GROUP BY p.id, p.name
		 ORDER BY total_qty DESC
		 LIMIT 1`,
		startDate, endDate,
	).Scan(&productName, &qtyTerjual)

	if err != nil && err != sql.ErrNoRows {
		return model.ReportResponse{}, err
	}

	if err == nil {
		report.ProdukTerlaris = model.ProdukTerlaris{
			Nama:       productName,
			QtyTerjual: qtyTerjual,
		}
	}

	return report, nil
}
