package enums

// USER ROLE
const (
	USERRole  = "USER"
	ADMINRole = "ADMIN"
)

// SALARY CALCULATIONS CONST
const (

	// UserWorkHours defines the standard number of working hours in a day.
	UserWorkHours uint64 = 8

	// UserWorkDays defines the standard number of working days in a month.
	UserWorkDays uint64 = 20

	// UserOvertimeMultiplier defines the multiplier for overtime duration.
	UserOvertimeMultiplier uint64 = 2
)

// AUTHORIZATION KEYS
const (
	AuthUserCtxKey  = "auth_user"
	SessionIDCtxKey = "session_id"
	RequestIDCtxKey = "request_id"
)
