package store

import (
	"time"
)

type Status uint8

type User struct {
	ID        uint64     `json:"id"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	Address   *string    `json:"address"`
	Status    Status     `json:"status"`
	BirthDay  *time.Time `json:"birth_day"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UserDao interface {
	// insert ignore into users(`username`, phone, address, status, birth_day, created, updated)
	// values (?,?,?,?,?,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	Insert(u *User) (int64, error)

	// insert into users(username, phone, address, status, birth_day, created, updated)
	// values (?,?,?,?,?,CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	// on duplicate key update
	//   username=values(username), phone=values(phone), address=values(address),
	//   status=values(status), birth_day=values(birth_day), updated=CURRENT_TIMESTAMP
	Upsert(u *User) (int64, error)

	// UPDATE users
	// SET [username=?,]
	//     [phone=?,]
	//     [address=?,]
	//     [status=?,]
	//     [birth_day=?,]
	//     updated=CURRENT_TIMESTAMP
	// WHERE id=?
	Update(u *User) (int64, error)

	// DELETE FROM users WHERE id=?
	Delete(id uint64) (int64, error)

	// select id, username, phone, address, status, birth_day, created, updated
	// FROM users WHERE id=?
	Get(id uint64) (*User, error)

	// select count(1)
	// from users
	Count() (int64, error)

	// select (select id from users where id=a.id) as id,
	// `username`, phone as phone, address, status, birth_day, created, updated
	// from users a
	// where id != -1 and  username <> 'admin' and username like ?
	// [
	// 	and address = ?
	// 	[and phone like ?]
	// 	and created > ?
	//  [{(u.BirthDay != nil && !u.BirthDay.IsZero()) || u.Id > 1 }
	//   [and birth_day > ?]
	//   [and id > ?]
	//  ]
	// ]
	// and status != ?
	// [and updated > ?]
	// and birth_day is not null
	// order by updated desc
	// limit ${offset}, ${size}
	List(offset, size int) ([]*User, error)
}