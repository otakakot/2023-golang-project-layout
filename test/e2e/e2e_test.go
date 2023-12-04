package e2e_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/otakakot/2023-golang-project-layout/internal/driver/postgres"
	"github.com/otakakot/2023-golang-project-layout/pkg/api"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	dsn := os.Getenv("POSTGRES_URL")

	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	db, err := postgres.New(dsn)
	if err != nil {
		t.Fatalf("failed to initialize postgres: %s", err)
	}

	t.Cleanup(func() {
		t.Helper()

		if _, err := db.ExecContext(context.Background(), "TRUNCATE TABLE todos"); err != nil {
			t.Errorf("failed to truncate table: %s", err)
		}

		if err := db.Close(); err != nil {
			t.Errorf("failed to close database: %s", err)
		}
	})

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:8080"
	}

	cli, err := api.NewClient(endpoint)
	if err != nil {
		t.Fatalf("failed to create client: %s", err)
	}

	t.Run("保存して更新して取得して削除して一覧取得する", func(t *testing.T) {
		t.Parallel()

		want, err := cli.CreateTodo(context.Background(), &api.CreateTodoRequest{
			Title: "test",
		})
		if err != nil {
			t.Fatalf("failed to create todo: %s", err)
		}

		got, err := cli.GetTodo(context.Background(), api.GetTodoParams{
			ID: want.ID.String(),
		})
		if err != nil {
			t.Fatalf("failed to get todo: %s", err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}

		if _, err := cli.UpdateTodo(context.Background(), &api.UpdateTodoRequest{
			Title:     "updated",
			Completed: true,
		}, api.UpdateTodoParams{
			ID: want.ID.String(),
		}); err != nil {
			t.Fatalf("failed to update todo: %s", err)
		}

		got, err = cli.GetTodo(context.Background(), api.GetTodoParams{
			ID: want.ID.String(),
		})
		if err != nil {
			t.Fatalf("failed to get todo: %s", err)
		}

		want.Completed = true

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}

		gots, err := cli.ListTodos(context.Background())
		if err != nil {
			t.Fatalf("failed to list todos: %s", err)
		}

		if len(gots) != 1 {
			t.Errorf("want: 1, got: %d", len(gots))
		}

		if err := cli.DeleteTodo(context.Background(), api.DeleteTodoParams{
			ID: want.ID.String(),
		}); err != nil {
			t.Fatalf("failed to delete todo: %s", err)
		}

		gots, err = cli.ListTodos(context.Background())
		if err != nil {
			t.Fatalf("failed to list todos: %s", err)
		}

		if len(gots) != 0 {
			t.Errorf("want: 0, got: %d", len(gots))
		}
	})
}
