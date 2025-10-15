package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)  //将私钥从 *ecdsa.PrivateKey 转换成字节切片[]byte,二进制私钥(256 位)转换成32个字节
	fmt.Println(privateKeyBytes)  
    fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 每个字节转换成两个16进制,同时去掉'0x'
    fmt.Println()

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)//类型断言，把接口转换为具体的 *ecdsa.PublicKey
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println("og pubKey:", hexutil.Encode(publicKeyBytes)) 
    fmt.Println("real pubKey:  ", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
    
    //方法一自动
    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Println(address)

    //方法二手动
    hash := sha3.NewLegacyKeccak256()
    hash.Write(publicKeyBytes[1:])
    fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
    fmt.Println("Processed:", hexutil.Encode(hash.Sum(nil)[12:])) // 处理后的  原长32位，截去12位，保留后20位

    // b := []byte{0x04, 0xAA, 0xBB, 0xCC}
    // fmt.Printf("%x\n", b[1:]) 
    // s := "0x04aabbcc"
    // fmt.Println(s[4:]) // 输出 "aabbcc"
}