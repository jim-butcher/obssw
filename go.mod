module github.com/jim-butcher/obssw

go 1.16

require obsws v0.0.0-20200720193653-c4fed10356a5

replace obsws => github.com/christopher-dG/go-obs-websocket v0.0.0-20200720193653-c4fed10356a5

require internal/util v0.0.0

replace internal/util => ./internal/pkg/util
