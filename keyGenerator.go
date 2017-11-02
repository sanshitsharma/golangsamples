package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := rand.Reader
	bitSize := 512

	var totalTime time.Duration
	iter := 100

	for i := 0; i < iter; i++ {
		start := time.Now()
		//key, err := rsa.GenerateKey(reader, bitSize)
		_, err := rsa.GenerateKey(reader, bitSize)
		generateTime := time.Now().Sub(start)

		totalTime += generateTime

		checkError(err)

		/*
		fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
		fmt.Println("Private key exponent", key.D.String())

		publicKey := key.PublicKey
		fmt.Println("Public key modulus", publicKey.N.String())
		fmt.Println("Public key exponent", publicKey.E)
		*/

		//fmt.Println(reflect.TypeOf(key))
		//fmt.Println("Iter[", i, "] --> Time to generate RSA key =", generateTime)

		// Save File
		//pemFilePath := "id_rsa_" + strconv.FormatInt(int64(bitSize), 10)
		//savePEMKey(pemFilePath, key)
	}

	avgTime := totalTime.Nanoseconds()/int64(iter)
	fmt.Println("Average time =", time.Duration(avgTime))
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	//start := time.Now()
	outFile, err := os.Create(fileName)
	checkError(err)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)

	outFile.Close()

	//generateTime := time.Now().Sub(start)
	//fmt.Println("Time to generate id_rsa = ", generateTime)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
