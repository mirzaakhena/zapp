package model

// =============================================

// RegisterResponse is
type RegisterResponse struct {
	RegisterToken string `json:"registerToken"` //
}

// LoginResponse is
type LoginResponse struct {
	LoginToken string `json:"loginToken"` //
}

// ActivateResponse is
type ActivateResponse struct{}

// ForgotPasswordInitResponse is
type ForgotPasswordInitResponse struct {
	ResetPasswordToken string `json:"resetPasswordToken"` //
}

// ForgotPasswordResetResponse is
type ForgotPasswordResetResponse struct{}

// UpdateStatusResponse is
type UpdateStatusResponse struct{}

// GenerateInvitationAccountResponse is
type GenerateInvitationAccountResponse struct {
	SpaceInvitationToken string
}

// UpdatePasswordResponse is
type UpdatePasswordResponse struct{}

// CreateSpaceResponse is
type CreateSpaceResponse struct{}

// GetAllPermissionResponse is
type GetAllPermissionResponse struct{}

// GetBasicUserInfoResponse is
type GetBasicUserInfoResponse struct {
	User *User `json:"user"` //
}

// UpdateBasicUserInfoResponse is
type UpdateBasicUserInfoResponse struct {
	Firstname string `json:"firstName"` //
	Lastname  string `json:"lastName"`  //
	Email     string `json:"email"`     //
	Phone     string `json:"phone"`     //
	Address   string `json:"address"`   //
}

// RemoveAccountResponse is
type RemoveAccountResponse struct{}

// RemoveWaitingAccountResponse is
type RemoveWaitingAccountResponse struct{}

// GetAllUserRolePermissionResponse is
type GetAllUserRolePermissionResponse struct{}

// CreateUserRolePermissionResponse is
type CreateUserRolePermissionResponse struct{}

// UpdateUserRolePermissionResponse is
type UpdateUserRolePermissionResponse struct{}

// DeleteUserRolePermissionResponse is
type DeleteUserRolePermissionResponse struct{}

// GetAllAccountUserRoleResponse is
type GetAllAccountUserRoleResponse struct{}

// UpdateAccountUserRoleResponse is
type UpdateAccountUserRoleResponse struct{}

// ApplyForPermissionResponse is
type ApplyForPermissionResponse struct{}

// GetAllAppliedPermissionResponse is
type GetAllAppliedPermissionResponse struct{}

// GrantAppliedPermissionResponse is
type GrantAppliedPermissionResponse struct{}

// RefuseAppliedPermissionResponse is
type RefuseAppliedPermissionResponse struct{}

// IsAccessableResponse is
type IsAccessableResponse struct{}

// ServiceContext is
type ServiceContext map[string]interface{}

// DaoContext is database implementation
type DaoContext interface{}

// ==============================================================================================================

// BaseErrorResponse is
type BaseErrorResponse struct {
	error
	Code int `` //
}

const (
	_ = iota + 4000

	// ConstErrorInternal is
	ConstErrorInternal

	// ConstErrorUnExistingEmailAddress is
	ConstErrorUnExistingEmailAddress

	// ConstErrorExistingEmailAddress is
	ConstErrorExistingEmailAddress

	// ConstErrorExistingEmailAddressButNotActiveYet is
	ConstErrorExistingEmailAddressButNotActiveYet

	// ConstErrorWhenSendingEmail is
	ConstErrorWhenSendingEmail

	// ConstErrorLogin is
	ConstErrorLogin

	// ConstErrorInvalidPassword is
	ConstErrorInvalidPassword

	// ConstErrorForgotPassword is
	ConstErrorForgotPassword

	// ConstErrorActivation is
	ConstErrorActivation

	// ConstErrorTokenNotFound is
	ConstErrorTokenNotFound

	// ConstErrorInvalidExpectedStatus is
	ConstErrorInvalidExpectedStatus
)
