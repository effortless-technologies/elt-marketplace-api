package clients

import (
	"net/http"
	"time"

	"github.com/effortless-technologies/elt-marketplace-api/models"

	"github.com/labstack/echo"
	"github.com/dghubble/sling"

)

const PostmatesClientTimeout = 5 * time.Second
const PostmatesBaseUrl = "https://api.postmates.com"
const PostmatesAPIKey = "0c55e671-dcd6-4d95-b237-abfe27952877"

func Authenticate() error {

	// TODO: implement

	return nil
}

func GetDeliveryQuote(c echo.Context) error {

	postmates_base := sling.New().Base(
		PostmatesBaseUrl).SetBasicAuth(
		PostmatesAPIKey,"").Set(
		"Accept", "application/json").Set(
		"Content-Type", "application/json")

	type form_params struct {
		PickupAddress		string		`url:"pickup_address"`
		DropoffAddress 		string 		`url:"dropoff_address"`
	}

	params := &form_params{
		PickupAddress: "20 McAllister St, San Francisco, CA",
		DropoffAddress: "101 Market St, San Francisco, CA",
	}

	response := &models.Quote{}
	req, err := postmates_base.New().Post(
		"/v1/customers/cus_LSo5Dq0t8ppZFF/delivery_quotes").
			BodyForm(params).Request()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, err = postmates_base.Do(req, &response, err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, response)
}
