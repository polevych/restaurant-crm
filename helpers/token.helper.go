package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"

	"github.com/polevych/restaurant-crm/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.Claims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, uid string)(accessToken string, resfreshToken string, err error) {
	
	accessClaims := jwt.MapClaims{
		"Email":  email,
		"First_name":  firstName,
		"lastName":  lastName,
		"Uid": uid,
		"ExpiresAt":  time.Now().Add(time.Duration(1) * time.Hour),
	}
	
	refreshClaims := jwt.MapClaims{}
	for key, value := range accessClaims {
		refreshClaims[key] = value
	}
	refreshClaims["ExpiresAt"] = time.Now().Local().Add(time.Duration(24) * time.Hour).Unix()
	
	accessTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err := accessTokenString.SignedString([]byte(SECRET_KEY))

	refreshTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refresh, err := refreshTokenString.SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return access, refresh, err
}

func UpdateAllTokens(accessToken string, refreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updatedObject primitive.D

	updatedObject = append(updatedObject, bson.E{"token", accessToken})
	updatedObject = append(updatedObject, bson.E{"refresh_token", refreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedObject = append(updatedObject, bson.E{"updated_at", Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updatedObject},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	return

}

func ValidateToken(accessToken string) (claims *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(accessToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if claims, ok := token.Claims.(*SignedDetails); !ok || !token.Valid {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()	
		return claims, msg
	}else{
		return claims, msg
	}
}