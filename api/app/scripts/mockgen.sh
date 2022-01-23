#!/bin/bash

# モックを作成する。

mockgen -source=./internal/domain/repository/bookmark.go -destination=./test/mock/domain/repository/bookmark.go
mockgen -source=./internal/domain/service/bookmark.go -destination=./test/mock/domain/service/bookmark.go
mockgen -source=./internal/application/usecase/bookmark.go -destination=./test/mock/application/usecase/bookmark.go
