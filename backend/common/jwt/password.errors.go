package jwt

import "errors"

// ErrPasswordMismatch is returned by ComparePassword when the hashed password
// does not match the provided plaintext password.
var ErrPasswordMismatch = errors.New("password mismatch")
