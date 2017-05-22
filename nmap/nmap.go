package nmap

import "fmt"

// Nmap a n leveled map for searching with multiple keys
type Nmap struct {
}

// Interface interface for nmap
type Interface interface {
	CreateMap(int) Nmap
	//this float array is ordered according to player order
	//the string stands for the action name.
	SearchMap([]string) []float64
}

func (n *Nmap) CreateMap(n int) Nmap {

}
