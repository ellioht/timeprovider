package main

import (
	"fmt"
	"time"
)

type TimeProvider interface {
	Now() time.Time
}

type RealTimeProvider struct {}

type MockTimeProvider struct {
	MockTime time.Time
}

// Declare Now() method on RealTimeprovider struct 
func (r RealTimeProvider) Now() time.Time {
	return time.Now()
}

// Declare Now() method on MockTimeProvider struct
func (m MockTimeProvider) Now() time.Time {
	return m.MockTime
}

type Subsciption struct {
	billDate string
	amount  float64
}

type Customer struct {
	name string
	subsciption Subsciption
}

// BillingSystem sctruct that uses TimeProvider interface
type BillingSystem struct {
	timeProvider TimeProvider
}

func (bs *BillingSystem) CalculateBill(c *Customer) float64 {
	// Get the current time from the time provider
	now := bs.timeProvider.Now()

	// Parse the bill date from the customer's subscription
	billDate, err := time.Parse("2006-01-02", c.subsciption.billDate)
	if err != nil {
			// Handle error
	}

	// Calculate the number of months between the bill date and the current time
	months := int(now.Sub(billDate).Hours() / 24 / 30)

	// Calculate the total bill amount
	total := c.subsciption.amount * float64(months)

	return total
}

func main() {
	mytime := time.Now()
	timeformat := mytime.Format("2006-01-02 15:04:05") 
	fmt.Println("Formatted time:", timeformat)

	rtp := RealTimeProvider{}
	rtpTime := rtp.Now()
	rtpFormat := rtpTime.Format("2006-01-02 15:04:05")
	fmt.Println("Using interface", rtpFormat)

	mtp := MockTimeProvider{MockTime: time.Date(2021, 10, 10, 0, 0, 0, 0, time.UTC)}
	mtpTime := mtp.Now()
	mtpFormat := mtpTime.Format("2006-01-02 15:04:05")
	fmt.Println("Mock time", mtpFormat)

	
	// Create a new billing system that uses the real time provider
	bs := BillingSystem{timeProvider: RealTimeProvider{}}

	// Create a new customer with a subscription that started on 2020-01-01
	c := Customer{
		name: "John Smith",
		subsciption: Subsciption{
			billDate: "2020-01-01",
			amount: 9.99,
		},
	}
	
	// Calculate the bill for the customer
	bill := bs.CalculateBill(&c)
	fmt.Println("Bill for", c.name, "is", bill)
	
}