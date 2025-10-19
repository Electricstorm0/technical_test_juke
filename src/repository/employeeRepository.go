package repository

import (
	"database/sql"
	"technical_test/src/config"
	"technical_test/src/model"
)

func Create(emp *model.Employee) (int64,error) {
	result,err := config.DB.Exec("INSERT INTO employee(name,email,position,salary) VALUES(?,?,?,?)",emp.Name,emp.Email,emp.Position,emp.Salary)
	if err != nil{
		return 0, err
	}
	id,err := result.LastInsertId()
	if err != nil{
		return 0, err
	}
	return id,nil
}

func GetByEmail(email string)(*model.Employee,error)  {
	var emp model.Employee
	rows:=config.DB.QueryRow("SELECT id,name,email,position,salary FROM employee Where email=?",email)
	if err:=rows.Scan(&emp.ID,&emp.Name,&emp.Email,&emp.Position,&emp.Salary);
	err !=nil{
	if err ==sql.ErrNoRows{
		return nil,nil
	}
	return nil,err
	}
	return &emp,nil
}

func GetAllRepo() ([]model.Employee,error)  {
	var employees []model.Employee
	rows,err:= config.DB.Query("SELECT id,name,email,position,salary FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
	var employee model.Employee
	if err:= rows.Scan(&employee.ID,&employee.Name,&employee.Email,&employee.Position,&employee.Salary)
	 err != nil {
		return nil, err
	}
	employees = append(employees, employee)
	}
	return  employees,nil
}

func GetById(id int64) (*model.Employee,error)  {
	rows:= config.DB.QueryRow("SELECT id,name,email,position,salary FROM employee Where id=?",id)
	var emp model.Employee
	if err:= rows.Scan(&emp.ID,&emp.Name,&emp.Email,&emp.Position,&emp.Salary); err != nil {
		if err == sql.ErrNoRows  {
		return nil, nil
	}
		return nil, err	
	}
	return &emp,nil
}


func Update(id int64, emp model.Employee) (*model.Employee, error) {
    _, err := config.DB.Exec(
        "UPDATE employee SET name=?, email=?, position=?, salary=? WHERE id=?",
        emp.Name, emp.Email, emp.Position, emp.Salary, id,
    )
    if err != nil {
        return nil, err
    }

    updatedEmp := emp
    updatedEmp.ID = id

    return &updatedEmp, nil
}

func Delete(id int64) (*model.Employee,error)  {
	result,err:=config.DB.Exec("DELETE FROM employee WHERE id=?",id)
	if err != nil {
		return nil, err
	}
	
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }
    if rowsAffected == 0 {
        return nil, nil
    }
	var emp model.Employee

    return &emp, nil
}