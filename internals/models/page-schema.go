package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// PageType defines the type of content a page holds.
type PageType string

const (
	NotesType   PageType = "notes"
	MCQType     PageType = "mcq"
	YouTubeType PageType = "ytvideo"
	CodingType  PageType = "coding"
)

// Page represents the page that will be returned from the database
type Page struct {
	ID        primitive.ObjectID `json:"id"`
	TopicID   string             `json:"topic_id"`
	Title     string             `json:"title"`
	Type      string             `json:"type"`
	Content   interface{}        `json:"content"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	DeletedAt string             `json:"deleted_at,omitempty"`
	Hidden    bool               `json:"isHidden"`
}

// NotesContent represents the content of a notes page.
type NotesContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// MCQContent represents the content of a multiple-choice question page.
type MCQContent struct {
	Question string        `json:"question"`
	Image    string        `json:"image"`
	Options  []interface{} `json:"options"`
	Correct  int           `json:"correct"`
}

// YouTubeContent represents the content of a YouTube video page.
type YouTubeContent struct {
	VideoID string `json:"video_id"` // YouTube video ID
	Title   string `json:"title"`    // Title of the video
}

// CodingContent represents the content of a coding problem page.
type CodingContent struct {
	ProblemStatement  string   `json:"problem_statement"`
	TestCases         []string `json:"test_cases"`
	SolutionTestCases []string `json:"solutions"`
}
