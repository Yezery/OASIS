import request from "@/utils/axiosRequest"
export function getToken(user) {
  return request({
    url: `/getToken`,
    method: 'post',
    data: user
  })
}

export function checkUserExist(user) {
  return request({
    url: `/checkMnemonic`,
    method: 'post',
    data: user
  })
}

export function setMnemonic(user) {
  return request({
    url: `/setMnemonic`,
    method: 'post',
    data: user
  })
}