# Microservicio de Minas

Este microservicio hace parte de la implementación de la capa de aplicación de un sistema IoT para el monitoreo de minas subterraneas.

## Ejecución (development)

1. Tener instalado [Make](https://www.gnu.org/software/make/) y [Go](https://go.dev/) (más adelante se incluirá un Dockerfile para que no se tenga que instalar Go directamente)
2. Copiar **.env.example** y renombrarlo a **.env.local** ingresando las variables indicadas en este.
3. Construir los contenedores de Docker necesarios

```bash
make run-docker-dev
```

4. Levantar el servidor de desarrollo:

````bash
make run
```a
````
