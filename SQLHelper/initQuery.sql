CREATE TABLE student(
    studentname VARCHAR(255) PRIMARY KEY,
    class INT,
    mark INT
);

INSERT INTO student VALUES('Karthikeyan', 11, 195);
INSERT INTO student VALUES('Sri Kanth', 11, 183);
INSERT INTO student VALUES('Kumar', 11, 199);

UPDATE student SET class = 11, mark = 195 WHERE name = 'sri ram';