.PHONY : watch clean bower

-include Makefile.local

gulp=node_modules/.bin/coffeegulp
bower=node_modules/.bin/bower

watch: bower
	@$(gulp) --require coffee-script watch

clean:
	@rm -rf www

bower:
	@$(bower) install

