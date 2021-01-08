make:
	go run compile_all.go

sets.html: compiler.go sets.tex
	go run compiler.go --in sets.tex > sets.html

index.html: compile_all.go
	rm sheets/*
	rm index.html
	go run compile_all.go

o: index.html
	open index.html
