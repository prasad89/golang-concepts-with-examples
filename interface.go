// https://go.dev/play/p/styxyJDeRsT

//go:build ignore

package main

import "fmt"

type SalaryCalculator interface {
	SalaryCalculator() int
}

type Permanent struct {
	EmpID int
	Name  string
	Basic int
	PF    int
}

type Intern struct {
	EmpID   int
	Name    string
	Stipend int
}

func (p Permanent) SalaryCalculator() int {
	return p.Basic + p.PF
}

func (i Intern) SalaryCalculator() int {
	return i.Stipend
}

func main() {
	p1 := Permanent{101, "Krishna", 5000, 900}
	i1 := Intern{201, "Amarja", 2700}

	fmt.Println("Permanent Employee Salary:", p1.SalaryCalculator())
	fmt.Println("Intern Salary:", i1.SalaryCalculator())
}
