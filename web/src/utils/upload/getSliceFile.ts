import { fileHash } from "./fileManipulation";

export /**
* @param {File} file 切片文件
* @param {number} pieceSizes 切片大小
* @param {string} fileKey 文件唯一标识
*/
    const getSliceFile = async (file: File, pieceSizes = 1, fileKey: string) => {
        const piece = 1024 * 1024 * pieceSizes;
        // 文件总大小
        const totalSize = file.size;
        const fileName = file.name;
        // 每次上传的开始字节
        let start = 0;
        let index = 1;
        // 每次上传的结尾字节
        let end = start + piece;
        const chunks = [];
        while (start < totalSize) {
            const current = Math.min(end, totalSize);
            // 根据长度截取每次需要上传的数据
            // File对象继承自Blob对象，因此包含slice方法
            const blob = file.slice(start, current);
            const hash = (await fileHash(blob)) as string;

            // 特殊处理，对接阿里云大文件上传
            chunks.push({
                file: blob,
                size: totalSize,
                index,
                fileSizeInByte: totalSize,
                name: fileName,
                fileName,
                hash,
                sliceSizeInByte: blob.size,
                progress: 0,
                fileKey,
            });
            start = current;
            end = start + piece;
            index += 1;
        }
        return chunks;
    };