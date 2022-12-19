import axios from "axios";
import cryptoJs from 'crypto-js'
//引入crypto-js
//引入axios
//定义baseUrl  接口网址
const baseUrl = "http://localhost:9090";

export function getUser() { 
    return axios.get(`${baseUrl}/user`) 
}

export function setUser(username,password) { 
    return axios.post(`${baseUrl}/user/login?username=${username}&password=${password}`) 
}
export function SetUser(username) { 
    return axios.post(`${baseUrl}/user/username?username=${username}`) 
}
export function insertUser(username,password,email) { 
    return axios.post(`${baseUrl}/user/userinset?username=${username}&password=${password}&email=${email}`) 
}

export const Url = "http://localhost:9090/user/login";



export const encryptDes = (message, key) => {
    let keyHex = cryptoJs.enc.Utf8.parse(key)
    let option = { mode: cryptoJs.mode.ECB, padding: cryptoJs.pad.Pkcs7 }
    let encrypted = cryptoJs.DES.encrypt(message, keyHex, option)
    return encrypted.ciphertext.toString()
}

export const decryptDes = (message, key) => {
    let keyHex = cryptoJs.enc.Utf8.parse(key);
    let decrypted = cryptoJs.DES.decrypt(
        {
            ciphertext: cryptoJs.enc.Hex.parse(message)
        },
        keyHex,
        {
            mode: cryptoJs.mode.CBC,
            padding: cryptoJs.pad.Pkcs7
        }
    )
    return decrypted.toString(cryptoJs.enc.Utf8)
}
