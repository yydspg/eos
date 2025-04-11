package mongoutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eos/tools/db/tx"
	"github.com/eos/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMaxPoolSize = 100
	defaultMaxRetry    = 3
)

type Config struct {
	Uri         string
	Address     []string
	Database    string
	Username    string
	Password    string
	AuthSource  string
	MaxPoolSize int
	MaxRetry    int
}
type Client struct {
	tx tx.Tx
	db *mongo.Database
}

func (c *Client) GetDB() *mongo.Database {
	return c.db
}

func (c *Client) GetTx() tx.Tx {
	return c.tx
}

func (c *Config) ValidateAndSetDefaults() error {
	if c.Uri == "" && len(c.Address) == 0 {
		return errs.Wrap(errs.New("either uri or address must be set"))
	}
	if c.Database == "" {
		return errs.Wrap(errs.New("database is requried"))
	}
	if c.MaxPoolSize <= 0 {
		c.MaxRetry = defaultMaxPoolSize
	}
	if c.MaxRetry <= 0 {
		c.MaxRetry = defaultMaxRetry
	}
	if c.Uri == "" {
		if c.AuthSource == "" {
			c.Uri = buildMongoURI(c, c.Database)
		} else {
			c.Uri = buildMongoURI(c, c.AuthSource)
		}
	}
	return nil
}

func NewMongoDB(ctx context.Context, config *Config) (*Client, error) {
	if err := config.ValidateAndSetDefaults(); err != nil {
		return nil, err
	}
	opts := options.Client().ApplyURI(config.Uri).SetMaxPoolSize(uint64(config.MaxPoolSize))
	var (
		cli *mongo.Client
		err error
	)
	for i := 0; i < config.MaxRetry; i++ {
		cli, err = connectMongo(ctx, opts)
		if err != nil && shouldRetry(ctx, err) {
			time.Sleep(time.Second / 2)
			continue
		}
		break
	}
	if err != nil {
		return nil, errs.WrapMsg(err, "failed connect to MongoDB", "URI", config.Uri)
	}
	mtx, err := NewMongoTx(ctx, cli)
	if err != nil {
		return nil, err
	}
	return &Client{
		tx: mtx,
		db: cli.Database(config.Database),
	}, nil
}
func shouldRetry(ctx context.Context, err error) bool {
	select {

	case <-ctx.Done():
		return false
	default:
		if cmdErr, ok := err.(mongo.CommandError); ok {
			return cmdErr.Code != 13 && cmdErr.Code != 18
		}
		return true
	}
}
func connectMongo(ctx context.Context, opts *options.ClientOptions) (*mongo.Client, error) {
	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return cli, nil
}

func NewMongoTx(ctx context.Context, client *mongo.Client) (tx.Tx, error) {
	mtx := &mongoTx{
		client: client,
	}
	if err := mtx.init(ctx); err != nil {
		return nil, err
	}
	return mtx, nil
}

type mongoTx struct {
	client *mongo.Client
	tx     func(context.Context, func(ctx context.Context) error) error
}

func (t *mongoTx) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if t.tx == nil {
		return fn(ctx)
	}
	return t.tx(ctx, fn)
}

func (t *mongoTx) init(ctx context.Context) error {
	var res map[string]any
	if err := t.client.Database("admin").RunCommand(ctx, bson.M{"isMaster": 1}).Decode(&res); err != nil {
		return errs.WrapMsg(err, "check wether mongo was deployed on cluster mode")
	}
	if _, allowTx := res["setName"]; !allowTx {
		return nil
	}
	t.tx = func(tx_ctx context.Context, fn func(ctx context.Context) error) error {
		session, err := t.client.StartSession()
		if err != nil {
			return errs.WrapMsg(err, "mongodb start sessino error")
		}
		defer session.EndSession(tx_ctx)
		_, err = session.WithTransaction(tx_ctx, func(session_ctx mongo.SessionContext) (any, error) {
			return nil, fn(session_ctx)
		})
		return errs.WrapMsg(err, "mongodb transaction failed")
	}
	return nil
}
func Check(ctx context.Context, config *Config) error {
	if err := config.ValidateAndSetDefaults(); err != nil {
		return err
	}
	clientOpts := options.Client().ApplyURI(config.Uri)
	mongoClient, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return errs.WrapMsg(err, "MongoDB connect failed", "URI", config.Uri, "Database", config.Database, "MaxPoolSize", config.MaxPoolSize)
	}

	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			_ = mongoClient.Disconnect(ctx)
		}
	}()

	if err = mongoClient.Ping(ctx, nil); err != nil {
		return errs.WrapMsg(err, "MongoDB ping failed", "URI", config.Uri, "Database", config.Database, "MaxPoolSize", config.MaxPoolSize)
	}
	return nil
}

func buildMongoURI(config *Config, authSource string) string {
	credentials := ""

	if config.Username != "" && config.Password != "" {
		credentials = fmt.Sprintf("%s:%s", config.Username, config.Password)
	}
	return fmt.Sprintf(
		"mongodb://%s@%s/%s?authSource=%s&maxPoolSize=%d,",
		credentials,
		strings.Join(config.Address, ","),
		config.Database,
		authSource,
		config.MaxPoolSize,
	)
}
// func init() {
// 	if err := specialerror.AddReplace(mongo.ErrNoDocuments, errs.ErrRecordNotFound); err != nil {
// 		panic(err)
// 	}
// }
