package apiserver

import "github.com/jonathanhu237/binding-manager/backend/internal/domain"

func (as *ApiServer) init() error {
	// Check the admin user exists or not.
	exists, err := as.repo.User.CheckAdminExists()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	as.logger.Info("No admin user found, creating first admin user")

	// Create the first admin user.
	first_admin := &domain.User{
		Username: as.cfg.FirstAdmin.Username,
		Email:    as.cfg.FirstAdmin.Email,
		IsAdmin:  true,
	}
	if err := first_admin.Password.Set(as.cfg.FirstAdmin.Password); err != nil {
		return err
	}

	// Insert the first admin user into the database.
	if err := as.repo.User.Insert(first_admin); err != nil {
		return err
	}

	return nil
}
