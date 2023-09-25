import request from "@/utils/axiosRequest"

export function SaleTypeList() {
  return request({
    url: `/getTypeList`,
    method: 'get',
  })
}