package controllers

import (
	"gin-mnc/initializers"
	model "gin-mnc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	// Binding request body to transaction struct
	var transaction struct {
		GiverID     uint64  `json:"giver_id"`
		ReceiverID  uint64  `json:"receiver_id"`
		Amount      float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi data
	if transaction.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	// Mulai transaksi database
	tx := initializers.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Buat objek transaksi
	newTransaction := model.Transaction{
		GiverID:    transaction.GiverID,
		ReceiverID: transaction.ReceiverID,
		Amount:     transaction.Amount,
	}

	// Simpan transaksi ke database
	if err := tx.Create(&newTransaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kurangi saldo dari giver_id
	if err := tx.Exec("UPDATE customers SET balance = balance - ? WHERE id = ?", transaction.Amount, transaction.GiverID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tambah saldo ke receiver_id
	if err := tx.Exec("UPDATE customers SET balance = balance + ? WHERE id = ?", transaction.Amount, transaction.ReceiverID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit transaksi database
	tx.Commit()

	// Kembalikan respons
	c.JSON(http.StatusOK, gin.H{"transaction": newTransaction})
}

