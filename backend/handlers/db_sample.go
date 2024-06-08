package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ken-harada/household-accounts/models"
	"github.com/labstack/echo/v4"
)

func DBSample(c echo.Context) error {
	db, err := models.ConnectDB()

	if c.FormValue("sample_search") != "" {
		fmt.Printf("%+v\n", c.FormValue("sample_search"))
	} else {
		fmt.Printf("%+v\n", "空文字")
	}

	if err != nil {
		log.Fatal(err)
	}
	sample, err := models.FindSampleAll(db)
	for _, v := range sample {
		fmt.Printf("%+v\n", v.Id)
		fmt.Printf("%+v\n", v.Name)
	}
	if err != nil {
		log.Fatal(err)
	}

	return c.Render(http.StatusOK, "sample_index.html", sample)
}

func SampleInput(c echo.Context) error {
	return c.Render(http.StatusOK, "sample_input.html", nil)
}

func SampleRegister(c echo.Context) error {
	sample := models.Sample{
		Name: c.FormValue("sample_user_name"),
	}
	db, err := models.ConnectDB()

	if err != nil {
		return err
	}
	result := db.Create(&sample)
	if result.Error != nil {
		panic("failed to create travel brochure")
	}
	return c.Redirect(http.StatusMovedPermanently, "/db_sample_index")
}
