import request from "@/utils/axiosRequest"

// export function positionCategoryList(data) {
//   return request({
//     url: '/category/position_category_list',
//     method: 'post',
//     data
//   })
// }
export function SaleTypeList() {
  return request({
    url: `/getTypeList`,
    method: 'get',
  })
}