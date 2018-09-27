package event

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

//nolint
func TestEventCollector_Post(t *testing.T) {
	sub1 := NewSubscription(100, []Event_Topic{Event_TransactionResult})
	sub2 := NewSubscription(100, []Event_Topic{Event_ContractEvent})
	sub3 := NewSubscription(100, []Event_Topic{Event_TransactionResult, Event_ContractEvent})

	ec := GetEventCollectorInstance()
	ec.Subscribe(sub1)
	ec.Subscribe(sub2)
	ec.Subscribe(sub3)
	count1 := 0
	count2 := 0
	count3 := 0

	var mux sync.Mutex
	go func(ch <-chan *Event) {
		t.Log("run sub1")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_TransactionResult {
					t.Fatalf("sub1 expect event topic Event_TransactionResult, got %s", e.Topic.String())
				}
				if e.Data != "test1" {
					t.Fatalf("sub1 expect event data test1, got %s", e.Data)
				}
				mux.Lock()
				count1++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub1.ReadChan())

	go func(ch <-chan *Event) {
		t.Log("run sub2")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_ContractEvent {
					t.Fatalf("sub2 expect event topic Event_ContractEvent, got %s", e.Topic.String())
				}
				if e.Data != "test2" {
					t.Fatalf("sub2 expect event data test2, got %s", e.Data)
				}
				mux.Lock()
				count2++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub2.ReadChan())

	go func(ch <-chan *Event) {
		t.Log("run sub3")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_TransactionResult && e.Topic != Event_ContractEvent {
					t.Fatalf("sub3 expect event topic Event_TransactionResult or Event_ContractEvent, got %s", e.Topic.String())
				}
				mux.Lock()
				count3++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub3.ReadChan())

	ec.Post(NewEvent(Event_TransactionResult, "test1"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))

	time.Sleep(time.Millisecond * 100)

	mux.Lock()
	if count1 != 1 || count2 != 2 || count3 != 3 {
		t.Fatalf("expect count1 = 1, count2 = 2, count3 = 3, got %d %d %d", count1, count2, count3)
	}
	mux.Unlock()

	ec.Unsubscribe(sub1)
	ec.Post(NewEvent(Event_TransactionResult, "test1"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))

	time.Sleep(time.Millisecond * 100)

	mux.Lock()
	if count1 != 1 || count2 != 4 || count3 != 6 {
		t.Fatalf("expect count1 = 1, count2 = 4, count3 = 6, got %d %d %d", count1, count2, count3)
	}
	mux.Unlock()
}

//nolint
func TestEventCollector_Full(t *testing.T) {
	sub1 := NewSubscription(1, []Event_Topic{Event_TransactionResult})
	sub2 := NewSubscription(1, []Event_Topic{Event_ContractEvent})
	sub3 := NewSubscription(1, []Event_Topic{Event_TransactionResult, Event_ContractEvent})

	ec := GetEventCollectorInstance()
	ec.Subscribe(sub1)
	ec.Subscribe(sub2)
	ec.Subscribe(sub3)
	count1 := 0
	count2 := 0
	count3 := 0

	ec.Post(NewEvent(Event_TransactionResult, "test1"))
	ec.Post(NewEvent(Event_TransactionResult, "test1"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))
	ec.Post(NewEvent(Event_ContractEvent, "test2"))
	time.Sleep(time.Millisecond * 100)

	var mux sync.Mutex
	go func(ch <-chan *Event) {
		t.Log("run sub1")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_TransactionResult {
					t.Fatalf("sub1 expect event topic Event_TransactionResult, got %s", e.Topic.String())
				}
				mux.Lock()
				count1++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub1.ReadChan())

	go func(ch <-chan *Event) {
		t.Log("run sub2")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_ContractEvent {
					t.Fatalf("sub2 expect event topic Event_ContractEvent, got %s", e.Topic.String())
				}
				mux.Lock()
				count2++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub2.ReadChan())

	go func(ch <-chan *Event) {
		t.Log("run sub3")
		loop:
		for {
			select {
			case e := <-ch:
				if e.Topic != Event_TransactionResult && e.Topic != Event_ContractEvent {
					t.Fatalf("sub3 expect event topic Event_TransactionResult or Event_ContractEvent, got %s", e.Topic.String())
				}
				mux.Lock()
				count3++
				mux.Unlock()
			case <- time.After(1 * time.Second):
				break loop
			}
		}
	}(sub3.ReadChan())

	time.Sleep(time.Millisecond * 100)
	mux.Lock()
	if count1 != 1 || count2 != 1 || count3 != 1 {
		t.Fatalf("expect count1 = 1, count2 = 1, count3 = 1, got %d %d %d", count1, count2, count3)
	}
	mux.Unlock()

	sub1 = NewSubscription(1000, []Event_Topic{Event_TransactionResult})
	sub2 = NewSubscription(1000, []Event_Topic{Event_ContractEvent})
	sub3 = NewSubscription(1000, []Event_Topic{Event_TransactionResult, Event_ContractEvent})
	ec.Subscribe(sub1)
	ec.Subscribe(sub2)
	ec.Subscribe(sub3)
	data := "test1"
	for i := 0; i < 5; i++ {
		data += data
	}

	t0 := time.Now().Nanosecond()
	// almost 6ms for 10000 post
	for i := 0; i < 1000; i++ {
		ec.Post(NewEvent(Event_TransactionResult, data))
		time.Sleep(time.Microsecond * 50)
	}
	t1 := time.Now().Nanosecond()
	fmt.Println(t1 - t0)
	time.Sleep(time.Millisecond * 100)
	mux.Lock()
	if count1 > 1001 || count2 != 1 || count3 > 1001 {
		t.Fatalf("expect count1 <= 1001, count2 = 1, count3 <= 1001, got %d %d %d", count1, count2, count3)
	}
	mux.Unlock()
}
