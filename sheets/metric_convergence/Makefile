# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_convergence.pdf
	make graph.pdf

metric_convergence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_convergence.tex

metric_convergence.pdf: ../../*.tex ../../trademark.pdf *.tex metric_convergence.tex
	pdflatex --file-line-error -interaction=nonstopmode metric_convergence.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_convergence.pdf
	open metric_convergence.pdf

o:
	make open

defs: metric_convergence.tex
	defs metric_convergence.tex

spell:
	aspell -c metric_convergence.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metric_convergence.*
	make reset
	bbk_sheet -mode mk > Makefile