package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

type Client struct {
	cluster *gocql.ClusterConfig

	session *gocql.Session
}

type User struct {
	Email string
	Name string
}

func NewClient(hosts... string) (*Client, error) {
	cluster := gocql.NewCluster(hosts...)
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	return &Client{cluster: cluster, session: session}, nil
}

func (c *Client) InitSchema() error {
	log.Println("initializing schema")
	if err := c.session.Query("CREATE KEYSPACE IF NOT EXISTS medusa_test WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 3}").Exec(); err != nil {
		return err
	}

	return c.session.Query("CREATE TABLE IF NOT EXISTS medusa_test.users (email text PRIMARY KEY, name text)").Exec()
}

func (c *Client) CreateUser(user User) error {
	return c.session.Query(`INSERT INTO medusa_test.users (email, name) VALUES (?, ?)`, user.Email, user.Name).Exec()
}

func (c *Client) GetUsers() ([]User, error) {
	var email, name string
	users := make([]User, 0)

	iter := c.session.Query("SELECT email, name FROM medusa_test.users").Iter()
	for iter.Scan(&email, &name) {
		users = append(users, User{Email: email, Name: name})
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}
