# Documentación Técnica del Proyecto

Este documento proporciona instrucciones claras y precisas para configurar y ejecutar el proyecto en un entorno local.

## Requisitos del Sistema

- **Go:** Versión 1.22.5 o superior.
- **PostgreSQL:** Una base de datos PostgreSQL en funcionamiento.
- **Variables de Entorno:** Un archivo `.env` con las siguientes variables:
  - `DB_HOST`: Host de la base de datos (por ejemplo, `localhost`).
  - `DB_PORT`: Puerto de la base de datos (por ejemplo, `5432`).
  - `DB_USER`: Usuario de la base de datos.
  - `DB_PASSWORD`: Contraseña de la base de datos.
  - `DB_NAME`: Nombre de la base de datos.
  - `JWT_SECRET`: Una clave secreta para la firma de tokens JWT.

## Dependencias

El proyecto utiliza las siguientes dependencias principales:

- **Gin:** Un framework web para Go.
- **GORM:** Un ORM para Go.
- **GoDotEnv:** Para cargar variables de entorno desde un archivo `.env`.
- **JWT-Go:** Para la generación y validación de JSON Web Tokens.

Para instalar todas las dependencias, ejecute el siguiente comando:

```bash
go mod tidy
```

## Ejecución del Proyecto

1. **Clonar el repositorio:**

   ```bash
   git clone <https://github.com/ChrisTheAbysswalker/rootly-backend>
   ```

2. **Crear el archivo `.env`:**

   Cree un archivo `.env` en la raíz del proyecto y añada las variables de entorno mencionadas en la sección "Requisitos del Sistema".

3. **Ejecutar la aplicación:**

   ```bash
   go run .
   ```

   El servidor se iniciará en el puerto `8000`.

## Endpoints de la API

A continuación se muestra una lista de los endpoints disponibles en la API:

### Autenticación

- **POST** `/api/auth/register`: Crea una nueva cuenta de usuario.
- **POST** `/api/auth/login`: Valida las credenciales y devuelve un token JWT.

### Especies

- **GET** `/api/especies`: Lista el inventario botánico completo.
- **GET** `/api/especies/:id`: Obtiene una especie por su ID.
- **POST** `/api/especies`: Registra una nueva planta en el sistema.
- **PUT** `/api/especies/:id`: Actualiza los datos de una planta existente.
- **DELETE** `/api/especies/:id`: Elimina una especie del inventario.
- **GET** `/api/ecosistema/stats`: Obtiene métricas (humedad, salud, alertas).
- **GET** `/api/familias`: Obtiene la lista de familias de plantas.
- **GET** `/api/estados`: Obtiene la lista de estados de salud de las plantas.
- **GET** `/api/registro_salud/:id`: Obtiene el registro de salud de una planta por su ID.
- **POST** `/api/registro_salud`: Crea un nuevo registro de salud para una planta.
- **PUT** `/api/registro_salud/:id`: Actualiza un registro de salud existente.

### Personal

- **GET** `/api/staff`: Obtiene la lista de personal.
