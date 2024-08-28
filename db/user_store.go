package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/sanLimbu/horroscope/types"
)

const userTable = "users"

type Map map[string]any

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

func (s *PostgresUserStore) Drop(ctx context.Context) error {
	fmt.Println("-------dropping user table-------------")
	_, err := s.db.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", userTable))
	return err
}

func (s *PostgresUserStore) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1, email=$s, password=$3 where id=$4", userTable)
	_, err := s.db.ExecContext(ctx, query, params.FirstName, params.LastName, params.Email, params.Password, filter["id"])
	return err

}

func (s *PostgresUserStore) DeleteUser(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}

func (s *PostgresUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	user.ID = uuid.New() //Generate a new UUID for the USER ID
	query := fmt.Sprintf(`INSERT INTO %s (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING id`, userTable)
	err := s.db.QueryRowContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.EncryptedPassword).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User

	query := fmt.Sprintf("SELECT id, name, email, password FROM %s WHERE email=$1", userTable)
	err := s.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	query := fmt.Sprintf("SELECT id, name, email, password FROM %s WHERE id=$1", userTable)
	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.Email, &user.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	query := fmt.Sprintf("SELECT id, firstName, email, password FROM %s", userTable)
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*types.User
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.Email, &user.EncryptedPassword); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/sanLimbu/horroscope/types"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// const userColl = "users"

// type Map map[string]any

// type Dropper interface {
// 	Drop(context.Context) error
// }

// type UserStore interface {
// 	Dropper
// 	GetUserByEmail(context.Context, string) (*types.User, error)
// 	GetUserByID(context.Context, string) (*types.User, error)
// 	GetUsers(context.Context) ([]*types.User, error)
// 	InsertUser(context.Context, *types.User) (*types.User, error)
// 	DeleteUser(context.Context, string) error
// 	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
// }

// type MongoUserStore struct {
// 	client *mongo.Client
// 	coll   *mongo.Collection
// }

// func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
// 	dbname := os.Getenv(MongoDBNameEnvName)
// 	return &MongoUserStore{
// 		client: client,
// 		coll:   client.Database(dbname).Collection(userColl),
// 	}
// }

// func (s *MongoUserStore) Drop(ctx context.Context) error {
// 	fmt.Println("---- dropping user collection ----")
// 	return s.coll.Drop(ctx)
// }

// func (s *MongoUserStore) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
// 	oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
// 	if err != nil {
// 		return err
// 	}

// 	filter["_id"] = oid
// 	update := bson.M{"$set": params.ToBSON()}
// 	_, err = s.coll.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
// 	res, err := s.coll.InsertOne(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user.ID = res.InsertedID.(primitive.ObjectID)
// 	return user, nil
// }

// func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
// 	var user types.User
// 	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var user types.User
// 	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
// 	curr, err := s.coll.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []*types.User
// 	if err := curr.All(ctx, &users); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
