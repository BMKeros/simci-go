BASEDIR := ${CURDIR}

mockgen:
	mockery --case=snake --outpkg=storagemocks --output=internal/platform/storage/storagemocks --dir=internal --name=UserRepository