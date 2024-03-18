// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "filmlib/server/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(username, password string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", username, password)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), username, password)
}

// MockActor is a mock of Actor interface.
type MockActor struct {
	ctrl     *gomock.Controller
	recorder *MockActorMockRecorder
}

// MockActorMockRecorder is the mock recorder for MockActor.
type MockActorMockRecorder struct {
	mock *MockActor
}

// NewMockActor creates a new mock instance.
func NewMockActor(ctrl *gomock.Controller) *MockActor {
	mock := &MockActor{ctrl: ctrl}
	mock.recorder = &MockActorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActor) EXPECT() *MockActorMockRecorder {
	return m.recorder
}

// CreateActor mocks base method.
func (m *MockActor) CreateActor(actor entity.Actor) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", actor)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockActorMockRecorder) CreateActor(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockActor)(nil).CreateActor), actor)
}

// DeleteActorById mocks base method.
func (m *MockActor) DeleteActorById(actorId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActorById", actorId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActorById indicates an expected call of DeleteActorById.
func (mr *MockActorMockRecorder) DeleteActorById(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActorById", reflect.TypeOf((*MockActor)(nil).DeleteActorById), actorId)
}

// GetActor mocks base method.
func (m *MockActor) GetActor(actorId int) (entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", actorId)
	ret0, _ := ret[0].(entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockActorMockRecorder) GetActor(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockActor)(nil).GetActor), actorId)
}

// GetActorsIdByFilmId mocks base method.
func (m *MockActor) GetActorsIdByFilmId(filmId int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsIdByFilmId", filmId)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsIdByFilmId indicates an expected call of GetActorsIdByFilmId.
func (mr *MockActorMockRecorder) GetActorsIdByFilmId(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsIdByFilmId", reflect.TypeOf((*MockActor)(nil).GetActorsIdByFilmId), filmId)
}

// UpdateActorById mocks base method.
func (m *MockActor) UpdateActorById(actorId int, actor entity.UpdateActorInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActorById", actorId, actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActorById indicates an expected call of UpdateActorById.
func (mr *MockActorMockRecorder) UpdateActorById(actorId, actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActorById", reflect.TypeOf((*MockActor)(nil).UpdateActorById), actorId, actor)
}

// MockFilm is a mock of Film interface.
type MockFilm struct {
	ctrl     *gomock.Controller
	recorder *MockFilmMockRecorder
}

// MockFilmMockRecorder is the mock recorder for MockFilm.
type MockFilmMockRecorder struct {
	mock *MockFilm
}

// NewMockFilm creates a new mock instance.
func NewMockFilm(ctrl *gomock.Controller) *MockFilm {
	mock := &MockFilm{ctrl: ctrl}
	mock.recorder = &MockFilmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilm) EXPECT() *MockFilmMockRecorder {
	return m.recorder
}

// CreateFilm mocks base method.
func (m *MockFilm) CreateFilm(film entity.Film) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilm", film)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFilm indicates an expected call of CreateFilm.
func (mr *MockFilmMockRecorder) CreateFilm(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilm", reflect.TypeOf((*MockFilm)(nil).CreateFilm), film)
}

// DeleteFilmById mocks base method.
func (m *MockFilm) DeleteFilmById(filmId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilmById", filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilmById indicates an expected call of DeleteFilmById.
func (mr *MockFilmMockRecorder) DeleteFilmById(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilmById", reflect.TypeOf((*MockFilm)(nil).DeleteFilmById), filmId)
}

// GetFilmsByActorId mocks base method.
func (m *MockFilm) GetFilmsByActorId(actorId int) ([]entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsByActorId", actorId)
	ret0, _ := ret[0].([]entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsByActorId indicates an expected call of GetFilmsByActorId.
func (mr *MockFilmMockRecorder) GetFilmsByActorId(actorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsByActorId", reflect.TypeOf((*MockFilm)(nil).GetFilmsByActorId), actorId)
}

// GetFilmsWithFragment mocks base method.
func (m *MockFilm) GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsWithFragment", actorNameFrag, filmNameFrag)
	ret0, _ := ret[0].([]entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsWithFragment indicates an expected call of GetFilmsWithFragment.
func (mr *MockFilmMockRecorder) GetFilmsWithFragment(actorNameFrag, filmNameFrag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsWithFragment", reflect.TypeOf((*MockFilm)(nil).GetFilmsWithFragment), actorNameFrag, filmNameFrag)
}

// GetFilmsWithSort mocks base method.
func (m *MockFilm) GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsWithSort", sortType, filmsId)
	ret0, _ := ret[0].([]entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsWithSort indicates an expected call of GetFilmsWithSort.
func (mr *MockFilmMockRecorder) GetFilmsWithSort(sortType, filmsId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsWithSort", reflect.TypeOf((*MockFilm)(nil).GetFilmsWithSort), sortType, filmsId)
}

// UpdateFilmById mocks base method.
func (m *MockFilm) UpdateFilmById(filmId int, deleteIds, addIds []int, film entity.UpdateFilmInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilmById", filmId, deleteIds, addIds, film)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilmById indicates an expected call of UpdateFilmById.
func (mr *MockFilmMockRecorder) UpdateFilmById(filmId, deleteIds, addIds, film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilmById", reflect.TypeOf((*MockFilm)(nil).UpdateFilmById), filmId, deleteIds, addIds, film)
}
