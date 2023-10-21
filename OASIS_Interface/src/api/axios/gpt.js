import request from "@/utils/axiosRequest"
export function sendToGPT(data,cancelToken) {
  return request({
    url: `/sendToGPT`,
    method: 'post',
    data: data,
    cancelToken
  })
}
