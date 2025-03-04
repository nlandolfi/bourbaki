all: */*.tex
	cd all/ && make open

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

graph: graph.pdf
	make graph.pdf
	open graph.pdf

open:
	make graph

o:
	open www/screen/index.html

clean:
	rm */*.aux
	rm */*.out
	rm */*.log

# sed -i -e 's_find_replace_g" ./sheets/*/*.tex

cp:
	git add *
	git commit -m "checkpoint"
	git push
