package models

import "time"

// Term or Concept or Glossary is a description of some core words of their fields
//
// the fields of term was inspired by lark lingo, baidu wiki, wikipedia and so on.
type Term struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	DisplayName      string    `json:"display-name"`
	Alias            string    `json:"alias"`
	ChineseName      string    `json:"chinese-name"`
	Abbreviation     string    `json:"abbreviation"`
	Fields           []string  `json:"fields"`
	Tags             []string  `json:"tags"`
	Relevance        []string  `json:"relevance"` // 相关词
	StrictDefinition string    `json:"strict-definition"`
	DateCreated      time.Time `json:"date-created"`
	DateModified     time.Time `json:"date-modified"`
}
