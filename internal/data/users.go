package data

import (
	"time"

	"github.com/supabase-community/supabase-go"
)

type password struct {
	hash      []byte
	plaintext *string
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"version"`
}

type UserModel struct {
	DB *supabase.Client
}

/* func (m UserModel) Create(ctx context.Context, u *User) error {
	resp, err := m.DB.
		From("users").
		Insert(u, false, "", "").
		Execute()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return supabase.ParseJSON(resp, &u)
}

func (m UserModel) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	resp, err := m.DB.
		From("users").
		Select("*", "exact", false).
		Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []map[string]interface{}
	if err := supabase.ParseJSON(resp, &users); err != nil {
		return nil, err
	}

	return users, nil
}
*/
