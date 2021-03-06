//go:generate gobatis users.go
package example_xml

import "github.com/runner-mei/GoBatis/tests"

type Users interface {
	FindByID(id int64) (*tests.User, error)

	Insert(u *tests.User) (int64, error)

	Update(id int64, u *tests.User) (int64, error)

	DeleteByID(id int64) (int64, error)

	DeleteAll() (int64, error)
}
