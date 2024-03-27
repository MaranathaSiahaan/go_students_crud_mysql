//Maranatha Siahaan, 11322018, D3 TI 1 22
package routes

import (
    "github.com/gorilla/mux"
    "github.com/MaranathaSiahaan/go_students_crud_mysql/pkg/controllers"
)

func RegisterStudentsRoutes(router *mux.Router) {
    router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")       
    router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")        
    router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")  
    router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")  
    router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE") 
}
