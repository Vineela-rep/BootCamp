package main

import "fmt"

type employee interface {
	Fulltime()
	Contractor()
	Freelancer()
}
type Salary struct {
	BasicFT int
	BasicCT int
	BasicFL int
}
func (s1 Salary) Fulltime() {
	var MSalaryFT int = s1.BasicFT*30
	var YsalaryFT int = MSalaryFT*12
	fmt.Println("Monthly Salary of a Fulltime employee if: ",MSalaryFT)
	fmt.Println("Yearly Salary of a Fulltime employee if: ", YsalaryFT)
}
func (s2 Salary) Contractor() {
	var MSalaryCT int = s2.BasicCT*30
	var YsalaryCT int = MSalaryCT*12
	fmt.Println("Monthly Salary of a Contractor employee if: ",MSalaryCT)
	fmt.Println("Yearly Salary of a Contractor employee if: ", YsalaryCT)
}
func (s3 Salary) Freelancer() {
	var MSalaryFL int = s3.BasicFL*20
	var YsalaryFL int = MSalaryFL*12
	fmt.Println("Monthly Salary of a Freelancer employee if: ",MSalaryFL)
	fmt.Println("Yearly Salary of a Freelancer employee if: ", YsalaryFL)
}
func main(){
	var e employee= Salary{500,100,10}
	e.Fulltime()
	e.Contractor()
	e.Freelancer()
}