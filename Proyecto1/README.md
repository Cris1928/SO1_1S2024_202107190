#             MANUAL TECNICO

El siguiente proyecto tiene como objetivo implementar un sistema de monitoreo derecursos del sistema y gestión de procesos, empleando varias tecnologías y lenguajes de programación.

## ‎ ‎ ‎ ‎ ‎ ‎ Kernel Modules

## Módulo de RAM
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "ram_so1_1s2024.ko" este recoge el porcentaje de la ram utilizada y el porcentaje de la ram libre.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "ram_so1_1s2024.ko" ejecutar el comando "sudo make all"
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod ram_so1_1s2024.ko"
- Leer el archivo escrito cat ram_so1_1s2024

## Módulo de CPU
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "cpu_so1_1s2024.ko" este recoge el porcentaje del cpu utilizado y el porcentaje del cpu libre, tambien recoge los procesos padres e hijos que la computadora utilia en ese momento.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "cpu_so1_1s2024.ko" ejecutar el comando "sudo make all"
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod cpu_so1_1s2024.ko"
- Leer el archivo escrito cat cpu_so1_1s2024
