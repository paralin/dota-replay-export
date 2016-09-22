package model

import (
	"time"
)

type SubmissionStatus int

type Submission struct {
	Id                    string           `json:"_id" bson:"_id"`
	LegacyUsed            bool             `json:"legacyUsed"`
	Rating                int              `json:"rating"`
	Reviewer              string           `json:"reviewer"`
	ReviewerUntil         time.Time        `json:"reviewerUntil"`
	Reviewed              bool             `json:"reviewed"`
	ReviewerDescription   string           `json:"reviewerDescription" bson:"reviewerDescription"`
	Description           string           `json:"description"`
	MatchId               int64            `json:"matchid" bson:"matchid"`
	MatchTime             int              `json:"matchtime" bson:"matchtime"`
	Show                  string           `json:"show" bson:"show"`
	Status                SubmissionStatus `json:"status" bson:"status"`
	UserId                string           `json:"uid" bson:"uid"`
	UserName              string           `json:"uname" bson:"uname"`
	CreatedAt             time.Time        `json:"createdAt" bson:"createdAt"`
	HeroToWatch           string           `json:"hero_to_watch" bson:"hero_to_watch"`
	FetchError            int              `json:"fetch_error" bson:"fetch_error"`
	FetchErrorReplayState string           `json:"fetch_error_replay_state" bson:"fetch_error_replay_state"`
	IngameTime            string           `json:"ingame_time" bson:"ingame_time"`
}
