package v1

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"my-echo/common"
	"my-echo/db"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	UserCollection = "users"
)

type User struct {
	Id        string    `json:"id" bson:"id"`
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Username  string    `json:"username" bson:"username"`
	Number    string    `json:"number" bson:"number"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

var UserListData []User = []User{
	{
		Id:        "1",
		FirstName: "Torine",
		LastName:  "Walker",
		Username:  "torine.walker@gmail.com",
		Number:    "01234567890",
	},
}

// Router
func (u User) FindAll(c echo.Context) error {
	objects := []User{}

	query := bson.M{}
	coll := db.GetDmManager().Db.Collection(UserCollection)

	cursor, err := coll.Find(db.GetDmManager().Ctx, query)

	if err != nil {
		log.Println("Error in find coll: ", err.Error())
	}
	for cursor.Next(context.TODO()) {
		elemValue := new(User)
		log.Println("elemValue", elemValue)
		err := cursor.Decode(elemValue)

		if err != nil {
			log.Println("[Error]", err.Error())
			break
		}
		objects = append(objects, *elemValue)
	}

	return common.GenerateSuccessResponse(c, objects, "Success! User list")
}

func (u User) FindById(c echo.Context) error {
	userId := c.Param("id")
	if userId == "" {
		log.Println("Id is valid")
		return common.GenerateErrorResponse(c, "[ERROR]: No username is provided!", "Please provide a user id!")
	}

	var user, found = u.findById(userId)

	if found {
		return common.GenerateSuccessResponse(c, user, "User is found!")
	}
	log.Println("User is not found")
	return common.GenerateNotFoundResponse(c, "[ERROR]: User is not found!", "")
}

func (u User) AddUser(c echo.Context) error {
	formData := User{}

	if err := c.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(c, err.Error(), "Failed to Bind Input!")
	}
	formData.Id = common.GenerateRandomString(10)
	formData.CreatedAt = time.Now().UTC()
	if err := u.addUser(formData); err != nil {
		return common.GenerateErrorResponse(c, err.Error(), "Failed! Can't add user.")
	}
	return common.GenerateSuccessResponse(c, formData, "Success! User added")
}

// End Route

func (u User) addUser(payload User) error {
	UserListData = append(UserListData, payload)
	return nil
}

func (u User) findById(id string) (*User, bool) {
	for _, item := range UserListData {
		if item.Id == id {
			return &item, true
		}
	}
	return nil, false
}
