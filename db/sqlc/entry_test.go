// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: entry.sql

package db

import (
	"awesomeProject1/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreateEntry(t *testing.T) {
	randomAccount := createRandomAccount(t)
	entry := createRandomEntry(t, randomAccount)
	require.NotEmpty(t, entry)
}

func TestQueries_GetEntry(t *testing.T) {
	randomAccount := createRandomAccount(t)
	entry1 := createRandomEntry(t, randomAccount)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestQueries_ListEntries(t *testing.T) {
	randomAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, randomAccount)
	}

	arg := ListEntriesParams{
		AccountID: randomAccount.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}

}

func createRandomEntry(t *testing.T, account Account) Entry {

	params := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, entry.AccountID, params.AccountID)
	require.Equal(t, entry.Amount, params.Amount)

	return entry
}
