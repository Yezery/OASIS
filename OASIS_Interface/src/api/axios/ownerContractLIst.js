import request from "@/utils/axiosRequest"
export function postOwnerContractList(currentAddress) {
  return request({
    url: `/getOwnerNFTs`,
    method: 'post',
    data:currentAddress
  })
}

export function getOwnerUpSaleNFTs(data) {
  return request({
    url: `/GetOwnerUpSaleNFTs`,
    method: 'post',
    data:data
  })
}

export function getOwnerNFTsByAddress(data) {
  return request({
    url: `/getOwnerNFTsByAddress`,
    method: 'post',
    data:data
  })
}

export function updateNFTOwnerListAfterBuy(data) {
  return request({
    url: `/UpdateNFTOwnerListAfterBuy`,
    method: 'post',
    data:data
  })
}


export function getSeriesByNFTAddress(data) {
  return request({
    url: `/getSeriesByNFTAddress`,
    method: 'post',
    data:data
  })
}


export function search(SearchVo) {
  return request({
    url: `/Search`,
    method: 'post',
    data:SearchVo
  })
}

export function mainSearch(SearchVo) {
  return request({
    url: `/mainSearch`,
    method: 'post',
    data:SearchVo
  })
}

