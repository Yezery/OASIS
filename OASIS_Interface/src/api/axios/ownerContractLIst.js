import request from "@/utils/axiosRequest"
export function postOwnerContractList(currentAddress) {
  return request({
    url: `/getOwnerNFTs`,
    method: 'post',
    data:currentAddress
  })
}

export function getOwnerUpSaleNFTs(nftInf) {
  return request({
    url: `/GetOwnerUpSaleNFTs`,
    method: 'post',
    data:nftInf
  })
}

export function getOwnerNFTsByAddress(nftInf) {
  return request({
    url: `/getOwnerNFTsByAddress`,
    method: 'post',
    data:nftInf
  })
}

export function updateNFTOwnerListAfterBuy(nftInf) {
  return request({
    url: `/UpdateNFTOwnerListAfterBuy`,
    method: 'post',
    data:nftInf
  })
}

export function search(SearchVo) {
  return request({
    url: `/Search`,
    method: 'post',
    data:SearchVo
  })
}