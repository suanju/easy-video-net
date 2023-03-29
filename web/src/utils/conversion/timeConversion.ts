import dayjs from "dayjs"
import isToday from "dayjs/plugin/isToday"

export const timetoRFC3339 = (date: Date) => {
	let y = date.getFullYear()
	let m = date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : (date.getMonth() + 1)
	let d = date.getDate() < 10 ? '0' + date.getDate() : date.getDate()
	let hh = date.getHours() < 10 ? '0' + date.getHours() : date.getHours();
	let mm = date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()
	let ss = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()
	var endDate = y + '-' + m + '-' + d + ' ' + hh + ':' + mm + ':' + ss
	endDate = endDate.replace(/\s+/g, 'T') + '+08:00'


	return endDate
}


export const rFC3339ToTime = (dateStr: string): string => {
	let date = new Date(dateStr).toJSON();
	let newDate: string
	return newDate = new Date(+new Date(date) + 8 * 3600 * 1000).toISOString().replace(/T/g, ' ').replace(/\.[\d]{3}Z/, '');
}



export const formattingSecondTime = (time: number) => {
	let t = '';
	if (time > -1) {
		let min: number | string
		let hour: number | string
		hour = Math.floor(time / 3600)
		if (hour < 10) {
			hour = "0" + hour
		}
		min = Math.floor(time / 60) % 60
		if (min < 10) {
			min = "0" + min
		}
		let sec: number | string
		sec = time % 60
		if (hour != 0) t += hour + ":"
		if (min != 0) t += min + ":"
		if (hour == 0 && min == 0) {
			if (sec == 0) {
				sec = "00"
			}
			t = "00 :" + sec
		} else {
			if (sec == 0) {
				sec = "00"
			}
			t += sec
		}
	}
	return t
}



export const timestampFormat = (timestamp: number | string) => {
	if (typeof timestamp == "string") {
		timestamp = Date.parse(timestamp);
	}

	if (!timestamp) return '';

	const zeroize = (num: number) => {
		return (String(num).length == 1 ? '0' : '') + num;
	}

	const curTimestamp = Date.parse(new Date().toDateString()); //当前时间戳
	const timestampDiff = curTimestamp - timestamp; // 参数时间戳与当前时间戳相差毫秒数
	const curDate = new Date(curTimestamp); // 当前时间日期对象
	const tmDate = new Date(timestamp); // 参数时间戳转换成的日期对象

	let Y = tmDate.getFullYear(),
		m = tmDate.getMonth() + 1,
		d = tmDate.getDate();
	let H = tmDate.getHours(),
		i = tmDate.getMinutes(),
		s = tmDate.getSeconds();
	if (timestampDiff < 60) { // 一分钟以内
		return "刚刚";
	} else if (timestampDiff < 3600) { // 一小时前之内
		return Math.floor(timestampDiff / 60) + "分钟前";
	} else if (curDate.getFullYear() == Y && curDate.getMonth() + 1 == m && curDate.getDate() == d) {
		return '今天' + zeroize(H) + ':' + zeroize(i);
	} else {
		let newDate = new Date((curTimestamp - 86400) * 1000); // 参数中的时间戳加一天转换成的日期对象
		if (newDate.getFullYear() == Y && newDate.getMonth() + 1 == m && newDate.getDate() == d) {
			return '昨天' + zeroize(H) + ':' + zeroize(i);
		} else if (curDate.getFullYear() == Y) {
			return zeroize(m) + '月' + zeroize(d) + '日 ' + zeroize(H) + ':' + zeroize(i);
		} else {
			return Y + '年' + zeroize(m) + '月' + zeroize(d) + '日 ' + zeroize(H) + ':' + zeroize(i);
		}
	}
}


export const recordTimeFormat = (timestamp: string) => {
	dayjs.extend(isToday)
	let is = dayjs(timestamp).isToday()
	if (is) {
		return dayjs(timestamp).format("HH:MM")
	} else {
		return dayjs(timestamp).format("YYYY:MM:DD")
	}
}

export const chatListTimeFormat = (timestamp: string) => {
	dayjs.extend(isToday)
	let is = dayjs(timestamp).isToday()
	if (is) {
		return dayjs(timestamp).format("HH:MM")
	} else {
		return dayjs(timestamp).format("MM:DD")
	}
}