module.exports = {
    serverRuntimeConfig: {
        // Will only be available on the server side
        NEWS_API_HOST_NAME: process.env.NEWS_API_HOST_NAME, // Pass through env variables
        NEWS_API_HOST_PORT: process.env.NEWS_API_HOST_PORT
      },
      publicRuntimeConfig: {
        // Will be available on both server and client
      }
}