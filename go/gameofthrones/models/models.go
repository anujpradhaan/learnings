package models

import "time"

type Root struct {
	Books      string `json:"books"`
	Characters string `json:"characters"`
	Houses     string `json:"houses"`
}
type Book struct {
	URL           string    `json:"url"`           //The hypermedia URL of this resource
	Name          string    `json:"name"`          //The name of this book
	Isbn          string    `json:"isbn"`          //The International Standard Book Number (ISBN-13) that uniquely identifies this book.
	Authors       []string  `json:"authors"`       //An array of names of the authors that wrote this book.
	NumberOfPages int       `json:"numberOfPages"` //The number of pages in this book.
	Publisher     string    `json:"publisher"`     //The company that published this book.
	Country       string    `json:"country"`       //The country that this book was published in
	MediaType     string    `json:"mediaType"`     //The type of media this book was released in.
	Released      time.Time `json:"released"`      //The date (ISO 8601) when this book was released.
	Characters    []string  `json:"characters"`    //An array of Character resource URLs that has been in this book.
	PovCharacters []string  `json:"povCharacters"` //An array of Character resource URLs that has had a POV-chapter in this book.
}

type Character struct {
	URL         string   `json:"url"`         // The hypermedia URL of this resource
	Name        string   `json:"name"`        // The name of this character
	Gender      string   `json:"gender"`      //The gender of this character.
	Culture     string   `json:"culture"`     //The culture that this character belongs to.
	Born        string   `json:"born"`        //Textual representation of when and where this character was born.
	Died        string   `json:"died"`        //Textual representation of when and where this character died.
	Titles      []string `json:"titles"`      //TThe titles that this character holds.
	Aliases     []string `json:"aliases"`     //The aliases that this character goes by.
	Father      string   `json:"father"`      //The character resource URL of this character's father.
	Mother      string   `json:"mother"`      //The character resource URL of this character's mother.
	Spouse      string   `json:"spouse"`      //An array of Character resource URLs that has had a POV-chapter in this book.
	Allegiances []string `json:"allegiances"` //An array of House resource URLs that this character is loyal to.
	Books       []string `json:"books"`       //An array of Book resource URLs that this character has been in.
	PovBooks    []string `json:"povBooks"`    //An array of Book resource URLs that this character has had a POV-chapter in.
	TvSeries    []string `json:"tvSeries"`    //An array of names of the seasons of Game of Thrones that this character has been in.
	PlayedBy    []string `json:"playedBy"`    //An array of actor names that has played this character in the TV show Game Of Thrones.

}

type House struct {
	URL              string   `json:"url"`              //The hypermedia URL of this resource
	Name             string   `json:"name"`             //The name of this house
	Region           string   `json:"region"`           //The region that this house resides in.
	CoatOfArms       string   `json:"coatOfArms"`       //Text describing the coat of arms of this house.
	Words            string   `json:"words"`            //The words of this house.
	Titles           []string `json:"titles"`           //The titles that this house holds.
	Seats            []string `json:"seats"`            //The seats that this house holds.
	CurrentLord      string   `json:"currentLord"`      //The Character resource URL of this house's current lord.
	Heir             string   `json:"heir"`             //The Character resource URL of this house's heir.
	Overlord         string   `json:"overlord"`         //The House resource URL that this house answers to.
	Founded          string   `json:"founded"`          //The year that this house was founded.
	Founder          string   `json:"founder"`          //The Character resource URL that founded this house.
	DiedOut          string   `json:"diedOut"`          //The year that this house died out.
	AncestralWeapons []string `json:"ancestralWeapons"` //An array of names of the noteworthy weapons that this house owns.
	CadetBranches    []string `json:"cadetBranches"`    //An array of House resource URLs that was founded from this house.
	SwornMembers     []string `json:"swornMembers"`     //An array of Character resource URLs that are sworn to this house.
}
