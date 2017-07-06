package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/bxcodec/go-clean-arch/models"

// ArticleUsecase is an autogenerated mock type for the ArticleUsecase type
type ArticleUsecase struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: cursor, num
func (_m *ArticleUsecase) Fetch(cursor string, num int64) ([]*models.Article, string, error) {
	ret := _m.Called(cursor, num)

	var r0 []*models.Article
	if rf, ok := ret.Get(0).(func(string, int64) []*models.Article); ok {
		r0 = rf(cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Article)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, int64) string); ok {
		r1 = rf(cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int64) error); ok {
		r2 = rf(cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *ArticleUsecase) GetByID(id int64) (*models.Article, error) {
	ret := _m.Called(id)

	var r0 *models.Article
	if rf, ok := ret.Get(0).(func(int64) *models.Article); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}