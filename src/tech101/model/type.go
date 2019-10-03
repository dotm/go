package model

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Price       int64     `json:"price"`
}

type PollResponse struct {
	Header struct {
		ProcessTime float64  `json:"process_time"`
		Messages    []string `json:"messages"`
		Reason      string   `json:"reason"`
		ErrorCode   string   `json:"error_code"`
	} `json:"header"`
	Data struct {
		Poll struct {
			PollID      int    `json:"poll_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Question    string `json:"question"`
			Options     []struct {
				OptionID     int    `json:"option_id"`
				Option       string `json:"option"`
				ImageOption  string `json:"image_option"`
				IsFifCorrect bool   `json:"is_fif_correct"`
			} `json:"options"`
			PollTypeID   int    `json:"poll_type_id"`
			PollType     string `json:"poll_type"`
			OptionTypeID int    `json:"option_type_id"`
			OptionType   string `json:"option_type"`
			StatusID     int    `json:"status_id"`
			Status       string `json:"status"`
			StartTime    int    `json:"start_time"`
			EndTime      int    `json:"end_time"`
			Statistic    struct {
				TotalVoter       int `json:"total_voter"`
				StatisticOptions []struct {
					OptionID   int    `json:"option_id"`
					Option     string `json:"option"`
					Voter      int    `json:"voter"`
					Percentage int    `json:"percentage"`
					IsSelected bool   `json:"is_selected"`
				} `json:"statistic_options"`
			} `json:"statistic"`
			IsAnswered bool   `json:"is_answered"`
			WinnerURL  string `json:"winner_url"`
		} `json:"poll"`
	} `json:"data"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}
