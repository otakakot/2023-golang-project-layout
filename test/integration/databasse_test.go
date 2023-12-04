package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/otakakot/2023-golang-project-layout/internal/adapter/gateway"
	"github.com/otakakot/2023-golang-project-layout/internal/domain/model"
	"github.com/otakakot/2023-golang-project-layout/internal/driver/postgres"
)

func TestTodoGateway(t *testing.T) {
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

	gw := gateway.NewTodo(db)

	t.Run("保存して取得して削除して一覧取得する", func(t *testing.T) {
		t.Parallel()

		todo := model.GenerateTodo("title")

		if err := gw.Save(context.Background(), *todo); err != nil {
			t.Errorf("failed to create todo: %s", err)
		}

		got, err := gw.Find(context.Background(), todo.ID)
		if err != nil {
			t.Errorf("failed to find todo: %s", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreFields(model.Todo{}, "CreatedAt", "UpdatedAt"),
		}

		if diff := cmp.Diff(*todo, *got, opts...); diff != "" {
			t.Errorf("todo mismatch (-want +got):\n%s", diff)
		}

		if err := gw.Delete(context.Background(), todo.ID); err != nil {
			t.Errorf("failed to delete todo: %s", err)
		}

		gots, err := gw.List(context.Background())
		if err != nil {
			t.Errorf("failed to list todo: %s", err)
		}

		if len(gots.Todos()) != 0 {
			t.Errorf("todo count mismatch: want 0, got %d", len(gots.Todos()))
		}
	})
}
