package factory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	// test cash
	payment, err := GetPaymentMethod(Cash)
	require.NoError(t, err, "A payment method of type 'Cash' must exist")
	msg := payment.Pay(10.30)
	require.Contains(t, msg, "paid using cash", "The cash payment method message wasn't correct")

	// test debit card
	payment, err = GetPaymentMethod(DebitCard)
	require.NoError(t, err, "A payment method of type 'DebitCard' must exist")
	msg = payment.Pay(22.30)
	require.Contains(t, msg, "paid using new debit card", "The debit card payment method message wasn't correct")

}
