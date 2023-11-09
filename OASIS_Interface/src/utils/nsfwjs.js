import * as nsfwjs from 'nsfwjs'
export async function process(img) {
  const model = await nsfwjs.load('/model/nsfwModel/', { size: 299 })
  const predictions = await model.classify(img)
  // Share results
  console.log('Predictions: ', predictions)
  return predictions
} 

export async function processGif(img) {
  const model = await nsfwjs.load('/model/nsfwModel/', { size: 299 })
  const predictions = await model.classifyGif(img)
  // Share results
  console.log('Predictions: ', predictions)
  return predictions
} 