package c

// import (
// 	"log"

// 	"golang.org/x/net/context"

// 	firebase "firebase.google.com/go"

// 	"google.golang.org/api/option"
// )

// func VerifyFirebaseToken(idToken string) {
// 	client, err := app.Auth(context.Background())
// 	if err != nil {
// 		log.Fatalf("error getting Auth client: %v\n", err)
// 	}

// 	token, err := client.VerifyIDToken(idToken)
// 	if err != nil {
// 		log.Fatalf("error verifying ID token: %v\n", err)
// 	}

// 	log.Printf("Verified ID token: %v\n", token)
// }

// func InitializeSDK() {
// 	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
// 	}
// }
