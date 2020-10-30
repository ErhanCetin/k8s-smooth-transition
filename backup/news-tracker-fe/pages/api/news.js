// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

export default (req, res) => {
  //fetch('http://'+Config.newsapiHostName+':'+Config.newsApiHostPort+'/news/getAll')
  return fetch('http://127.0.0.1:53939/news/getAll')
  .then((response) => response.json())
  .then((data) => {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json')
    res.json(data);
  });
}
