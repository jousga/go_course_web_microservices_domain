# go_course_web Microservice: Domain

URL del curso: https://www.udemy.com/course-dashboard-redirect/?course_id=4809678

Este proyecto forma parte de la transformación del proyecto creado en el repositorio [go_course_web](https://github.com/jousga/go_course_web) pero transformándolo a microservicios

El conjunto de microservicios que forman parte de este proyecto son:
1. Domain: https://github.com/jousga/go_course_web_microservices_domain
2. Meta: https://github.com/jousga/go_course_web_microservices_meta
3. User: https://github.com/jousga/go_course_web_microservices_user
4. Course: https://github.com/jousga/go_course_web_microservices_course
5. Enrollment: https://github.com/jousga/go_course_web_microservices_enrollment
6. Lib Response: https://github.com/jousga/go_course_web_microservices_lib_response

Para limpiar el proyecto y descargar las dependencias tenemos que usar el comando de GO:

```text
go mod tidy
```

#GIT
Para este proyecto hemos definido TAGS en GITHUB que nos servirán para gestionar las versiones de nuestros packages y poder
importar diferentes versiones mediante los tags. En git este TAG lo podemos crear dentro de RELEASE, por lo tanto, lo que tenemos que hacer 
es:
1. Seleccionar la opción "Create a new release" y
2. Introducimos un nombre para la release con el formato recomendado por Go. Ej: v0.0.1
3. En la misma pantalla seleccionamos "Choose a tag" y le ponemos como tag la misma versión que en el punto anterior y clicamos en "Create new tag: V0.0.1"
4. Pulsamos el botón "Publish release"

Es importante que el formato de la versión/tag sea vx.x.x para que Go lo entienda como versión

También se puede importar mediante commit o branch, pero es mejor utilizar los tags para que quede ordenado.

En la página de Go https://go.dev/blog/publishing-go-modules se explica como recomienda Go el generar las versiones 