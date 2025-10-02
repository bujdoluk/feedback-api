package data

import (
	"time"

	"github.com/supabase-community/supabase-go"
)

type Reply struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Picture   string    `json:"picture"`
	Text      string    `json:"text"`
}

type ReplyModel struct {
	DB *supabase.Client
}

/* func (m ReplyModel) Create(ctx context.Context, r *Reply) error {
	resp, err := m.DB.
		From("replies").
		Insert(r, false, "", "").
		Execute()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return supabase.ParseJSON(resp, &r)
}

func (m ReplyModel) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	resp, err := m.DB.
		From("replies").
		Select("*", "exact", false).
		Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var replies []map[string]interface{}
	if err := supabase.ParseJSON(resp, &replies); err != nil {
		return nil, err
	}

	return replies, nil
}
*/
