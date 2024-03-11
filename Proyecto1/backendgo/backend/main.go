package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	//	"strconv"

	//"io/ioutil"
	"log"
	//"math"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"

	//"github.com/rs/cors"
	"os"
	//"github.com/shirou/gopsutil/cpu"
)

var conexion = ConectarBD()

func ConectarBD() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	conexionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	conexion, err := sql.Open("mysql", conexionString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Conexion con MySQL Correcta")
	}
	return conexion
}

type Datacpu struct {
	Fecha      string `json:"fecha"`
	Porcentaje int    `json:"porcentaje"`
}

type child struct {
	Pid    int    `json:"pid"`
	Nombre string `json:"name"`
	Estado int    `json:"state"`
	Padre  int    `json:"pidPadre"`
}

type Process struct {
	Pid     int     `json:"pid"`
	Nombre  string  `json:"name"`
	Usuario int     `json:"usuario"`
	Estado  int     `json:"state"`
	Ram     int     `json:"ram"`
	Padre   int     `json:"pidPadre"`
	Chil    []child `json:"child"`
}
type Cpu struct {
	Porcentaje int       `json:"cpu_porcentaje"`
	Procesos   []Process `json:"processes"`
}

type Ram struct {
	Total      int `json:"totalRam"`
	En_uso     int `json:"memoriaEnUso"`
	Libre      int `json:"libre"`
	Porcentaje int `json:"porcentaje"`
}
type Dataram struct {
	Fecha      string `json:"fecha"`
	Porcentaje int    `json:"porcentaje"`
}

type Ip struct {
	Ip string `json:"ip"`
}
type Respuestacpu struct {
	Mensaje string `json:"mensaje"`
}

type Respuestaram struct {
	Mensaje string `json:"mensaje"`
}

func Index(x http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(x, "sserver")

}

var process *exec.Cmd

func StartProcess(w http.ResponseWriter, r *http.Request) {
	// Crear un nuevo proceso con un comando de espera
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	// Obtener el comando con PID
	process = cmd

	fmt.Fprintf(w, "Proceso iniciado con PID: %d y estado en espera", process.Process.Pid)
}

func StopProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGSTOP al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al detener el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d detenido", pid)
}

func ResumeProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al reanudar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d reanudado", pid)
}

func KillProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d ha terminado", pid)
}

func Insertdatacpu(info map[string]interface{}) {
	jsonData, err := json.Marshal(info)

	if err != nil {
		fmt.Println("eram1")
		log.Fatal(err)
	}
	fmt.Println("string(jsonData)")
	fmt.Println(string(jsonData))

	response := &http.Response{}
	request, err := http.NewRequest("POST", "http://localhost:5200/registrocpu", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("eram2")
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("eram3")
		log.Fatal(err)
	}
	defer response.Body.Close()

	var respuesta Respuestacpu
	json.NewDecoder(response.Body).Decode(&respuesta)
	fmt.Println(respuesta.Mensaje)
	fmt.Println("respuesta.Mensaje")

}
func Insertdataram(info map[string]interface{}) {

	jsonData, err := json.Marshal(info)

	if err != nil {
		fmt.Println("eram1")
		log.Fatal(err)
	}
	response := &http.Response{}
	request, err := http.NewRequest("POST", "http://localhost:5200/registroram", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("eram2")
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("eram3")
		log.Fatal(err)
	}
	defer response.Body.Close()

	var respuesta Respuestaram
	json.NewDecoder(response.Body).Decode(&respuesta)
	fmt.Println(respuesta.Mensaje)
	fmt.Println("respuesta.Mensaje")

}
func Registrocpu(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	//var discoReg Disco
	//var respuesta Respuesta
	var dataCpu Datacpu
	//now := time.Now()
	//re := fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	var respuesta Respuestacpu

	json.NewDecoder((request.Body)).Decode(&dataCpu)
	respuesta.Mensaje = "Registro Exitoso"

	query := `INSERT INTO cpu (Fecha, Porcentaje) VALUES (?,?);`
	result, err := conexion.Exec(query, dataCpu.Fecha, dataCpu.Porcentaje)
	if err != nil {
		fmt.Println(err)
		respuesta.Mensaje = "Error al registrar"
	}
	fmt.Println(result)
	fmt.Println(respuesta)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}
func Registroram(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	//var discoReg Disco
	//var respuesta Respuesta
	var ramReg Dataram
	//now := time.Now()
	//re := fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	var respuesta Respuestaram

	json.NewDecoder((request.Body)).Decode(&ramReg)
	respuesta.Mensaje = "Registro Exitoso"

	query := `INSERT INTO ram (Fecha, Porcentaje) VALUES (?,?);`
	result, err := conexion.Exec(query, ramReg.Fecha, ramReg.Porcentaje)
	if err != nil {
		fmt.Println(err)
		respuesta.Mensaje = "Error al registrar"
	}
	fmt.Println(result)
	fmt.Println(respuesta)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(respuesta)
}

func Getdatcpu(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var lista []Datacpu
	//now := time.Now()
	//re := fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	query := "select * from cpu;"
	result, err := conexion.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		var datcpu Datacpu

		err = result.Scan(&datcpu.Fecha, &datcpu.Porcentaje)
		if err != nil {
			fmt.Println(err)
		}
		lista = append(lista, datcpu)
	}
	json.NewEncoder(response).Encode(lista)
}
func Getdatram(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var lista []Dataram
	//now := time.Now()
	//re := fmt.Sprint("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	query := "select * from ram;"
	result, err := conexion.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		var datram Dataram

		err = result.Scan(&datram.Fecha, &datram.Porcentaje)
		if err != nil {
			fmt.Println(err)
		}
		lista = append(lista, datram)
	}
	json.NewEncoder(response).Encode(lista)
}

func CPUModuleHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(" ==================== DATOS MODULO CPU ==================== ")
	//fmt.Println(" ")

	cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	outCpu, err := cmdCpu.CombinedOutput()
	if err != nil {
		fmt.Println("eerror", err)
		//fmt.Println(outCpu)
	}
	//	output := string(outCpu[:])
	//	fmt.Println(output)

	//---CPU
	//fmt.Println("-------------------- CPU --------------------")
	var cpu_info Cpu
	err = json.Unmarshal([]byte(outCpu), &cpu_info)
	if err != nil {
		fmt.Println(err)
	}
	//Mandar respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((cpu_info))
}
func RAMModuleHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(" ==================== DATOS MODULO RAM ==================== ")
	//fmt.Println(" ")

	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		fmt.Println("eerror", err)
		//fmt.Println(outCpu)
	}
	//	outputt := string(outRam[:])
	//	fmt.Println(outputt)
	//---RAM
	//fmt.Println("-------------------- RAM --------------------")
	var ram_info Ram
	err = json.Unmarshal([]byte(outRam), &ram_info)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("ram_info")
	//fmt.Println(ram_info)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((ram_info))

}
func main() {
	router := mux.NewRouter()
	//	routersql := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// No hacer nada aquí, solo devolver un código 200 o 404
		w.WriteHeader(http.StatusNoContent)
	}).Methods("GET", "POST")
	router.HandleFunc("/cpu", CPUModuleHandler).Methods("GET", "POST")
	router.HandleFunc("/ram", RAMModuleHandler).Methods("GET", "POST")
	router.HandleFunc("/datram", Getdatram).Methods("GET")
	router.HandleFunc("/registroram", Registroram).Methods("POST")
	router.HandleFunc("/datcpu", Getdatcpu).Methods("GET")
	router.HandleFunc("/registrocpu", Registrocpu).Methods("POST")

	router.HandleFunc("/api/data/ram", GetDataram).Methods("GET")
	router.HandleFunc("/api/data/cpu", GetDatacpu).Methods("GET")

	router.HandleFunc("/start", StartProcess)
	router.HandleFunc("/stop", StopProcess)
	router.HandleFunc("/resume", ResumeProcess)
	router.HandleFunc("/kill", KillProcess)

	go func() {
		log.Fatal(http.ListenAndServe(":5200", handlers.CORS()(router)))
	}()
	interval := 5

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			//	time.Sleep(1000)
			now := time.Now()
			second1 := now.Second()
			re := strconv.Itoa(now.Year()) + "-" + fmt.Sprintf("%02d", now.Month()) + "-" + fmt.Sprintf("%02d", now.Day()) + "-" + fmt.Sprintf("%02d", now.Hour()) + "-" + fmt.Sprintf("%02d", now.Minute()) + "-" + fmt.Sprintf("%02d", second1)
			fmt.Println("re")
			fmt.Println(re)

			fmt.Println(" ==================== DATOS MODULO CPU ==================== ")
			fmt.Println(" ")

			cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
			outCpu, err := cmdCpu.CombinedOutput()
			if err != nil {
				fmt.Println("eerror", err)
				//fmt.Println(outCpu)
			}
			//	output := string(outCpu[:])
			//	fmt.Println(output)

			//---CPU
			fmt.Println("-------------------- CPU --------------------")
			var cpu_info Cpu
			err = json.Unmarshal([]byte(outCpu), &cpu_info)
			if err != nil {
				fmt.Println(err)
			}
			//Mandar respuesta
			go sendToAPI("/cpu", cpu_info)
			infocpu := map[string]interface{}{
				"Fecha":      re,
				"Porcentaje": cpu_info.Porcentaje,
			}

			Insertdatacpu(infocpu)

			//	time.Sleep(1000)
			noww := time.Now()
			second := noww.Second()
			if now.Second() == noww.Second() {
				second = second + 1
			}
			ree := strconv.Itoa(now.Year()) + "-" + fmt.Sprintf("%02d", now.Month()) + "-" + fmt.Sprintf("%02d", now.Day()) + "-" + fmt.Sprintf("%02d", now.Hour()) + "-" + fmt.Sprintf("%02d", now.Minute()) + "-" + fmt.Sprintf("%02d", second)

			fmt.Println(" ==================== DATOS MODULO RAM ==================== ")
			fmt.Println(" ")

			cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
			outRam, err := cmdRam.CombinedOutput()
			if err != nil {
				fmt.Println("eerror", err)
				//fmt.Println(outCpu)
			}
			//	outputt := string(outRam[:])
			//	fmt.Println(outputt)
			//---RAM
			fmt.Println("-------------------- RAM --------------------")
			var ram_info Ram
			err = json.Unmarshal([]byte(outRam), &ram_info)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println("ram_info")
			///fmt.Println(ram_info)
			go sendToAPI("/ram", ram_info)
			fmt.Println(ram_info.Porcentaje)
			fmt.Println(ram_info.En_uso)
			inforam := map[string]interface{}{
				"Fecha":      ree,
				"Porcentaje": ram_info.Porcentaje,
			}
			Insertdataram(inforam)
			//Mandar respuesta

		}
	}
}

func sendToAPI(route string, data interface{}) {

	url := fmt.Sprintf("http://localhost:5200%s", route)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error al convertir datos a JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error al enviar datos a la API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("La API respondió con un código de estado no válido:", resp.StatusCode)
		return
	}

	log.Printf("Datos enviados a la ruta %s\n", route)
}

func GetDataram(w http.ResponseWriter, r *http.Request) {
	// Configuración de la conexión a la base de datos
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/Proyecto1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Realizar la consulta
	query := "SELECT Fecha, Porcentaje FROM ram ORDER BY Fecha DESC LIMIT 10"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Almacenar los resultados en un slice de Datos
	var data []Dataram
	for rows.Next() {
		var d Dataram
		err := rows.Scan(&d.Fecha, &d.Porcentaje)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d)
	}

	// Convertir a formato JSON y enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetDatacpu(w http.ResponseWriter, r *http.Request) {
	// Configuración de la conexión a la base de datos
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/Proyecto1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Realizar la consulta
	query := "SELECT Fecha, Porcentaje FROM cpu ORDER BY Fecha DESC LIMIT 10"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Almacenar los resultados en un slice de Datos
	var data []Datacpu
	for rows.Next() {
		var d Datacpu
		err := rows.Scan(&d.Fecha, &d.Porcentaje)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, d)
	}

	// Convertir a formato JSON y enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

/*

func getRam(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			fmt.Println(" ==================== DATOS MODULO RAM ==================== ")
			fmt.Println(" ")

			cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
			outRam, err := cmdRam.CombinedOutput()
			if err != nil {
				fmt.Println("eerror", err)
				//fmt.Println(outCpu)
			}
			//	outputt := string(outRam[:])
			//	fmt.Println(outputt)
			//---RAM
			fmt.Println("-------------------- RAM --------------------")
			var ram_info Ram
			err = json.Unmarshal([]byte(outRam), &ram_info)
			if err != nil {
				fmt.Println(err)
			}
			//-fmt.Println("ram_info")
			//-fmt.Println(ram_info.Libre)
			//Mandar respuesta
			json.NewEncoder(response).Encode(ram_info)

		}

	}

}


func postScheduledData() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println(" ==================== DATOS MODULO CPU ==================== ")
			fmt.Println(" ")

			cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
			outCpu, err := cmdCpu.CombinedOutput()
			if err != nil {
				fmt.Println("eerror", err)
				//fmt.Println(outCpu)
			}
			//	output := string(outCpu[:])
			//	fmt.Println(output)

			//---CPU
			fmt.Println("-------------------- CPU --------------------")
			var cpu_info Cpu
			err = json.Unmarshal([]byte(outCpu), &cpu_info)
			if err != nil {
				fmt.Println(err)
			}
			//Mandar respuesta
			url := "http://34.74.125.199:5200/cpu"
			//Mandar cpu_info que es un json
			p_cpu, err := cpu.Percent(time.Second, false)
			if err != nil {
				fmt.Println(err)
			}
			cpu_info.Porcentaje = int(math.Round(p_cpu[0]))
			jsonValue, _ := json.Marshal(cpu_info)
			//Mandar el json a la url
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Println("ji")
				fmt.Println(err)
			} else {
				defer resp.Body.Close()
				responseBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("ja")
					fmt.Println(err)
				} else {
					// Imprimir el contenido del mensaje enviado desde el servidor

					fmt.Println("jo")
					fmt.Println("\x1b[32m" + string(responseBody) + "\x1b[0m")
				}
			}
			fmt.Println(" ")
			fmt.Println(" ==================== DATOS MODULO RAM ==================== ")
			fmt.Println(" ")

			cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
			outRam, err := cmdRam.CombinedOutput()
			if err != nil {
				fmt.Println("eerror", err)
				//fmt.Println(outCpu)
			}
			//	outputt := string(outRam[:])
			//	fmt.Println(outputt)
			//---RAM
			fmt.Println("-------------------- RAM --------------------")
			var ram_info Ram
			err = json.Unmarshal([]byte(outRam), &ram_info)
			if err != nil {
				fmt.Println(err)
			}
			//-fmt.Println("ram_info")
			//-fmt.Println(ram_info.Libre)
			//Mandar respuesta
			url = "http://34.74.125.199:5200/ram"
			//Mandar ram_info que es un json
			jsonValue, _ = json.Marshal(ram_info)
			//fmt.Println("jsonValue")
			//fmt.Println(jsonValue)
			//Mandar el json a la url
			resp, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Println(err)
			} else {
				defer resp.Body.Close()
				responseBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				} else {
					// Imprimir el contenido del mensaje enviado desde el servidor
					fmt.Println("\x1b[32m" + string(responseBody) + "\x1b[0m")
				}
			}
			fmt.Println(" ")

		}
	}
}
*/
