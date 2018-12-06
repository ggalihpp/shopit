package item

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	mw "github.com/ggalihpp/shopit/middleware"
	"github.com/labstack/echo"
)

// Handler is
type Handler struct {
	Repository *Repository
}

// SetRoutes is
func (h *Handler) SetRoutes(r *echo.Group) {
	r.GET("", h.searchItem)
	r.POST("", h.insertItem)
	r.POST("/batch", h.insertItemBatch)

}

func (h *Handler) insertItem(c echo.Context) error {
	input := new(Item)

	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	err := h.Repository.InsertItem(input)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, input)
}

func (h *Handler) searchItem(c echo.Context) error {
	query := c.QueryParam("query")

	res, err := h.Repository.SearchItem(query)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	cc := c.(*mw.UserContext)

	fmt.Println("INI USERNAME DARI MIDDLEWARE DI PASSING: ", cc.GetUserName())

	return c.JSON(200, res)
}

func (h *Handler) insertItemBatch(c echo.Context) error {
	jsonFile, err := os.Open("item/mock/MOCK_DATA.json")
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	defer jsonFile.Close()

	b, _ := ioutil.ReadAll(jsonFile)

	var items []Item
	var counter int

	json.Unmarshal(b, &items)

	for _, item := range items {
		err := h.Repository.InsertItem(&item)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		counter++
	}

	return c.JSON(200, echo.Map{
		"error":         false,
		"data Inserted": counter,
	})
}
