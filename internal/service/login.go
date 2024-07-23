package service

import "log"

func (s *Service) CheckEmail(email string) (bool, error) {
	exists, err := s.dao.UserExistsByEmail(email)
	if err != nil {
		log.Printf("debug:x error: (%s)", err)
		return false, err
	}

	return exists, nil
}
