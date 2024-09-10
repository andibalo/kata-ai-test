package request

type RegisterUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	ChannelUsername string `json:"channel_username"`
	ChannelType     string `json:"channel_type"`
}

type LoginRequest struct {
	ChannelUsername string `json:"channel_username"`
	ChannelType     string `json:"channel_type"`
}
