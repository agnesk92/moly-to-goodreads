package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToUniversalDate(t *testing.T) {
	date := FormatDate("2010. február  3., 20:24")
	require.Equal(t, "2010-02-03", date)

	date = FormatDate("2010. február 13., 20:24")
	require.Equal(t, "2010-02-13", date)

	date = FormatDate("2030. december 23., 20:24")
	require.Equal(t, "2030-12-23", date)
}

func TestToUniversalDateIncomplete(t *testing.T) {
	date := FormatDate("2010. december")
	require.Equal(t, "", date)
}

func TestToUniversalDateMonthNotFound(t *testing.T) {
	require.PanicsWithError(
		t,
		"Month not found: random",
		func() { FormatDate("2010. random 13., 20:24") },
	)
}

func TestCleanString(t *testing.T) {
	require.Equal(t, "2010", cleanString("2010.,"))
	require.Equal(t, "2010", cleanString("2010,"))
	require.Equal(t, "2010", cleanString("2010"))

	require.Equal(t, "", cleanString(""))
	require.Equal(t, "", cleanString("."))
	require.Equal(t, "", cleanString(".,"))
}
