package entities

type User struct {
	Id         string `firestore:"id"`
	FamilyName string `firestore:"familyName"`
	FirstName  string `firestore:"firstName"`
}
