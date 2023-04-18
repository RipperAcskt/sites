package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"sites/config"
	"sites/internal/model"
	"sites/internal/service"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB      *sql.DB
	Migrate *migrate.Migrate
	cfg     *config.Config
}

func New(cfg *config.Config) (*Postgres, error) {
	DB, err := sql.Open("pgx", cfg.GetPostgresUrl())
	if err != nil {
		return nil, fmt.Errorf("open failed: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("with instance failed: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(cfg.MIGRATE_PATH, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("new with database instance failed: %w", err)
	}

	return &Postgres{
		DB,
		m,
		cfg,
	}, nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}

func (p *Postgres) CreateUser(ctx context.Context, user model.User) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var name string
	err := p.DB.QueryRowContext(queryCtx, "SELECT name FROM users WHERE email = $1", user.Email).Scan(&name)
	if err == nil {
		return fmt.Errorf("user: %v: %w", user.Name, service.ErrUserAlreadyExists)

	}
	fmt.Println(user.Password)
	_, err = p.DB.ExecContext(ctx, "INSERT INTO users (name, email, password, balance) VALUES($1, $2, $3, $4)", user.Name, user.Email, []byte(user.Password), 0)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}

func (p *Postgres) CheckUserByPhoneNumber(ctx context.Context, email string) (*model.User, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := p.DB.QueryRowContext(queryCtx, "SELECT id, password FROM users WHERE email = $1", email)

	var user model.User

	err := row.Scan(&user.ID, &user.Password)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, service.ErrUserDoesNotExists
		}

		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return &user, nil
}

func (p *Postgres) GetUserById(ctx context.Context, id string) (*model.User, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := &model.User{}
	err := p.DB.QueryRowContext(queryCtx, "SELECT id, name, email, balance FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, service.ErrUserDoesNotExists
		}
		return nil, fmt.Errorf("query row context failed: %w", err)
	}

	return user, err
}

func (p *Postgres) CreatePayment(ctx context.Context, order model.Order) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.DB.ExecContext(queryCtx, "INSERT INTO orders (user_id, order_id, state, balance) VALUES($1, $2, $3, $4)", order.UserId, order.Data.OrderId, order.State, order.Data.Pricing.Local.Amount)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}

func (p *Postgres) GetPaymentByUserId(ctx context.Context, id string) (*model.Order, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var order model.Order
	err := p.DB.QueryRowContext(queryCtx, "SELECT * FROM orders WHERE user_id = $1", id).Scan(&order.ID, order.UserId, order.State, order.Data.Pricing.Local.Amount)
	if err != nil {
		return nil, fmt.Errorf("query row context failed: %w", err)
	}
	return &order, nil
}

func (p *Postgres) GetPaymentId(ctx context.Context, id string) (string, float64, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var order_id string
	var balance float64
	fmt.Println(id)
	err := p.DB.QueryRowContext(queryCtx, "SELECT order_id, balance FROM orders WHERE user_id = $1", id).Scan(&order_id, &balance)
	if err != nil {
		return "", 0, fmt.Errorf("query row context failed: %w", err)
	}
	return order_id, balance, nil
}

func (p *Postgres) CompletePayment(ctx context.Context, order_id, user_id string, balance float64) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.DB.ExecContext(queryCtx, "UPDATE orders SET state = $1 WHERE order_id = $2 AND state = $3", model.Completed, order_id, model.Submitted)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	_, err = p.DB.ExecContext(queryCtx, "UPDATE users SET balance = $1 WHERE user_id = $2", balance, user_id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}

func (p *Postgres) SetState(ctx context.Context, order_id string) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.DB.ExecContext(queryCtx, "UPDATE orders SET state = $1 WHERE order_id = $2", model.Expired, order_id)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}

func (p *Postgres) GetBalance(ctx context.Context, userId string) (float64, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var balance float64

	err := p.DB.QueryRowContext(queryCtx, "SELECT balance FROM users WHERE id = $1", userId).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("exec failed: %w", err)
	}
	return balance, nil
}

func (p *Postgres) SetBalance(ctx context.Context, userId string, balance float64) error {
	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.DB.ExecContext(queryCtx, "UPDATE users SET balance = $1 WHERE id = $2", balance, userId)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}
