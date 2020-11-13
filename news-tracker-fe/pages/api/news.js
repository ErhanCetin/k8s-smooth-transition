// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import getConfig from 'next/config'

const { serverRuntimeConfig, publicRuntimeConfig } = getConfig()

export default (req, res) => {
  //return fetch('http://'+serverRuntimeConfig.NEWS_API_HOST_NAME+':'+serverRuntimeConfig.NEWS_API_HOST_PORT+'/news/getAll')
  return fetch('http://127.0.0.1:61560/news/getAll')
  .then((response) => response.json())
  .then((data) => {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json')
    res.json(data);
  });
}
