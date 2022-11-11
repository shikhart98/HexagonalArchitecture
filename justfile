run:
    docker-compose down > /dev/null 2>&1
    docker-compose up -d

initdb:
    docker exec -it hex_db bash -c "mysql < docker-entrypoint-initdb.d/init.sql"

enterdb:
    docker exec -it hex_db bash -c "mysql"