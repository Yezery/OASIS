import request from "@/utils/axiosRequest"
export function getToken(data) {
  return request({
    url: `/getToken`,
    method: 'post',
    data: data
  })
}

export function checkUserExist(data) {
  return request({
    url: `/checkMnemonic`,
    method: 'post',
    data: data
  })
}

export function setMnemonic(data) {
  return request({
    url: `/setMnemonic`,
    method: 'post',
    data: data
  })
}

export function setAuthenticationMetaInformation(data) {
  return request({
    url: `/setAuthenticationMetaInformation`,
    method: 'post',
    data: data
  })
}

export function forgetMnemonic(data) {
  return request({
    url: `/forgetMnemonic`,
    method: 'post',
    data: data
  })
}
export function resetMnemonic(data) {
  return request({
    url: `/resetMnemonic`,
    method: 'post',
    data: data
  })
}
