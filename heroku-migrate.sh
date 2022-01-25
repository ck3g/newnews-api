# from https://g14a.dev/posts/schema-migrations-heroku/

wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz &&
tar -zxvf /tmp/migrate.linux-amd64.tar.gz -C /tmp &&
/tmp/migrate -database "${DATABASE_URL}" -path db/migrations up
exit_status=$?
RED='\033[0;31m'

if [ $exit_status -ne 0 ]; then
    echo "${RED}Migration did not complete smoothly. Rolling back..."
    /tmp/migrate -database "${DATABASE_URL}" -path db/migrations down --all
    exit 1;
fi