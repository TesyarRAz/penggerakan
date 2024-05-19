package errors

type Unauthorized struct {
	Message string
}

func NewUnathorized() Unauthorized {
	return Unauthorized{
		Message: "Unauthorized",
	}
}

func (e Unauthorized) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func NewInternalServerError() InternalServerError {
	return InternalServerError{
		Message: "Internal server error",
	}
}
func (e InternalServerError) Error() string {
	return e.Message
}

type NotFound struct {
	Message string
}

func NewNotFound() NotFound {
	return NotFound{
		Message: "Not found",
	}
}

func (e NotFound) Error() string {
	return e.Message
}

type BadRequest struct {
	Message string
}

func NewBadRequest() BadRequest {
	return BadRequest{
		Message: "Bad request",
	}
}

func (e BadRequest) Error() string {
	return e.Message
}

type Conflict struct {
	Message string
}

func NewConflict() Conflict {
	return Conflict{
		Message: "Conflict",
	}
}

func (e Conflict) Error() string {
	return e.Message
}
