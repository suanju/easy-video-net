import { lib, SHA256 } from 'crypto-js'
import fileToArrayBuffer from 'file-to-array-buffer'

// 获取文件后缀
export const fileSuffix = (filename: string) => {
    const pos = filename.lastIndexOf('.')
    let suffix = ''
    if (pos !== -1) {
        suffix = filename.substring(pos)
    }
    return suffix;
}

export const arrayBufferToWordArray = (ab: any) => {
    const i8a = new Uint8Array(ab)
    const a = []
    for (let i = 0; i < i8a.length; i += 4) {
        a.push((i8a[i] << 24) | (i8a[i + 1] << 16) | (i8a[i + 2] << 8) | i8a[i + 3])
    }
    return lib.WordArray.create(a, i8a.length)
}

// 获取文件Hash
export const fileHash = async (file: Blob): Promise<String> => {
    const buffer = await fileToArrayBuffer(file)
    return SHA256(arrayBufferToWordArray(buffer)).toString()
}
