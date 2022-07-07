package main

import "fmt"
import "errors"
import "encoding/json"

type Matrix struct {
	rows int
	columns int
	elements [][] int
}

func (m Matrix) GetRows() int {
	return m.rows
}

func (m Matrix) GetColoumns() int {
	return m.columns
}

func (m Matrix) SetElements(i,j,num int) (error) {
	if i<0 || j<0 || i>=m.GetRows() || j>=m.GetColoumns(){
		return errors.New("out of bounds")
	}
	m.elements[i][j]=num
	return nil
}

func (m Matrix) Add(mat2 [][] int)([][] int){
	var i,j int
	for i=0;i<m.GetRows();i++ {
		for j = 0; j < m.GetColoumns(); j++ {
			mat2[i][j] = m.elements[i][j] + mat2[i][j];
		}
	}
	return mat2
}

func main(){
	m:= Matrix{rows:2,columns:2,elements:[][]int{{1,2},{3,4}}}

	fmt.Println(m.GetRows())
	fmt.Println(m.GetColoumns())
	fmt.Println((m.SetElements(1,1,25)))

	var AddMatrix=[][]int{{1,1}, {1,1}}
	AddMatrix=m.Add(AddMatrix)
	fmt.Println(AddMatrix)

}