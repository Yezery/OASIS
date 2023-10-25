import request from "@/utils/axiosRequest"
export function updateSaleactive(data) {
  return request({
    url: `/UpdateNFTOwnerList`,
    method: 'post',
    data: data
  })
}

export function insertSale(data) {
  return request({
    url: `/createSale`,
    method: 'post',
    data: data
  })
}

export function deleteSale(data) {
  return request({
    url: `/DeleteSale`,
    method: 'post',
    data: data
  })
}

export function getOnSaleNFTByNFTAddress(data) {
  return request({
    url: `/getOnSaleNFTByNFTAddress`,
    method: 'post',
    data: data
  })
}

export function makeNewTransaction(data) {
  return request({
    url: `/makeNewTransaction`,
    method: 'post',
    data: data
  })
}

export function scheduleDailySummary(data) {
  return request({
    url: `/scheduleDailySummary`,
    method: 'post',
    data: data
  })
}


