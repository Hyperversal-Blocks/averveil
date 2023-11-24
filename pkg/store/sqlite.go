package store

func (s *store) InsertZipRecord(zipFileName, fileName string) error {
	// Prepare the insert statement
	s.sqlMutex.Lock()
	defer s.sqlMutex.Unlock()
	stmt, err := s.sqlDB.Prepare("INSERT INTO zipped(zipFileName, fileName) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided values
	_, err = stmt.Exec(zipFileName, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) DeleteZipRecord(zipFileName string) error {
	// Prepare the insert statement
	s.sqlMutex.Lock()
	defer s.sqlMutex.Unlock()
	stmt, err := s.sqlDB.Prepare("DELETE from zipped WHERE zipFileName = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided values
	_, err = stmt.Exec(zipFileName)
	if err != nil {
		return err
	}

	return nil
}
