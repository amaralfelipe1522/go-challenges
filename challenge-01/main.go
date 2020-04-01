package main

/*
Acceptance Criteria
In order to successfully complete this challenge, your project will have to:

	Collect the CPU utilization of your machine
	Collect the RAM Utilization of your machine
	Collect the Backing Storage Utilization of your machine
	Display the results in a friendly fashion in the console when go run main.go is executed.
*/

import (
	"fmt"
	"log"
	"runtime"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func printMemory() {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)

	fmt.Printf("Memória alocada no momento é %v bytes (%vmb).\n", mem.Alloc, bToMb(mem.Alloc))
	fmt.Printf("Memória utilizada no Sistema Operacional no momento é %v bytes (%vmb).\n", mem.Sys, bToMb(mem.Sys))
}

func printCPU() {
	status, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("Falha na leitura")
	}

	fmt.Println("Iniciando leitura da CPU...")
	for _, s := range status.CPUStats {
		fmt.Printf("O uso de CPU atual no Sistema Operacional é %v bytes (%vmb).\n", s.System, bToMb(s.System))
		fmt.Printf("O uso de CPU atual do usuário é %v bytes (%vmb).\n", s.User, bToMb(s.User))
	}
	fmt.Println("Leitura da CPU finalizada!")
}

func printDisc() {
	status, err := linuxproc.ReadDiskStats("/proc/diskstats")
	if err != nil {
		log.Fatal("Falha na leitura")
	}
	fmt.Println("Iniciando leitura do disco...")

	for _, s := range status {
		fmt.Printf("Major %v.\n", s.Major)
		fmt.Printf("Minor %v.\n", s.Minor)
		fmt.Printf("Name %v.\n", s.Name)
		fmt.Printf("ReadIOs %v.\n", s.ReadIOs)
	}
	fmt.Println("Leitura de disco finalizada.")

	/*
		status, err := linuxproc.ReadDisk("/proc/diskstats")
		if err != nil {
			log.Fatal("Falha na leitura")
		}
		fmt.Println("Iniciando leitura do disco...")
		fmt.Printf("Uso de disco atual é %v.\n", status.Used)
		fmt.Printf("Qtd de disco livre é %v.\n", status.Free)
		fmt.Printf("Total de disco é %v.\n", status.All)
		fmt.Println("Leitura de disco finalizada.")
	}*/
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	printMemory()
	printCPU()
	printDisc()
}
