package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	return db, mock
}

func TestCreateQuestion_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	text := "Test question"

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "questions" ("text","created_at") VALUES ($1,$2) RETURNING "id"`)).
		WithArgs(text, sqlmock.AnyArg()). // Только 2 аргумента!
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	question, err := repo.CreateQuestion(ctx, text)

	require.NoError(t, err)
	require.NotNil(t, question)
	assert.Equal(t, uint(1), question.ID)
	assert.Equal(t, text, question.Text)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetQuestions_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "text", "created_at"}).
		AddRow(1, "Question 1", now).
		AddRow(2, "Question 2", now)

	mock.ExpectQuery(`SELECT \* FROM "questions"`).
		WillReturnRows(rows)

	questions, err := repo.GetQuestions(ctx)

	require.NoError(t, err)
	assert.Len(t, questions, 2)
	assert.Equal(t, "Question 1", questions[0].Text)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetQuestionByID_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	questionID := uint(1)
	now := time.Now()
	userID := uuid.New()

	questionRows := sqlmock.NewRows([]string{"id", "text", "created_at"}).
		AddRow(questionID, "Test question", now)
	mock.ExpectQuery(`SELECT \* FROM "questions" WHERE "questions"\."id" = \$1 ORDER BY "questions"\."id" LIMIT \$2`).
		WithArgs(questionID, 1).
		WillReturnRows(questionRows)

	answerRows := sqlmock.NewRows([]string{"id", "question_id", "user_id", "text", "created_at"}).
		AddRow(1, questionID, userID, "Answer 1", now)
	mock.ExpectQuery(`SELECT \* FROM "answers" WHERE "answers"\."question_id" = \$1`).
		WithArgs(questionID).
		WillReturnRows(answerRows)

	question, err := repo.GetQuestionByID(ctx, questionID)

	require.NoError(t, err)
	require.NotNil(t, question)
	assert.Equal(t, questionID, question.ID)
	assert.Len(t, question.Answers, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteQuestion_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	questionID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "questions" WHERE`)).
		WithArgs(questionID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.DeleteQuestion(ctx, questionID)

	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateAnswer_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	questionID := uint(1)
	userID := uuid.New()
	text := "Test answer"

	countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(`SELECT count\(\*\) FROM "questions"`).
		WithArgs(questionID).
		WillReturnRows(countRows)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "answers" ("question_id","user_id","text","created_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(questionID, userID, text, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	answer, err := repo.CreateAnswer(ctx, questionID, userID, text)

	require.NoError(t, err)
	require.NotNil(t, answer)
	assert.Equal(t, uint(1), answer.ID)
	assert.Equal(t, text, answer.Text)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAnswerByID_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	answerID := uint(1)
	userID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "question_id", "user_id", "text", "created_at"}).
		AddRow(answerID, 1, userID, "Test answer", now)

	mock.ExpectQuery(`SELECT \* FROM "answers" WHERE "answers"\."id" = \$1 ORDER BY "answers"\."id" LIMIT \$2`).
		WithArgs(answerID, 1).
		WillReturnRows(rows)

	answer, err := repo.GetAnswerByID(ctx, answerID)

	require.NoError(t, err)
	require.NotNil(t, answer)
	assert.Equal(t, answerID, answer.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAnswer_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	answerID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "answers" WHERE`)).
		WithArgs(answerID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.DeleteAnswer(ctx, answerID)

	require.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
