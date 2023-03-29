export const fileReader = (file: any, options ?: any) : Promise<any> =>  {
    options = options || {};
    return new Promise(function (resolve, reject) {
        let reader = new FileReader();

        reader.onload = function () {
            resolve(reader);
        };
        reader.onerror = reject;

        if (options.accept && !new RegExp(options.accept).test(file.type)) {
            reject({
                code: 1,
                msg: 'wrong file type'
            });
        }

        if (!file.type || /^text\//i.test(file.type)) {
            reader.readAsText(file);
        } else {
            reader.readAsDataURL(file);
        }
    });
}


 
 
