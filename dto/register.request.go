package dto

/*
Important notes:
I use @DateOfBirth data_type as 'string' rather than time.Time due to
automatically parse by golang-gin that handles error not too properly.

With this string data type, i parsed the date manually inside the controller.
I think, it is much better than a blackmagic that gin user inside ctx.ShouldBind
*/

type UserRegisterRequest struct {
	Email       string `json:"email" form:"email" binding:"required,email"`
	Password    string `json:"password" form:"password" binding:"required,min=8"`
	Username    string `json:"username" form:"username" binding:"required"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth" binding:"required"`
}
