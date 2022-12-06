package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Verificar se um diretório existe
	_, err := os.Stat("teste")
	if err != nil {
		fmt.Println("O arquivo não existe")

		// Para criar pastas
		err = os.Mkdir("teste", os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Para criar pastas com subdiretórios
	err = os.MkdirAll("teste2/novo", os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	// Para saber o diretório atual
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)

	// Para renomear um diretório
	_, err = os.Stat("teste2/atualizado")
	if err != nil {
		err = os.Rename("teste2/novo", "teste2/atualizado")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Para remover um diretório
	err = os.Remove("teste")
	if err != nil {
		log.Fatal(err)
	}

	// Para remover um diretório com subdiretórios
	err = os.RemoveAll("teste2/novo")
	if err != nil {
		log.Fatal(err)
	}

	// Para encontrar arquivos dentro de um diretório
	files, err := filepath.Glob(path + string(os.PathSeparator) + "**" + string(os.PathSeparator) + "*.go")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	// Para criar arquivos
	file, err := os.Create("teste")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("Esse é o texto que irá para dentro do arquivo")
	bn, err := file.Write(data)
	fmt.Printf("Quantidade de bytes escritos no arquivo: %d\n", bn)

	file.Close()

	// Para ler um arquivo
	// Obs.: Como o arquivo é carregado na memória é importante ler pequenas partes
	// do arquivo para não sobrecarregá-la
	file, err = os.Open("teste")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, 8)

	for {
		nb, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		fmt.Println(string(buf[:nb]))
	}
}
