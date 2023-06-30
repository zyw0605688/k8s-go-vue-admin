import dayjs              from "dayjs";
import relativeTime       from "dayjs/plugin/relativeTime"
import "dayjs/locale/zh-cn";

dayjs.extend(relativeTime)
dayjs.locale("zh-cn");

export const getRelativeTime = (val) =>{
  const num = new Date(val).getTime() / 1000
  return dayjs(dayjs.unix(num).format("YYYY-MM-DD hh:mm")).fromNow();
}
