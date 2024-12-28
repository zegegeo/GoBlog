.PHONY: start stop

start:
	@echo "Starting all services..."
	@cd service/user && go run user.go & echo $$! > user.pid
	@cd service/article && go run article.go & echo $$! > article.pid
	@cd service/comment && go run comment.go & echo $$! > comment.pid
	@cd api && go run gateway.go & echo $$! > gateway.pid
	@echo "All services started"

stop:
	@echo "Stopping all services..."
	@-kill `cat service/user/user.pid` 2>/dev/null || true
	@-kill `cat service/article/article.pid` 2>/dev/null || true
	@-kill `cat service/comment/comment.pid` 2>/dev/null || true
	@-kill `cat api/gateway.pid` 2>/dev/null || true
	@rm -f service/user/user.pid service/article/article.pid service/comment/comment.pid api/gateway.pid
	@echo "All services stopped" 