package data

import (
	"time"

	"github.com/supabase-community/supabase-go"
)

type Comment struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Picture   string    `json:"picture"`
	Status    string    `json:"status"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
}

func ValidateComment() {

}

type CommentModel struct {
	DB *supabase.Client
}

/* func (m CommentModel) Create(ctx context.Context, c *Comment) error {
	resp, err := m.DB.
		From("comments").
		Insert(c, false, "", "").
		Execute()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return supabase.ParseJSON(resp, &c)
}

func (m CommentModel) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	resp, err := m.DB.
		From("comments").
		Select("*", "exact", false).
		Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var comments []map[string]interface{}
	if err := supabase.ParseJSON(resp, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
*/
