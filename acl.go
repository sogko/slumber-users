package users

import (
	"github.com/gorilla/mux"
	"github.com/sogko/slumber/domain"
	"net/http"
)

func (resource *Resource) HandleListUsersACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	return true, ""
}

func (resource *Resource) HandleGetUserACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous to get user information
	return true, ""
}

func (resource *Resource) HandleCreateUserACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous to create a user account
	// if authenticated, only admin can create new users
	// no point for non-admin to create new users
	// TODO: only allow authorized but unauthenticated client
	if user == nil {
		// enforce authenticated access
		return true, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if !u.HasRole(RoleAdmin) {
		// must have an admin role
		return false, ""
	}
	return true, ""
}

func (resource *Resource) HandleUpdateUsersACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if !u.HasRole(RoleAdmin) {
		// must have an admin role
		return false, ""
	}
	// only logged-in admins can update users in batch
	return true, ""
}

func (resource *Resource) HandleDeleteAllUsersACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if !u.HasRole(RoleAdmin) {
		// must have an admin role
		return false, ""
	}
	// only logged-in admins can update users in batch
	return true, ""
}

func (resource *Resource) HandleConfirmUserACL(req *http.Request, user domain.IUser) (bool, string) {
	// allow anonymous access. user is expected to specify `code` (business logic)
	return true, ""
}

func (resource *Resource) HandleUpdateUserACL(req *http.Request, user domain.IUser) (bool, string) {
	params := mux.Vars(req)
	id := params["id"]
	repo := resource.UserRepository(req)

	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if u.HasRole(RoleAdmin) {
		// must have an admin role
		return true, ""
	}

	// retrieve target user
	_userTarget, err := repo.GetUserById(id)
	if err != nil {
		return false, "Invalid user"
	}
	userTarget := _userTarget.(*User)
	if userTarget != nil && u.ID == userTarget.ID {
		// this is his own account
		return true, ""
	}
	// a user can only `update` its own user account or if user is an admin
	return false, ""
}

func (resource *Resource) HandleDeleteUserACL(req *http.Request, user domain.IUser) (bool, string) {
	// only an admin can `delete` a user account
	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if !u.HasRole(RoleAdmin) {
		// must have an admin role
		return false, ""
	}
	// only logged-in admins can update users in batch
	return true, ""
}

func (resource *Resource) HandleCountUsersACL(req *http.Request, user domain.IUser) (bool, string) {
	if user == nil {
		// enforce authenticated access
		return false, ""
	}
	u := user.(*User)
	if u.Status != StatusActive {
		// must be an active user
		return false, ""
	}
	if !u.HasRole(RoleAdmin) {
		// must have an admin role
		return false, ""
	}
	// only logged-in admins can update users in batch
	return true, ""
}
