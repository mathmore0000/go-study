package service_test

import (
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get total tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetAllCount = func() int {
			return 100
		}

		rp.FuncGetTicketsCountByDestinationCountry = func(country string) int {
			return 10
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total := sv.GetTicketsCountByDestinationCountry("Brazil")

		// assert
		expectedTotal := 10
		require.Equal(t, expectedTotal, total)
	})
	t.Run("success to get percentage of tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetAllCount = func() int {
			return 100
		}

		rp.FuncGetTicketsCountByDestinationCountry = func(country string) int {
			return 10
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		percentage, err := sv.GetPercentageTicketsByDestinationCountry("Brazil")

		// assert
		var expectedPercentage float32 = 10.0
		require.NoError(t, err, "should not return an error")
		require.Equal(t, expectedPercentage, percentage)
	})
	t.Run("fail to get percentage of tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetAllCount = func() int {
			return 100
		}

		rp.FuncGetTicketsCountByDestinationCountry = func(country string) int {
			return 0
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		percentage, err := sv.GetPercentageTicketsByDestinationCountry("Brazil")

		// assert
		var expectedPercentage float32 = 0
		require.Error(t, err, "should return an error")
		require.Equal(t, expectedPercentage, percentage)
	})
}
