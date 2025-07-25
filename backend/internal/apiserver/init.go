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
	passwordHash, err := domain.GeneratePasswordHash(as.cfg.FirstAdmin.Password)
	if err != nil {
		return err
	}

	first_admin := &domain.User{
		Username:     as.cfg.FirstAdmin.Username,
		PasswordHash: passwordHash,
		IsAdmin:      true,
	}

	// Insert the first admin user into the database.
	if err := as.repo.User.Insert(first_admin); err != nil {
		return err
	}

	return nil
}
