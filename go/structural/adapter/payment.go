package adapter

import (
	"fmt"
)

type PaymentStatus string

const (
	Success PaymentStatus = "SUCCESS"
	Failure PaymentStatus = "FAILURE"
)

type PaymentRequest struct {
	Name   string
	Phone  string
	Email  string
	Amount int
}

type PaymentProviderInterface interface {
	GenerateLink() string
	Pay(paymentRequest PaymentRequest)
	CheckStatus() PaymentStatus
}

type CashfreeApi struct{}

func (c *CashfreeApi) CreateURL() string {
	return "Cashfree"
}

func (c *CashfreeApi) DoPayment(amount int) {
	fmt.Printf("Cashfree Payment: %d\n", amount)
}

func (c *CashfreeApi) VerifyStatus() string {
	// Simulated status check
	return "OK"
}

type CashFreePayProvider struct {
	cashfreeApi *CashfreeApi
}

func NewCashFreePayProvider() *CashFreePayProvider {
	return &CashFreePayProvider{cashfreeApi: &CashfreeApi{}}
}

func (c *CashFreePayProvider) GenerateLink() string {
	return c.cashfreeApi.CreateURL()
}

func (c *CashFreePayProvider) Pay(paymentRequest PaymentRequest) {
	c.cashfreeApi.DoPayment(paymentRequest.Amount)
}

func (c *CashFreePayProvider) CheckStatus() PaymentStatus {
	status := c.cashfreeApi.VerifyStatus()
	return c.toPaymentStatus(status)
}

func (c *CashFreePayProvider) toPaymentStatus(status string) PaymentStatus {
	if status == "OK" {
		return Success
	}
	return Failure
}

type RazorPayApi struct{}

func (r *RazorPayApi) MakeLink() string {
	return "RazorPay"
}

func (r *RazorPayApi) PrePay() {
	fmt.Println("RazorPay PrePayment")
}

func (r *RazorPayApi) Pay(name string, amount int) {
	fmt.Printf("RazorPay Payment for %s of amount %d\n", name, amount)
}

func (r *RazorPayApi) CheckStatus() string {
	// Simulated status check
	return "PASS"
}

type RazorPayProvider struct {
	razorpayApi *RazorPayApi
}

func NewRazorPayProvider() *RazorPayProvider {
	return &RazorPayProvider{razorpayApi: &RazorPayApi{}}
}

func (r *RazorPayProvider) GenerateLink() string {
	return r.razorpayApi.MakeLink()
}

func (r *RazorPayProvider) Pay(paymentRequest PaymentRequest) {
	r.razorpayApi.PrePay()
	r.razorpayApi.Pay(paymentRequest.Name, paymentRequest.Amount)
}

func (r *RazorPayProvider) CheckStatus() PaymentStatus {
	status := r.razorpayApi.CheckStatus()
	return r.toPaymentStatus(status)
}

func (r *RazorPayProvider) toPaymentStatus(status string) PaymentStatus {
	if status == "PASS" {
		return Success
	}
	return Failure
}

func ProcessPayment(provider PaymentProviderInterface, paymentRequest PaymentRequest) {
	fmt.Println(provider.GenerateLink())
	provider.Pay(paymentRequest)
	status := provider.CheckStatus()
	fmt.Printf("Payment Status: %s\n", status)
}
