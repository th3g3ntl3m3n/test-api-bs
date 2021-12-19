package api

import (
	"net/http"
	"strconv"
	"test-api-bs/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Apis interface {
	SaveEntry(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserPortfolio(c *gin.Context)
}

type api struct {
	db *gorm.DB
}

func New(db *gorm.DB) *gin.Engine {

	newApi := &api{db}

	router := gin.New()

	router.POST("/portfolio/:id/entry", newApi.SaveEntry)
	router.GET("/user/:id/portfolio", newApi.GetUserPortfolio)
	router.GET("/users", newApi.GetUsers)
	router.GET("/users/:id", newApi.GetUserByID)

	return router
}

func (api *api) SaveEntry(c *gin.Context) {
	pid := c.Params.ByName("id")
	portfolioID, err := strconv.Atoi(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	var entry model.Entry
	if err = c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	var portfolio model.Portfolio
	api.db.Raw("SELECT id,name FROM portfolios WHERE id = ?", portfolioID).Scan(&portfolio)
	if portfolio.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "portfolio not found"})
		return
	}

	if err := api.db.Exec("INSERT INTO entries (folio_id, coin_name, amount, price, transaction_fee) values (?,?,?,?,?)", portfolio.Id, entry.CoinName, entry.Amount, entry.Price, entry.TransactionFee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "done"})
}

func (api *api) GetUsers(c *gin.Context) {
	var users []model.User
	api.db.Raw("SELECT id,name,email,verified, locked FROM users").Scan(&users)
	c.JSON(http.StatusOK, users)
}

func (api *api) GetUserByID(c *gin.Context) {
	userid := c.Params.ByName("id")
	uid, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	var user model.User
	api.db.Raw("SELECT id,name,email,verified, locked FROM users WHERE id = ?", uid).Scan(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, user)

}

func (api *api) GetUserPortfolio(c *gin.Context) {
	userid := c.Params.ByName("id")
	uid, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	var user model.User
	api.db.Raw("SELECT id,name,email,verified, locked FROM users WHERE id = ?", uid).Scan(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var portfolio model.Portfolio
	api.db.Raw("SELECT portfolios.id,portfolios.name FROM portfolios LEFT JOIN users ON users.id = portfolios.user_id WHERE portfolios.user_id = ?", user.Id).Scan(&portfolio)

	if portfolio.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "portfolio not found"})
		return
	}

	var entries []model.Entry
	api.db.Raw("SELECT coin_name, amount, price, transaction_fee FROM entries WHERE folio_id = ?", portfolio.Id).Scan(&entries)
	if len(entries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "entries not found"})
		return
	}

	c.JSON(http.StatusOK, entries)
}
