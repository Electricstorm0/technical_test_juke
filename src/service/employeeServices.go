package service

import (
	"errors"
	"technical_test/src/model"
	"technical_test/src/repository"
)

func CreateEmployeeServices(emp *model.Employee)(int64,error)   {
	checkAvailableEmail,err := repository.GetByEmail(emp.Email)
	if err != nil {
		return 0, err
	}
	if checkAvailableEmail != nil  {
		return 0, errors.New("email has registered")
	}
	newEmployee,err:= repository.Create(emp)
	if err != nil {
		return 0, err
	}
	if emp.Name== "" || emp.Email=="" {
		return 0, errors.New("name and email are reuqired")
	}


	return newEmployee,nil
}

func GetAllService() ([]model.Employee,error) {
return repository.GetAllRepo()
}
func GetByIdService(id int64) (*model.Employee,error) {
employee,err:= repository.GetById(id)
if err != nil {
	return nil, err
}
if employee == nil {
	return nil, errors.New("employee not found")
}
return employee,nil
}

func UpdateServices(id int64, emp model.Employee) (*model.Employee,error)  {
checkExistingEmployee,err:= repository.GetById(id)
if err != nil {
	return nil, err
}
if checkExistingEmployee == nil{
	return nil, errors.New("employee not found")
} 

updateEmployees,err:= repository.Update(id,emp)
if err != nil {
	return nil, err
}

return updateEmployees,nil
}

func DeleteServices(id int64) (*model.Employee,error)  {
checkExistingEmployee,err:= repository.GetById(id)
if err != nil {
	return nil, err
}
if checkExistingEmployee == nil{
	return nil, errors.New("employee not found")
} 
deletedEmployee,err:=repository.Delete(id)
if err != nil {
	return nil, err
}
return deletedEmployee, nil
}