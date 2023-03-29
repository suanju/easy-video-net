export const liveKeyDesensitization = (key: string) :string => {
    const str = key.slice(0, 8) + "****" + key.slice(key.length - 4);
    return str
}

export const  getLocation = (url :string) =>{
    if(url){
      let aDom = document.createElement('a');
      aDom.href = url;
      let j = {
          hostname: aDom.hostname,
          host: aDom.host,
          origin: aDom.origin,
          protocol: aDom.protocol,
          pathname: aDom.pathname,
          hash: aDom.hash,
          domain : ""
      }
      let Domain = j.hostname.match(/([a-z0-9][a-z0-9\-]*?\.(?:com\.cn|net\.cn|org\.cn|com\.au|imipo\.shop|com|cn|co|net|org|gov|cc|biz|info|hn|xyz|hk|icu|us|mobi|art|wang|me|so|top|win|vip|ltd|red|ru|ac\.cn|xn--kput3i|xin|xn--3ds443g|shop|site|club|fun|online|link|gov\.cn|name|pro|work|tv|kim|group|tech|store|cx|ren|ink|pub|live|wiki|design|xn--fiq228c5hs|xn--6qq986b3xl|xn--fiqs8s|xn--ses554g|xn--hxt814e|xn--55qx5d|xn--io0a7i|xn--3bst00m)(?:\.(?:cn|jp))?)$/);
      if(Domain){
          j.domain = Domain[0];
      }
      return j;
    }
  }
