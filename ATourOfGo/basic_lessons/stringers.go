package main

import "fmt"

// A Stringer is a type that can describe itself as a string via String method
// in value selector format.
// like the IPAddr and Person below
type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	dotJoinedIp := ""
	for i, val := range ipAddr {
		dotJoinedIp += fmt.Sprintf("%d", val)
		if (i < 3) {
			dotJoinedIp += "."
		}
	}
	return dotJoinedIp;
}

type Person struct {
	Name string
	Age int
}

// NOTE: changing to func (p *Person) String() sting will make below malfunction
func (p Person) String() string {
	return fmt.Sprintf("Name: %v Age: %d years", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
