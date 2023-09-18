// 格式化日期格式
export function parseTime(date, format) {
  let curDate = new Date(date);
  const o = {
    "M+": curDate.getMonth() + 1, // 月份
    "d+": curDate.getDate(), // 日
    "h+":
      curDate.getHours() >= 10
        ? curDate.getHours()
        : "0" + curDate.getHours(), // 小时
    "m+":
      curDate.getMinutes() >= 10
        ? curDate.getMinutes()
        : "0" + curDate.getMinutes(), // 分
    "s+":
      curDate.getSeconds() >= 10
        ? curDate.getSeconds()
        : "0" + curDate.getSeconds(), // 秒
    "q+": Math.floor((curDate.getMonth() + 3) / 3), // 季度
    S: curDate.getMilliseconds(), // 毫秒
  };
  if (/(y+)/.test(format)) {
    format = format.replace(
      RegExp.$1,
      (curDate.getFullYear() + "").substr(4 - RegExp.$1.length)
    );
  }
  for (let k in o) {
    if (new RegExp("(" + k + ")").test(format)) {
      format = format.replace(
        RegExp.$1,
        RegExp.$1.length === 1
          ? o[k]
          : ("00" + o[k]).substr(("" + o[k]).length)
      );
    }
  }
  return format;
}