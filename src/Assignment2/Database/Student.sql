CREATE database students;

USE students; 

CREATE TABLE Student (
	StudentID varchar (9) NOT NULL,
    SName varchar (100) NOT NULL,
    DOB date NOT NULL,
    Address varchar (100) NOT NULL,
    PhoneNo char(8) NOT NULL,
    
    PRIMARY KEY (StudentID)
    
    );