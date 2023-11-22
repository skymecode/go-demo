package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// go解密算法
func main() {
	// 回调地址接受到的加密内容
	msgSecret := "jF5gtaRtxPX1XqOhuN+uFs+FgkYfTLBUh29OrghmTEAfVMqXGYAO4k1mbKuzhZnksf5BPW2/8jmi1zIYwUG9jtZTWbLgwZB57EowtXP5H8qfsu4gBClufVT9cebZHYyFtYZ64btsIF8EUtsU9+p5bvNNB/k7urXfR+e7hk8XOMhfZI70ZRa+BEoNY8Br6kkcfZdV4S6mimViJJkhLBkbbCN22ZKDx/swzp8wGmI9y9RLLtjGPzUv6DJ+xyKu8fYTQjeYXHu9T/fGOhT6nbXbod0KBv8zODT9glbAXQg5ZSEKb8H2wSTBMy2SzswAksUPfRPnhMtSMbL164gybAd7TAXA7F+HLYxfvN2gPd4yDTlH1s2/CYGVLXAMoZVxR+hBY8QHn+f+sxLiAU1TstuRscT1AxIh5UY8KGKIrZi86qEa657JEPGFS2m1i4s5/eZwHxTj2D79l5uV0ek6vmqAumz/VLC0Uf7tpK96pAwdkA8ttifxkn3aknfzPgZgsp56czqk4sci88GHCQwv4IQn5zqA/GrmmmHNLxjsQ4tojMavrGMjrOI28k+s8H1YUWwvoBH6J5cyMi9Yb9JdL7/BKz9mtv4zpr7rQBuVe+rVyMHNa+bZ/eL2HCbqoj7OT8FGogPZQntmTfoXjfPn4AscUQrAIwpT5Z6edo19Ae4qj3kjGJ/Q0JwoTjTslIeAfRc9LRCjECNpHLTzHICi7NDDLrVE2L+OT6FeRrbYwpNGFQvcKNtL5ad3dXXj17jIob9SaijRKsDJZBHfIgxOGjHNhA20NPF/mRfxdU1erqX1N3gZsOEtPIui0yQt1X+Q02hD" //加密之后的消息
	// 企业的 app_secret
	appSecret := "nkwpdVyagzTcFXJt"
	// 获取解密后的内容
	sourceMsg := AesDecrypt(msgSecret, appSecret)
	fmt.Printf("解密后:%s\n", sourceMsg)
}

func AesDecrypt(msgSecret, appSecret string) string {
	bytesPass, err := base64.StdEncoding.DecodeString(msgSecret)
	if err != nil {
		fmt.Println(err)
		return "解密失败！！！"
	}
	sourceMsg, err := DoAesDecrypt(bytesPass, []byte(appSecret))
	if err != nil {
		fmt.Println(err)
		return "解密失败！！！"
	}
	return string(sourceMsg)
}

func DoAesDecrypt(encryptedMsg, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(encryptedMsg))
	blockMode.CryptBlocks(origData, encryptedMsg)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// 去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unfilledNum := int(origData[length-1])
	return origData[:(length - unfilledNum)]
}
