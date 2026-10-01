package service

import "github.com/Rizal-Nurochman/repository"

func GetAll() ([]repository.Book, error)  {
	books, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}