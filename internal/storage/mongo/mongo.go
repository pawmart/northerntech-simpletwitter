package mongo

import (
	"log"
	"strings"
	"github.com/pawmart/northerntech-simpletwitter/internal/models"
	"github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
	"gopkg.in/mgo.v2/bson"
	"github.com/pawmart/northerntech-simpletwitter/config"
	"gopkg.in/mgo.v2"
	"time"
)

func NewStorage(dbConfig *config.DbConfig) *Storage {
	return &Storage{config: dbConfig}
}
// Storage struct.
type Storage struct {
	config *config.DbConfig
	db     *mgo.Database
}

const collectionName = "tweets"

// InsertTweet handling.
func (s *Storage) InsertTweet(entity *models.Tweet) error {
	return s.getDB().C(collectionName).Insert(entity)
}

// FindTweet handling.
func (s *Storage) FindTweet(id string) (entity *models.Tweet, err error) {
	entity = new(models.Tweet)
	err = s.getDB().C(collectionName).Find(bson.M{"id": id}).One(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// UpdateTweet handling.
func (s *Storage) UpdateTweet(id string, p *models.Tweet) error {
	return s.getDB().C(collectionName).Update(bson.M{"id": id}, p)
}

// RemoveTweet handling.
func (s *Storage) RemoveTweet(id string) error {
	return s.getDB().C(collectionName).Remove(bson.M{"id": id})
}

// FindTweets handling.
func (s *Storage) FindTweets(params operations.GetTweetsParams) []*models.Tweet {
	var q interface{}
	if len(params.FilterTag) > 0 {
		var tagQueries []bson.M
		for _, v := range params.FilterTag {
			tagQueries = append(tagQueries, bson.M{"tag": v})
		}

		q = bson.D{{"$and", tagQueries}}
	}

	var desiredCount int
	desiredCount = 0
	if params.Count != nil {
		desiredCount = int(*params.Count)
	}

	var result []*models.Tweet
	s.getDB().C(collectionName).Find(q).All(&result)



	// NOTE: as mongo is lame handle filtering yourself for now...
	var filteredResults []*models.Tweet
	collected := 1
	for _, item := range result {

		// filter count limit
		if desiredCount != 0 && collected > desiredCount {
			break
		}

		// filter by year
		if params.Year != nil {
		createdOn := time.Unix(*item.CreatedOn, 0)
			if createdOn.Year() != int(*params.Year) {
				continue
			}
		}

		filteredResults = append(filteredResults, item)
		collected++
	}

	// TODO: sort out sort ;)
	//sort.SliceStable(filteredResults, func(i, j int) bool {
	//	return filteredResults[i].CreatedOn < filteredResults[j].CreatedOn
	//})

	return filteredResults
}

// Ping database.
func (s *Storage) Ping() error {
	return s.getDB().Session.Ping()
}

// Drop database.
func (s *Storage) Drop() {
	s.getDB().DropDatabase()
}

func (s *Storage) getDB() *mgo.Database {
	if s.db != nil {
		return s.db
	}

	conf := s.config
	var session *mgo.Session

	session, err := s.dial(s.config)
	if err != nil {
		log.Fatalln(err)
	}

	if conf.User != "" && conf.Password != "" {
		err := session.Login(&mgo.Credential{Username: conf.User, Password: conf.Password, Source: conf.Auth})
		if err != nil {
			log.Fatalln(err)
		}
	}

	session.SetSafe(&mgo.Safe{})

	s.db = session.DB(conf.Database)
	return s.db
}

func (s *Storage) dial(conf *config.DbConfig) (*mgo.Session, error) {
	// TODO: Enforce secure connection for production.
	host := s.formatHost(conf)
	sess, err := mgo.Dial(host)
	if err != nil {
		log.Print("Connection to mongo error: " + err.Error())
		return nil, err
	}
	cloned := sess.Clone()
	defer sess.Close()

	return cloned, err
}

func (s *Storage) formatHost(conf *config.DbConfig) string {
	h := strings.TrimSpace(conf.Host)
	h = strings.TrimPrefix(h, "mongodb://")
	return h
}
