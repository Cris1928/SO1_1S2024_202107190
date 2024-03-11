CREATE DATABASE Proyecto1;

Use Proyecto1;

CREATE TABLE ram (
    Fecha VARCHAR(50) PRIMARY KEY,
    Porcentaje INT
);

select * from ram;

INSERT INTO ram (Fecha, Porcentaje) VALUES
    ('0-00-00 00:00:00', 0);
   
 DROP  table ram;

 TRUNCATE TABLE ram;



CREATE TABLE cpu (
    Fecha VARCHAR(50) PRIMARY KEY,
    Porcentaje INT
);

select * from cpu;

INSERT INTO cpu (Fecha, Porcentaje) VALUES
    ('0-00-00 00:00:00', 0);
   
 DROP  table cpu;

 TRUNCATE TABLE cpu;