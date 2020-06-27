package models

// Relation : Graba la relación de un usuario con otro
type Relation struct {
	UserID             string `bson:"userID" json:"userID"`
	UserRelationshipID string `bson:"userRelationshipID" json:"userRelationshipID"`
}
