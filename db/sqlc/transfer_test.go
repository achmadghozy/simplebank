package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/achmadghozy/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfers {
	user := CreateRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: "USD",
	}

	account1, _ := testQueries.CreateAccount(context.Background(), arg)
	account2, _ := testQueries.CreateAccount(context.Background(), arg)

	arg1 := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg1)

	require.NoError(t, err)
	require.Equal(t, arg1.Amount, transfer.Amount)

	return transfer
}

func TestCreateRandomTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetAccount(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfer(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2 := CreateRandomUser(t)

	arg1 := CreateAccountParams{
		Owner:    user1.Username,
		Balance:  300,
		Currency: util.USD,
	}

	arg2 := CreateAccountParams{
		Owner:    user2.Username,
		Balance:  250,
		Currency: util.USD,
	}

	account1, _ := testQueries.CreateAccount(context.Background(), arg1)
	account2, _ := testQueries.CreateAccount(context.Background(), arg2)

	arg3 := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        10,
	}

	for i := 0; i < 5; i++ {
		_, _ = testQueries.CreateTransfer(context.Background(), arg3)
	}

	arg4 := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg4)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Equal(t, len(transfers), 5)
}
