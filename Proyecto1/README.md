# ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ MANUAL TECNICO

El siguiente proyecto tiene como objetivo implementar un sistema de monitoreo derecursos del sistema y gestión de procesos, empleando varias tecnologías y lenguajes de programación.

## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ Kernel Modules

## Módulo de RAM
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "ram_so1_1s2024.ko" este recoge el porcentaje de la ram utilizada y el porcentaje de la ram libre.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "ram_so1_1s2024.ko" ejecutar el comando "sudo make all" y este generara los archivos necesarios para su despliegue en el directorio /proc
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod ram_so1_1s2024.ko"
- Leer el archivo escrito cat ram_so1_1s2024

## Módulo de CPU
Este modulo está ubicado en el directorio /proc, desplegado desde el archvivo "cpu_so1_1s2024.ko" este recoge el porcentaje del cpu utilizado y el porcentaje del cpu libre, tambien recoge los procesos padres e hijos que la computadora utilia en ese momento.
Pasos para utilizarla desde el directorio /proc:
- En el directorio del archivo "cpu_so1_1s2024.ko" ejecutar el comando "sudo make all" y este generara los archivos necesarios para su despliegue en el directorio /proc
- Para desplegarlo en el directorio /proc ejecutar el comando "sudo insmod cpu_so1_1s2024.ko"
- Leer el archivo escrito cat cpu_so1_1s2024

## ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ ‎ PLATAFORMA DE MONITOREO 
## Monitoreo en tiempo real
Este posee el porcentaje de ram utilizada y libre desplegada en una grafica circular, de igual manera con el cpu, esto utilizando la libreria "react-cjartjs-2".
Desde el backend posee los ednpoints "localhost:5200/api/cpu" y "localhost:5200/api/ram" estos poseen la informacion del ultimo registro de su porcentaje en utilizacion de la ram y el cpu.

func RAMModuleHandler(w http.ResponseWriter, r *http.Request) {

	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		fmt.Println("eerror", err)
	}
	//fmt.Println("-------------------- RAM --------------------")
	var ram_info Ram
	err = json.Unmarshal([]byte(outRam), &ram_info)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((ram_info))

}






![WhatsApp Image 2024-03-17 at 4 28 27 AM](https://github.com/Cris1928/SO1_1S2024_202107190/assets/98928867/973c43d9-1fde-4ef8-b3a3-b4a22fa54710)


