GOPKG ?=	moul.io/srand
DOCKER_IMAGE ?=	moul/srand
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
