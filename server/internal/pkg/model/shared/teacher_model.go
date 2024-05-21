package shared_model

type TeacherResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
}
