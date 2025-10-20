# Task Manager - Go & Gin

Este es un proyecto de gestor de tareas simple con un backend desarrollado en Go utilizando el framework Gin y un frontend básico en HTML, CSS y JavaScript. La aplicación permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) sobre tareas, organizándolas en tres estados: Pendiente, En Progreso y Hecho.

## Características

- **Gestión de Tareas**: Crea, visualiza, actualiza y elimina tareas.
- **Múltiples Estados**: Organiza las tareas en tres columnas: `Pending`, `In Progress`, y `Done`.
- **Arrastrar y Soltar (Drag and Drop)**: Cambia el estado de una tarea simplemente arrastrándola a otra columna.
- **Operaciones en Lote**: Agrega, actualiza y elimina múltiples tareas en una sola operación haciendo clic en el botón "Accept".
- **Soporte para Múltiples Bases de Datos**: Configurable para usar **PostgreSQL** o **MongoDB** como backend de base de datos.
- **Contenerización**: Totalmente contenerizado con Docker y Docker Compose para un despliegue y desarrollo sencillos.

## Tech Stack

- **Backend**: Go, Gin, GORM (para PostgreSQL), Mongo-Driver (para MongoDB)
- **Frontend**: HTML, CSS, JavaScript
- **Base de Datos**: PostgreSQL, MongoDB
- **Contenerización**: Docker, Docker Compose

## Prerrequisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Cómo Empezar

Sigue estos pasos para levantar el proyecto localmente.

### 1. Clonar el Repositorio

```bash
git clone git@github-tariq:tariq-ventura/Proyecto-go.git
cd Proyecto-go
```

### 2. Configurar variables de entorno

crea un archivo ```.env``` a partir del ejemplo proporcionado 

```bash
cp .env.example .env
```

Abre el archivo .env y ajusta las variables según tu configuración. La variable más importante es DB_CONTEXT, que te permite elegir la base de datos a utilizar.

- Para PostgreSQL
```
DB_CONTEXT="postgres"
```

- Para MongoDB
```
DB_CONTEXT="mongo"
```

### 3. Levantar los Contenedores

Usa Docker Compose para construir y levantar todos los servicios.


```
docker-compose up --build -d
```

### 4. Uso
Una vez que los contenedores estén en funcionamiento, puedes acceder a la aplicación en tu navegador:

- [ http://localhost:3000]( http://localhost:3000)

## API Endpoints

- GET /: Sirve la aplicación frontend (index.html).
- GET /api/tasks: Obtiene todas las tareas.
- POST /api/tasks: Crea una o más tareas nuevas.
- PUT /api/tasks: Actualiza una o más tareas existentes.
- DELETE /api/tasks: Elimina una o más tareas.
- GET /healthz: Endpoint de health check.


## Conexion a la base

Para conectarte a la base de datos PostgreSQL usando psql:

### PostgreSQL

```
docker exec -it postgres-db psql -U <USER_DB> -d <DB_NAME>
```

(Reemplaza <USER_DB> y <DB_NAME> con los valores de tu archivo .env).

### MongoDB

Puedes conectarte a MongoDB usando una herramienta como Mongsh:

```
mongosh "mongodb://localhost:27017"
```