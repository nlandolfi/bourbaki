# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_balls.pdf
	make graph.pdf

metric_balls.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_balls.tex

metric_balls.pdf: ../../*.tex ../../trademark.pdf *.tex metric_balls.tex
	# pdflatex --file-line-error -interaction=nonstopmode metric_balls.tex
	../../latexrun metric_balls.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_balls.pdf
	open metric_balls.pdf

o:
	make open

terms: metric_balls.tex
	bbk_sheet -mode ts metric_balls.tex

spell:
	aspell -c metric_balls.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f metric_balls.*
	make reset
	bbk_sheet -mode mk > Makefile