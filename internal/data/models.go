package data

import (
	"github.com/supabase-community/supabase-go"
)

type Models struct {
	Suggestions SuggestionModel
	Permissions PermissionModel
	Users       UserModel
	Comments    CommentModel
	Replies     ReplyModel
}

func NewModels(db *supabase.Client) Models {
	return Models{
		Suggestions: SuggestionModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Users:       UserModel{DB: db},
		Comments:    CommentModel{DB: db},
		Replies:     ReplyModel{DB: db},
	}
}
