PASSWORD=dummy_passowrd
LIMIT=60
NEXT_WAIT_TIME=0
until echo '0' | arangoinspect --server.ask-jwt-secret false --server.username "root" --server.password "${PASSWORD}" --quiet true | grep "dignostics collected" || [ $NEXT_WAIT_TIME -eq $LIMIT ]; do
   sleep 1
   echo "retry init arangodb: $(( ++NEXT_WAIT_TIME ))"
done

if [ $NEXT_WAIT_TIME -eq $LIMIT ]; then
  echo "

  ArangoDB not available!
  Initial script don't run

  "
  exit
fi

arangosh --server.password ${PASSWORD} --javascript.execute-string '
var list = db._databases(); 
if (list.includes("dummy_discount_db")) {
  db._dropDatabase("dummy_discount_db");
} 
db._createDatabase("dummy_discount_db");'

arangosh --server.password ${PASSWORD} --server.database dummy_discount_db --javascript.execute-string '
db._drop("product");
db._create("product", { keyOptions: { type: "autoincrement" } });'

arangosh --server.password ${PASSWORD} --server.database dummy_discount_db --javascript.execute-string '
db._drop("user");
db._create("user", { keyOptions: { type: "autoincrement" } });'

arangoimp --file /opt/tools/products.json --collection product --create-collection true --server.database dummy_discount_db --server.password ${PASSWORD}
arangoimp --file /opt/tools/users.json --collection user --create-collection true --server.database dummy_discount_db --server.password ${PASSWORD}