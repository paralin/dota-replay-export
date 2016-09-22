package csv

import (
	"encoding/csv"
	"github.com/paralin/replayexport/pkg/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"strconv"
	"time"
)

func ExportSubmissions(writer io.Writer, collection *mgo.Collection, since, minRating int) error {
	query := bson.M{"status": bson.M{"$in": []int{4, 6}}}
	if since != 0 {
		query["createdAt"] = bson.M{"$gt": time.Unix(int64(since), 0)}
	}
	if minRating != 0 {
		query["rating"] = bson.M{"$gt": minRating}
	}
	var result model.Submission
	iter := collection.Find(query).Sort("-createdAt").Iter()

	csw := csv.NewWriter(writer)
	csw.Write([]string{
		"Rating",
		"Show",
		"Match ID",
		"Submitter",
		"Hero to Watch",
		"Reviewer Description",
	})
	for iter.Next(&result) {
		csw.Write([]string{
			strconv.Itoa(result.Rating),
			result.Show,
			strconv.FormatInt(result.MatchId, 10),
			result.UserName,
			result.HeroToWatch,
			result.ReviewerDescription,
		})
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
