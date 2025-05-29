.PHONY: run
run: swaybar-panel
	./swaybar-panel

swaybar-panel: *.go
	go build .
