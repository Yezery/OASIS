import request from "@/utils/axiosRequest"
export function updateSaleactive(nftInf) {
  console.log(nftInf);
  return request({
    url: `/UpdateNFTOwnerList`,
    method: 'post',
    data: nftInf
  })
}

export function insertSale(nftInf) {
  return request({
    url: `/createSale`,
    method: 'post',
    data: nftInf
  })
}

export function deleteSale(nftInf) {
  return request({
    url: `/DeleteSale`,
    method: 'post',
    data: nftInf
  })
}



