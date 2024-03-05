package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	//"github.com/rs/cors"
	"github.com/shirou/gopsutil/cpu"
)

type Process struct {
	Pid     int    `json:"pid"`
	Nombre  string `json:"name"`
	Usuario int    `json:"usuario"`
	Estado  int    `json:"state"`
	Ram     int    `json:"ram"`
	Padre   int    `json:"pidPadre"`
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

type Ip struct {
	Ip string `json:"ip"`
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

func Index(x http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(x, "sserver")

}

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
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// No hacer nada aquí, solo devolver un código 200 o 404
		w.WriteHeader(http.StatusNoContent)
	}).Methods("GET", "POST")
	router.HandleFunc("/cpu", CPUModuleHandler).Methods("GET", "POST")
	router.HandleFunc("/ram", RAMModuleHandler).Methods("GET", "POST")

	go func() {
		log.Fatal(http.ListenAndServe(":5200", handlers.CORS()(router)))
	}()

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
			go sendToAPI("/cpu", cpu_info)

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
