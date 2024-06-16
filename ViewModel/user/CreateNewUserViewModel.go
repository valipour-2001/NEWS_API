package user

type CreateNewUserViewModel struct {
	LastName      string `validate:"required"`
	FirstName     string
	Email         string
	UserName      string
	Password      string
	CreatorUserId string
	AvatarName    string
}
type EditUserViewModel struct {
	TargetUserId string
	LastName     string `validate:"required"`
	FirstName    string `validate:"required"`
	Email        string `validate:"required"`
	UserName     string `validate:"required"`
	Password     string `validate:"required"`
}
type EditUserRoleViewModel struct {
	TargetUserId string
	Roles        []string `validate:"required"`
}
type EditUserPasswordViewModel struct {
	TargetUserId string
	Password     string `validate:"required"`
}
