package cd_nodes_go

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/RulezKT/cd_consts_go"
)

func LoadNodesCoords() *[]cd_consts_go.NodesJsonStruct {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// pathToFile := filepath.Join(dir, "files", "nodes_file.json")
	pathToFile := filepath.Join(dir, "files", "nodes.json")

	file, _ := os.ReadFile(pathToFile)

	data := string(file)

	//fmt.Println(data)

	stringArr := strings.Split(data, "]")
	nodesInfo := make([]cd_consts_go.NodesJsonStruct, len(stringArr)-1)
	for i := 0; i < len(stringArr)-2; i++ { //29530  -2 чтобы не забирать последнюю кавычку в файле ,там просто пробел и после ещё 1 строка с пробелом

		/*
			if i == 0 || i == 1 || i == len(stringArr)-1 {

				fmt.Println("i = ", i, "arr = ", stringArr[i])
			}
		*/

		substr := strings.Split(stringArr[i][2:], ",") //для 0 элемента убрать [[ , для остальных убрать запятую и начальную скобку

		nodesInfo[i].Time, _ = strconv.ParseFloat(substr[0], 64)
		nodesInfo[i].Which = substr[1][1 : len(substr[1])-1] //-кавычки
		/*
			if i == 0 || i == 1 || i == len(stringArr)-3 {
				fmt.Println("i == ", i, " Time: ", nodesInfo[i].Time, " Which: ", nodesInfo[i].Which, " Eps: ", nodesInfo[i].Eps)
			}
		*/

	}

	return &nodesInfo
}
