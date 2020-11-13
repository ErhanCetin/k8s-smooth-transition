// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import getConfig from 'next/config'

const { serverRuntimeConfig, publicRuntimeConfig } = getConfig()

export default (req, res) => {
  return fetch('http://'+serverRuntimeConfig.NEWS_API_HOST_NAME+':'+serverRuntimeConfig.NEWS_API_HOST_PORT+'/news/getAll')
  .then((response) => response.json())
  .then((data) => {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json')
    res.json(data);
  });
}
