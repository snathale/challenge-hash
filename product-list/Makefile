# RUN ALL PROJECT TESTS
run-tests:
	docker-compose exec product-list php /opt/app/vendor/phpunit/phpunit/phpunit --configuration /opt/app/phpunit.xml

# GET ALL PROJECT DEPENDENCE
get-deps:
	docker-compose exec product-list composer update

