build:
	@go build -o inventory *.go

clean:
	@rm inventory

test:
	@go test ./...

install: build
	@mv ./inventory /etc/ansible/hosts/libvirt-inventory

uninstall:
	@rm -f /etc/ansible/hosts/libvirt-inventory
