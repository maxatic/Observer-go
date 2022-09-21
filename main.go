package main

import "fmt"

type Subject interface {
	Subscribe()
	UnSubscribe()
	Send()
}

type ob interface {
	RecID() int
	HandleEvent(objects []string)
}

type Customer struct {
	id       int
	username string
}

func (customer *Customer) HandleEvent(news []string) {
	fmt.Println("Greetings, ", customer.username, "\nFreshly received News: ")
	for _, val := range news {
		fmt.Println(val)
	}
	fmt.Println("\n")
}

func (customer *Customer) RecID() int {
	return customer.id
}

type NewYorkTimes struct {
	subs    []ob
	objects []string
}

func (NYtimes *NewYorkTimes) addObject(object string) {
	NYtimes.objects = append(NYtimes.objects, object)
	NYtimes.Send()
}

func (NYtimes *NewYorkTimes) remObject(object string) {
	size := len(NYtimes.objects)
	for i, item := range NYtimes.objects {
		if item == object {
			NYtimes.objects[i] = NYtimes.objects[size-1]
		}
	}
	NYtimes.objects = NYtimes.objects[:size-1]
	NYtimes.Send()
}

func (NYtimes *NewYorkTimes) Subscribe(observer ob) {
	NYtimes.subs = append(NYtimes.subs, observer)
}

func (NYtimes *NewYorkTimes) UnSubscribe(observer ob) {
	size := len(NYtimes.subs)
	for i, item := range NYtimes.subs {
		if item.RecID() == observer.RecID() {
			NYtimes.subs[i] = NYtimes.subs[size-1]
		}
	}
	NYtimes.subs = NYtimes.subs[:size-1]
}

func (NYtimes *NewYorkTimes) Send() {
	for _, val := range NYtimes.subs {
		val.HandleEvent(NYtimes.objects)
	}

}

func main() {
	NYtimes := NewYorkTimes{objects: []string{
		"Mobilization Comes After Humiliating Losses in Ukraine",
		"Biden to Address U.N. on Day of Ukraine Debates",
	}}
	cust1 := &Customer{id: 1, username: "Volodymyr Zelensky"}
	cust2 := &Customer{id: 2, username: "vladimir putin"}
	NYtimes.Subscribe(cust1)
	NYtimes.addObject("Ukraine’s advances are exposing deep vulnerabilities in Russia’s military. Is it just the beginning? \n")
	NYtimes.Subscribe(cust2)
	NYtimes.addObject("Analysis: With his latest speech, President Vladimir Putin showed that he is at his most dangerous when he is cornered.\n")
	NYtimes.remObject("Biden to Address U.N. on Day of Ukraine Debates\n")
	NYtimes.UnSubscribe(cust1)
	NYtimes.addObject("Special Master Skeptical of Declassification Claims by Trump’s Lawyers\n")
}
