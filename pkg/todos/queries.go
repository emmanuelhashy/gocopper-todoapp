package todos

import (
	"context"
	"database/sql"

	"github.com/gocopper/copper/csql"
)

var ErrRecordNotFound = sql.ErrNoRows

func NewQueries(querier csql.Querier) *Queries {
	return &Queries{
		querier: querier,
	}
}

type Queries struct {
	querier csql.Querier
}

/*
Here are some example queries that use Querier to unmarshal results into Go strcuts

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	const query = "SELECT * FROM posts ORDER BY created_at DESC"

	var (
	    posts []Post
	    err = q.querier.Select(ctx, &posts, query)
    )

	return posts, err
}

func (q *Queries) GetPostByID(ctx context.Context, id string) (*Post, error) {
	const query = "SELECT * from posts where id=?"

	var (
	    post Post
	    err = q.querier.Get(ctx, &post, query, id)
    )

	return &post, err
}

func (q *Queries) SavePost(ctx context.Context, post *Post) error {
	const query = `
	INSERT INTO posts (id, title, url, poster)
	VALUES (?, ?, ?, ?)
	ON CONFLICT (id) DO UPDATE SET title=?, url=?`

	_, err := q.querier.Exec(ctx, query,
		post.ID,
		post.Title,
		post.URL,
		post.Poster,
		post.Title,
		post.URL,
	)

	return err
}
*/

func (q *Queries) SaveTodo(ctx context.Context, todo *Todo) error {
	const query = `
	INSERT INTO todos (name)
	VALUES (?)`

	_, err := q.querier.Exec(ctx, query,
		todo.Name,
	)

	return err
}


func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
    const query = "SELECT * FROM todos"
    var (
        todos []Todo
        err   = q.querier.Select(ctx, &todos, query)
    )
    return todos, err
}

func (q *Queries) UpdateTodo(ctx context.Context, oldName string, todo *Todo) error {
    const query = `
    UPDATE todos SET name=(?) WHERE name=(?)`
    _, err := q.querier.Exec(ctx, query,
        todo.Name,
        oldName,
    )
    return err
}

func (q *Queries) DeleteTodo(ctx context.Context, todo *Todo) error {
	const query = `
	DELETE from todos WHERE name=(?)`

	_, err := q.querier.Exec(ctx, query,
		todo.Name,
	)

	return err
}