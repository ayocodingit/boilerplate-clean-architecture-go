#app
APP_NAME=Clean Architecture
APP_ENV=local
APP_PORT_HTTP=3000
APP_LOG=info

#database
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=root
DB_NAME=testing
DB_POOL_MIN=1
DB_POOL_MAX=10
DB_POOL_ACQUIRE=60000
DB_POOL_IDLE=60000
DB_KEEP_ALIVE=true

#JWT
JWT_ACCESS_SECRET=token-secret
JWT_ALGORITHM=HS256

#Redis
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_TTL=0
