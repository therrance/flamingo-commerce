package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"flamingo.me/flamingo-commerce/v3/cart/domain/cart"
)

// AllBehaviour is an autogenerated mock type for the AllBehaviour type
type AllBehaviour struct {
	mock.Mock
}

var (
	_ cart.ModifyBehaviour             = new(AllBehaviour)
	_ cart.CompleteBehaviour           = new(AllBehaviour)
	_ cart.GiftCardBehaviour           = new(AllBehaviour)
	_ cart.GiftCardAndVoucherBehaviour = new(AllBehaviour)
)

// AddToCart provides a mock function with given fields: ctx, _a1, deliveryCode, addRequest
func (_m *AllBehaviour) AddToCart(ctx context.Context, _a1 *cart.Cart, deliveryCode string, addRequest cart.AddRequest) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, addRequest)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.AddRequest) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, cart.AddRequest) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, cart.AddRequest) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ApplyVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *AllBehaviour) ApplyVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, couponCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CleanCart provides a mock function with given fields: ctx, _a1
func (_m *AllBehaviour) CleanCart(ctx context.Context, _a1 *cart.Cart) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart) *cart.Cart); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CleanDelivery provides a mock function with given fields: ctx, _a1, deliveryCode
func (_m *AllBehaviour) CleanDelivery(ctx context.Context, _a1 *cart.Cart, deliveryCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, deliveryCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteItem provides a mock function with given fields: ctx, _a1, itemID, deliveryCode
func (_m *AllBehaviour) DeleteItem(ctx context.Context, _a1 *cart.Cart, itemID string, deliveryCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemID, deliveryCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, string) error); ok {
		r2 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *AllBehaviour) RemoveVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, couponCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateAdditionalData provides a mock function with given fields: ctx, _a1, additionalData
func (_m *AllBehaviour) UpdateAdditionalData(ctx context.Context, _a1 *cart.Cart, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, additionalData)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateBillingAddress provides a mock function with given fields: ctx, _a1, billingAddress
func (_m *AllBehaviour) UpdateBillingAddress(ctx context.Context, _a1 *cart.Cart, billingAddress cart.Address) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, billingAddress)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.Address) *cart.Cart); ok {
		r0 = rf(ctx, _a1, billingAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.Address) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, billingAddress)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.Address) error); ok {
		r2 = rf(ctx, _a1, billingAddress)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateDeliveryInfo provides a mock function with given fields: ctx, _a1, deliveryCode, deliveryInfo
func (_m *AllBehaviour) UpdateDeliveryInfo(ctx context.Context, _a1 *cart.Cart, deliveryCode string, deliveryInfo cart.DeliveryInfoUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, deliveryInfo)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateDeliveryInfoAdditionalData provides a mock function with given fields: ctx, _a1, deliveryCode, additionalData
func (_m *AllBehaviour) UpdateDeliveryInfoAdditionalData(ctx context.Context, _a1 *cart.Cart, deliveryCode string, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, additionalData)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateItem provides a mock function with given fields: ctx, _a1, itemUpdateCommand
func (_m *AllBehaviour) UpdateItem(ctx context.Context, _a1 *cart.Cart, itemUpdateCommand cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemUpdateCommand)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateItems provides a mock function with given fields: ctx, _a1, itemUpdateCommands
func (_m *AllBehaviour) UpdateItems(ctx context.Context, _a1 *cart.Cart, itemUpdateCommands []cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemUpdateCommands)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePaymentSelection provides a mock function with given fields: ctx, _a1, paymentSelection
func (_m *AllBehaviour) UpdatePaymentSelection(ctx context.Context, _a1 *cart.Cart, paymentSelection cart.PaymentSelection) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, paymentSelection)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.PaymentSelection) *cart.Cart); ok {
		r0 = rf(ctx, _a1, paymentSelection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.PaymentSelection) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, paymentSelection)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.PaymentSelection) error); ok {
		r2 = rf(ctx, _a1, paymentSelection)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePurchaser provides a mock function with given fields: ctx, _a1, purchaser, additionalData
func (_m *AllBehaviour) UpdatePurchaser(ctx context.Context, _a1 *cart.Cart, purchaser *cart.Person, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, purchaser, additionalData)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Complete provides a mock function with given fields: _a0, _a1
func (_m *AllBehaviour) Complete(_a0 context.Context, _a1 cart.Cart) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, cart.Cart) *cart.Cart); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, cart.Cart) cart.DeferEvents); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, cart.Cart) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Restore provides a mock function with given fields: _a0, _a1
func (_m *AllBehaviour) Restore(_a0 context.Context, _a1 cart.Cart) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, cart.Cart) *cart.Cart); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, cart.Cart) cart.DeferEvents); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, cart.Cart) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ApplyGiftCard provides a mock function with given fields: ctx, _a1, giftCardCode
func (_m *AllBehaviour) ApplyGiftCard(ctx context.Context, _a1 *cart.Cart, giftCardCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, giftCardCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, giftCardCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, giftCardCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, giftCardCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveGiftCard provides a mock function with given fields: ctx, _a1, giftCardCode
func (_m *AllBehaviour) RemoveGiftCard(ctx context.Context, _a1 *cart.Cart, giftCardCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, giftCardCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, giftCardCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, giftCardCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, giftCardCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ApplyAny provides a mock function with given fields: ctx, _a1, anyCode
func (_m *AllBehaviour) ApplyAny(ctx context.Context, _a1 *cart.Cart, anyCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, anyCode)

	var r0 *cart.Cart
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, anyCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	var r1 cart.DeferEvents
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, anyCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, anyCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}