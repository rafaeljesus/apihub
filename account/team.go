package account

import (
	"encoding/json"

	"github.com/backstage/backstage/db"
	"github.com/backstage/backstage/errors"
	. "github.com/mrvdot/golang-utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Team struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"-""`
	Name     string        `json:"name"`
	Alias    string        `json:"alias"`
	Users    []string      `json:"users"`
	Owner    string        `json:"owner"`
	Services []*Service    `json:"services,omitempty"`
}

func (team *Team) Save(owner *User) error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	if team.Name == "" {
		message := "Name cannot be empty."
		return &errors.ValidationError{Message: message}
	}

	team.Users = []string{owner.Email}
	team.Owner = owner.Email
	if team.Alias == "" {
		team.Alias = GenerateSlug(team.Name)
	} else {
		team.Alias = GenerateSlug(team.Alias)
	}
	err = conn.Teams().Insert(team)
	if mgo.IsDup(err) {
		message := "Someone already has that team alias. Could you try another?"
		return &errors.ValidationError{Message: message}
	}

	return nil
}

func DeleteTeamByAlias(alias string, user *User) (*Team, error) {
	team, err := FindTeamByAlias(alias, user)
	if err != nil || team.Owner != user.Email {
		return nil, &errors.ForbiddenError{Message: errors.ErrOnlyOwnerHasPermission.Error()}
	}
	team.Services, err = FindServicesByTeam(alias)
	if err != nil {
		return nil, err
	}
	return team, team.delete()
}

func (team *Team) delete() error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Teams().RemoveId(team.Id)
	if err == mgo.ErrNotFound {
		message := "Team not found."
		return &errors.ValidationError{Message: message}
	}

	_, err = conn.Services().RemoveAll(bson.M{"team": team.Alias})
	if err != nil {
		return err
	}
	return nil
}

func (team *Team) AddUsers(emails []string) error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	var newUser bool
	var user *User
	for _, email := range emails {
		user = &User{Email: email}
		if !user.Valid() {
			continue
		}
		if _, err := team.ContainsUser(user); err != nil {
			team.Users = append(team.Users, user.Email)
			newUser = true
		}
	}
	if newUser {
		conn.Teams().Update(bson.M{"name": team.Name}, team)
	}
	return nil
}

func (team *Team) RemoveUsers(emails []string) error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	var (
		removedUsers       bool
		errOwnerNotRemoved *errors.ValidationError
		user               *User
	)
	for _, email := range emails {
		if team.Owner == email {
			message := "It is not possible to remove the owner from the team."
			errOwnerNotRemoved = &errors.ValidationError{Message: message}
			continue
		}

		user = &User{Email: email}
		if !user.Valid() {
			continue
		}
		if i, err := team.ContainsUser(user); err == nil {
			hi := len(team.Users) - 1
			if hi > i {
				team.Users[i] = team.Users[hi]
			}
			team.Users = team.Users[:hi]
			removedUsers = true
		}
	}
	if removedUsers {
		conn.Teams().Update(bson.M{"name": team.Name}, team)
	}
	if errOwnerNotRemoved != nil {
		return errOwnerNotRemoved
	}
	return nil
}

func DeleteTeamByName(name string) error {
	conn, err := db.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = conn.Teams().Remove(bson.M{"name": name})
	if err == mgo.ErrNotFound {
		message := "Team not found."
		return &errors.ValidationError{Message: message}
	}

	return nil
}

func FindTeamByName(name string) (*Team, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var team Team
	err = conn.Teams().Find(bson.M{"name": name}).One(&team)
	if err == mgo.ErrNotFound {
		message := "Team not found."
		return nil, &errors.ValidationError{Message: message}
	}
	team.Services, err = FindServicesByTeam(team.Alias)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func FindTeamByAlias(alias string, user *User) (*Team, error) {
	team, err := findTeamByAlias(alias)
	if err != nil {
		return nil, errors.ErrTeamNotFound
	}
	_, err = team.ContainsUser(user)
	if err != nil {
		return nil, &errors.ForbiddenError{Message: err.Error()}
	}
	team.Services, err = FindServicesByTeam(alias)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func findTeamByAlias(alias string) (*Team, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var team Team
	err = conn.Teams().Find(bson.M{"alias": alias}).One(&team)
	if err == mgo.ErrNotFound {
		message := "Team not found."
		return nil, &errors.ValidationError{Message: message}
	}

	return &team, nil
}

func FindTeamById(id string) (*Team, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var errNotFound = &errors.ValidationError{Message: "Team not found."}
	if !bson.IsObjectIdHex(id) {
		return nil, errNotFound
	}

	var team Team
	err = conn.Teams().FindId(bson.ObjectIdHex(id)).One(&team)
	if err != nil {
		return nil, errNotFound
	}

	team.Services, err = FindServicesByTeam(team.Alias)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (team *Team) GetTeamUsers() ([]*User, error) {
	conn, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var users []*User
	var user *User
	for _, email := range team.Users {
		user, _ = FindUserByEmail(email)
		users = append(users, user)
	}

	return users, nil
}

func getEmails(users []*User) []string {
	emails := make([]string, len(users))
	for i, u := range users {
		emails[i] = u.Email
	}
	return emails
}

func (team *Team) ContainsUser(user *User) (int, error) {
	for i, u := range team.Users {
		if u == user.Email {
			return i, nil
		}
	}
	return -1, errors.ErrUserNotInTeam
}

//Return a representation of a team without sensitive data.
func (team *Team) ToString() string {
	team.Id = ""
	t, _ := json.Marshal(team)
	return string(t)
}
