# Taller de código: Desafío final


> [!IMPORTANTE]
> El proyecto corre solo en Local, se intento Dockerizar pero en el proceso testear la API via postman se generaban algunos errores. Igual se dejaron los cambios para terminar el proceso de manera exitosa en el futuro.

> [!TIP]
> Modificar el Archivo .env con la informacion correspondiente." 


### Integrantes

[Ismael Carvajal](https://github.com/ismaelc511)

[Juan Ignacio Delena](https://github.com/JuanIgnacioDelena)

[Marlon Yepes](https://github.com/myepes82)

[Martha De la Ossa ](https://github.com/marthadelaossa)
         


## Repositorio Proyecto

[Repositorio](https://github.com/marthadelaossa/FinalBackIIIGo.git)



## Objetivo
A continuación se plantea un desafío integrador que nos permitirá evaluar todos los temas que hemos visto en la cursada.

## Sistema de reserva de turnos
Se desea implementar una API que permita administrar la reserva de turnos para una clínica odontológica. Esta debe cumplir con los siguientes requerimientos:


### Administración de Odontólogos

Listar, agregar, modificar y eliminar odontólogos. Registrar apellido, nombre y matrícula de los mismos. Se desea el desarrollo de un CRUD para la entidad Dentista.

- POST  : Agregar un dentista.
- GET   : Traer dentista por ID.
- PUT   : Actualizar dentista.
- PATCH : Actualizar dentista por alguno de sus campos.
- DELETE: Eliminar el dentista.

### Administración de Pacientes

Listar, agregar, modificar y eliminar pacientes. De cada uno se almacenan: nombre, apellido, domicilio, DNI y fecha de alta. Se desea el desarrollo de un CRUD para la entidad Paciente.

- POST  : Agregar un paciente.
- GET   : Traer paciente por ID.
- PUT   : Actualizar paciente.
- PATCH : Actualizar paciente por alguno de sus campos.
- DELETE: Eliminar el paciente.

### Registrar Turnos

Se tiene que poder permitir asignar a un paciente un turno con un odontólogo a una determinada fecha y hora. Al turno se le debe poder agregar una
descripción. Se desea el desarrollo de un CRUD para la entidad Turno. 

- POST  : Agregar un Turno.
- GET   : Traer Turno por ID.
- PUT   : Actualizar Turno.
- PATCH : Actualizar Turno por alguno de sus campos.
- DELETE: Eliminar Turno.
- POST  : agregar turno por DNI del paciente y matrícula del dentista.
- GET   : traer turno por DNI del paciente. Debe traer el detalle del turno (Fecha-Hora, descripción, Paciente y Dentista) y el dni deberá ser recibido por
          QueryParams.

### Seguridad y Middleware
- Se tiene que proveer cierta seguridad al momento de realizar POST, PUT, PATCH y DELETE. Esta seguridad mediante
autenticación deberá estar implementada mediante un middleware.


## Requerimientos técnicos

La aplicación debe ser desarrollada en diseño orientado a paquetes:

- Capa/dominio de entidades de negocio.
- Capa/dominio de acceso a datos (Repository).
- Capa de acceso a datos (base de datos): es la base de datos de nuestro sistema.

Podrás utilizar cualquier base de datos relacional modelado a través de un modelo entidad-relación, como H2 o MySQL, o no relacional, como MongoDB.

- Capa/dominio service.
- Capa/dominio handler.

¡Mucha suerte!
